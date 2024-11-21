package repository

import (
	"context"
	"database/sql"
	"fmt"

	models "example.com/expenses-tracker/internal/pkg"
)

const (
	errFailedToCreateExpense      = "failed to create expense: %w"
	errFailedToGetExpenseById     = "failed to get expense by id: %w"
	errFailedToUpdateExpense      = "failed to update expense: %w"
	errFailedToDeleteExpense      = "failed to delete expense: %w"
	errFailedToGetExpensesForUser = "failed to get expenses for user: %w"
)

type expensesRepository struct {
	db *sql.DB
}

// NewExpensesRepository creates a new Postgresql repository for expenses
func NewExpensesRepository(db *sql.DB) *expensesRepository {
	return &expensesRepository{db}
}

func (r *expensesRepository) CreateExpense(ctx context.Context, model *models.Expense) error {
	query := "INSERT INTO expenses(id, amount, user_id, category_id, description, created_at) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.ExecContext(ctx, query, model.ID, model.Amount, model.User.ID, model.Category.ID, model.Description, model.CreatedAt.Format("2006-01-02 15:04:05"))
	if err != nil {
		return fmt.Errorf(errFailedToCreateExpense, err)
	}

	return nil
}

func (r *expensesRepository) GetExpense(ctx context.Context, id string) (*models.Expense, error) {
	var expense models.Expense

	query := "SELECT e.*, u.first_name, u.last_name, u.email FROM expenses WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&expense.ID, &expense.Amount, &expense.User.ID, &expense.Category.ID, &expense.Description, &expense.User.CreatedAt, &expense.User.FirstName, &expense.User.LastName, expense.User.Email)
	if err != nil {
		return nil, fmt.Errorf(errFailedToGetExpenseById, err)
	}

	return &expense, nil
}

func (r *expensesRepository) UpdateExpense(ctx context.Context, model *models.Expense) error {
	query := "UPDATE expenses SET amount = $1, description = $2, category_id = $3 WHERE id = $4"
	_, err := r.db.ExecContext(ctx, query, model.Amount, model.Description, model.Category.ID, model.ID)
	if err != nil {
		return fmt.Errorf(errFailedToUpdateExpense, err)
	}

	return nil
}

func (r *expensesRepository) DeleteExpense(ctx context.Context, model *models.Expense) error {
	query := "DELETE FROM expenses WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, model.ID)
	if err != nil {
		return fmt.Errorf(errFailedToDeleteExpense, err)
	}

	return nil
}

func (r *expensesRepository) GetAllForUser(ctx context.Context, user *models.User) (models.Expenses, error) {
	query := "SELECT * FROM expenses WHERE user_id = $1 ORDER BY id DESC"
	rows, err := r.db.QueryContext(ctx, query, user.ID)
	if err != nil {
		return nil, fmt.Errorf(errFailedToGetExpensesForUser, err)
	}

	defer rows.Close()
	var expenses models.Expenses
	for rows.Next() {
		expense := &models.Expense{
			User: user,
		}

		if err := rows.Scan(&expense.ID, &expense.Amount, &expense.User.ID, &expense.Category.ID, &expense.Description, &expense.CreatedAt); err != nil {
			return nil, fmt.Errorf(errFailedToGetExpensesForUser, err)
		}

		expenses = append(expenses, expense)
	}

	return expenses, nil
}
