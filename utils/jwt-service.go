package utils

import (
	"errors"
	"loan-tracking/domain"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Claims struct to hold JWT claims
type Claims struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Role     string             `json:"role"`
	jwt.StandardClaims
}

type JWTService struct {
	secretKey        string
	issuer           string
	refreshSecretKey string
}

// NewJWTService creates a new JWTService
func NewJWTService(secretKey, issuer, refreshSecretKey string) *JWTService {
	return &JWTService{
		secretKey:        secretKey,
		issuer:           issuer,
		refreshSecretKey: refreshSecretKey,
	}
}
func (j *JWTService) GenerateValidationToken(user *domain.VerificationClaims) (string, error) {
	claims := &domain.VerificationClaims{
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) GenerateForgotPasswordToken(user *domain.Email) (string, error) {
	claims := &domain.VerificationClaims{
		Email: user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

// GenerateToken generates a new JWT token
func (j *JWTService) GenerateToken(userID primitive.ObjectID, role string) (string, error) {
	claims := &Claims{
		ID:   userID,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(), // Shorter expiry time
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) GenerateRefreshToken(userID primitive.ObjectID, role string) (string, error) {
	// Set expiration time for refresh token
	expirationTime := time.Now().Add(time.Hour * 24 * 7)

	claims := &Claims{
		ID:   userID,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.refreshSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JWTService) ValidateRefreshToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.refreshSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	// Check if the token is expired
	if time.Unix(claims.ExpiresAt, 0).Before(time.Now()) {
		return nil, errors.New("refresh token expired")
	}

	return claims, nil
}

func (j *JWTService) ValidateToken(tokenString string) (*domain.VerificationClaims, error) {
	claims := &domain.VerificationClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Validate the token and claims
	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	return claims, nil
}
func (j *JWTService) ValidateValidateToken(tokenString string) (*domain.VerificationClaims, error) {
	claims := &domain.VerificationClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Validate the token and claims
	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	return claims, nil
}
