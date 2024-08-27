package usecase

import (
	"errors"
	"fmt"
	"loan-tracking/domain"
	"loan-tracking/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	jwtSvc    *utils.JWTService
	userRepo  domain.UserRepositoryInterface
	tokenRepo domain.TokenRepositoryInterface
}

func NewUserUsecase(js *utils.JWTService, userRepo domain.UserRepositoryInterface, tokenRepo domain.TokenRepositoryInterface) *UserUsecase {
	return &UserUsecase{
		jwtSvc:    js,
		tokenRepo: tokenRepo,
		userRepo:  userRepo,
	}
}

func (u *UserUsecase) GetUserByUsername(username string) (*domain.User, error) {
	return u.userRepo.GetUserByUsername(username)
}
func (u *UserUsecase) GetUserByEmail(email *string) (*domain.User, error) {
	return u.userRepo.GetUserByEmail(*email)
}

// Login authenticates a user and returns JWT and refresh tokens if successful
func (u *UserUsecase) Login(authUser *domain.AuthUser) (string, string, error) {
	fmt.Println("authuser: ", authUser)
	user, err := u.userRepo.GetUserByUsername(authUser.Username)
	if err != nil {
		return "", "", err
	}

	fmt.Println("user: ", user)

	if err := utils.CheckPasswordHash(user.Password, authUser.Password); err != nil {
		return "", "", errors.New("invalid username or password2")
	}

	// Generate JWT and refresh tokens for the authenticated user
	token, err := u.jwtSvc.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := u.jwtSvc.GenerateRefreshToken(user.ID, user.Role)
	if err != nil {
		return "", "", err
	}

	refreshedTokenClaim := &domain.RefreshToken{
		UserID:    user.ID,
		Role:      user.Role,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
	}

	// Save the refresh token in the database
	err = u.tokenRepo.SaveRefreshToken(refreshedTokenClaim)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func (v *UserUsecase) GenerateVerificationToken(user *domain.VerificationClaims) error {
	token, err := v.jwtSvc.GenerateValidationToken(user)
	if err != nil {
		return err
	}
	utils.SendTokenEmail(user.Email, token)
	return nil
}

func (v *UserUsecase) ForgotPassword(userdata *domain.VerificationClaims) error {
	fmt.Println(userdata)
	user, err := v.GetUserByEmail(&userdata.Email)
	if err != nil {
		return err
	}
	veruser := &domain.VerificationClaims{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Role:     user.Role,
	}
	token, err := v.jwtSvc.GenerateValidationToken(veruser)
	if err != nil {
		return err
	}
	fmt.Println("email: ", user.Email)
	err = utils.SendForgotPasswordTokenEmail(user.Email, token)
	if err != nil {
		return err
	}
	return nil
}

func (v *UserUsecase) VerifyForgotPassword(token string) (*domain.VerificationClaims, error) {
	claims, err := v.jwtSvc.ValidateValidateToken(token)
	if err != nil {
		return &domain.VerificationClaims{}, err
	}

	v.UpdateUser(claims.Name, claims.Password)

	return claims, nil
}

func (v *UserUsecase) VerifyUser(token string) error {
	claims, err := v.jwtSvc.ValidateValidateToken(token)
	if err != nil {
		return err
	}
	user := &domain.User{
		Email:    claims.Email,
		Name:     claims.Name,
		Password: claims.Password,
	}
	return v.userRepo.Create(user)
}

func (u *UserUsecase) UpdateUser(username, newPassword string) error {
	existingUser, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}
	existingUser.Password = hashedPassword

	return u.userRepo.UpdateUser(username, existingUser)
}

func (u *UserUsecase) GetAllUsers() ([]*domain.User, error) {
	users, err := u.userRepo.GetAllUsers()
	return users, err
}

// DeleteUser deletes a user by ID
func (u *UserUsecase) DeleteUser(objectID primitive.ObjectID) error {
	return u.userRepo.DeleteUser(objectID)
}
