package usecase

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"loan-tracking/domain"
)

type loanUseCase struct {
	loanRepo domain.LoanRepositoryInterface
}

func NewLoanUseCase(loanRepo domain.LoanRepositoryInterface) domain.LoanUseCaseInterface {
	return &loanUseCase{
		loanRepo: loanRepo,
	}
}

func (u *loanUseCase) CreateLoan(ctx context.Context, borrowerID primitive.ObjectID, amount float64, interestRate float64, termInMonths int) (*domain.Loan, error) {
	// Validate input
	if amount <= 0 || interestRate <= 0 || termInMonths <= 0 {
		return nil, errors.New("invalid loan parameters")
	}

	// Create the loan struct
	loan := &domain.Loan{
		Issuer:       borrowerID,
		Amount:       amount,
		InterestRate: interestRate,
		TermInMonths: termInMonths,
		CreatedAt:    time.Now(),
		EndDate:      time.Now().AddDate(0, termInMonths, 0),
		Approved:     false,
	}

	// Save the loan using the domain
	createdLoan, err := u.loanRepo.CreateLoan(ctx, loan)
	if err != nil {
		return nil, err
	}

	return createdLoan, nil
}

func (u *loanUseCase) GetLoanByID(ctx context.Context, id primitive.ObjectID) (*domain.Loan, error) {
	loan, err := u.loanRepo.GetLoanByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if loan == nil {
		return nil, errors.New("loan not found")
	}

	return loan, nil
}

func (u *loanUseCase) UpdateLoanStatus(ctx context.Context, id primitive.ObjectID, status bool) error {

	// Get the loan to update
	loan, err := u.loanRepo.GetLoanByID(ctx, id)
	if err != nil {
		return err
	}

	if loan == nil {
		return errors.New("loan not found")
	}

	// Update the status
	loan.Approved = status
	loan.UpdatedAt = time.Now()

	// Save the updated loan
	err = u.loanRepo.UpdateLoan(ctx, loan)
	if err != nil {
		return err
	}

	return nil
}

func (u *loanUseCase) DeleteLoan(ctx context.Context, id primitive.ObjectID) error {
	// Check if the loan exists
	loan, err := u.loanRepo.GetLoanByID(ctx, id)
	if err != nil {
		return err
	}

	if loan == nil {
		return errors.New("loan not found")
	}

	// Delete the loan
	err = u.loanRepo.DeleteLoan(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *loanUseCase) ListLoans(ctx context.Context, filter map[string]interface{}) ([]domain.Loan, error) {

	// Convert the filter map to bson.M for MongoDB query
	bsonFilter := make(map[string]interface{})
	for key, value := range filter {
		bsonFilter[key] = value
	}

	loans, err := u.loanRepo.ListLoans(ctx, bsonFilter)
	if err != nil {
		return nil, err
	}

	return loans, nil
}
