package domain

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name"`
	Email       string             `json:"email"`
	Role        string             `json:"role"`
	Password    string             `json:"password"`
	LoanHistory []Loan             `json:"loanHistory"`
}

type Claims struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	UserID primitive.ObjectID `json:"username"`
	Role   string             `json:"role"`
	jwt.StandardClaims
}

type ForgotClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type VerificationClaims struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Email    string             `json:"username"`
	Name     string             `json:"name"`
	Password string             `json:"password"`
	Role     string             `json:"role"`
	jwt.StandardClaims
}

type AuthUser struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

type UserUsecaseInterface interface {
	GenerateVerificationToken(user *VerificationClaims) error
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email *string) (*User, error)
	VerifyUser(token string) error
	Login(user *AuthUser) (string, string, error)
	ForgotPassword(user *Email) error
	VerifyForgotPassword(token string) (*VerificationClaims, error)

	GetAllUsers() ([]*User, error)
	UpdateUser(username, newPassword string) error
	DeleteUser(objectID primitive.ObjectID) error
}

type UserRepositoryInterface interface {
	//User operations
	Create(user *User) error
	GetUserByUsername(username string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id primitive.ObjectID) (*User, error)
	GetAllUsers() ([]*User, error)
	UpdateUser(username string, user *User) error
	DeleteUser(id primitive.ObjectID) error
}
