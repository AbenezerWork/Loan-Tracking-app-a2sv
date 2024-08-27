package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Loan struct {
	ID           primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Issuer       primitive.ObjectID `json:"issuer"`
	Amount       float64            `json:"amount"`
	Approved     bool               `json:"approved"`
	InterestRate float64            `json:"interest"`
	TermInMonths int                `json:"duration"`
	Status       time.Time          `json:"status"`
	CreatedAt    time.Time          `json:"createdAt"`
	EndDate      time.Time          `json:"enddate"`
	UpdatedAt    time.Time          `json:"updatedAt"`
}
type LoanUseCaseInterface interface {
	CreateLoan(ctx context.Context, borrowerID primitive.ObjectID, amount float64, interestRate float64, termInMonths int) (*Loan, error)
	GetLoanByID(ctx context.Context, id primitive.ObjectID) (*Loan, error)
	UpdateLoanStatus(ctx context.Context, id primitive.ObjectID, status bool) error
	DeleteLoan(ctx context.Context, id primitive.ObjectID) error
	ListLoans(ctx context.Context, filter map[string]interface{}) ([]Loan, error)
}
type LoanRepositoryInterface interface {
	CreateLoan(ctx context.Context, loan *Loan) (*Loan, error)
	GetLoanByID(ctx context.Context, id primitive.ObjectID) (*Loan, error)
	UpdateLoan(ctx context.Context, loan *Loan) error
	DeleteLoan(ctx context.Context, id primitive.ObjectID) error
	ListLoans(ctx context.Context, filter bson.M) ([]Loan, error)
}
