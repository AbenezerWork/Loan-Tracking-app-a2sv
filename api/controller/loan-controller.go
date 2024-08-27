package controller

import (
	"context"
	"net/http"

	"loan-tracking/domain"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanController struct {
	loanUseCase domain.LoanUseCaseInterface
}

func NewLoanController(loanUseCase domain.LoanUseCaseInterface) *LoanController {
	return &LoanController{
		loanUseCase: loanUseCase,
	}
}

// CreateLoan handles POST /loans
func (lc *LoanController) CreateLoan(c *gin.Context) {
	var input domain.Loan

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	borrowerID, err := primitive.ObjectIDFromHex(input.Issuer.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid borrower ID"})
		return
	}

	loan, err := lc.loanUseCase.CreateLoan(context.Background(), borrowerID, input.Amount, input.InterestRate, input.TermInMonths)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loan)
}

// GetLoan handles GET /loans/:id
func (lc *LoanController) GetLoan(c *gin.Context) {
	id := c.Param("id")
	loanID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan ID"})
		return
	}

	loan, err := lc.loanUseCase.GetLoanByID(context.Background(), loanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if loan == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "loan not found"})
		return
	}

	c.JSON(http.StatusOK, loan)
}

// UpdateLoanStatus handles PUT /loans/:id/status
func (lc *LoanController) UpdateLoanStatus(c *gin.Context) {
	id := c.Param("id")
	loanID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan ID"})
		return
	}

	var input struct {
		Status bool `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = lc.loanUseCase.UpdateLoanStatus(context.Background(), loanID, input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "loan status updated successfully"})
}

// DeleteLoan handles DELETE /loans/:id
func (lc *LoanController) DeleteLoan(c *gin.Context) {
	id := c.Param("id")
	loanID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid loan ID"})
		return
	}

	err = lc.loanUseCase.DeleteLoan(context.Background(), loanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "loan deleted successfully"})
}

// ListLoans handles GET /loans
func (lc *LoanController) ListLoans(c *gin.Context) {
	filter := map[string]interface{}{}

	if borrowerID := c.Query("borrower_id"); borrowerID != "" {
		id, err := primitive.ObjectIDFromHex(borrowerID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid borrower ID"})
			return
		}
		filter["borrower_id"] = id
	}

	if status := c.Query("status"); status != "" {
		filter["status"] = status
	}

	loans, err := lc.loanUseCase.ListLoans(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loans)
}
