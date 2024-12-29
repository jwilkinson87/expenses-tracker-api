package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"example.com/expenses-tracker/pkg/models"
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
	_, err := r.db.ExecContext(ctx, query, model.ID, model.Amount, model.User.ID, model.Category.ID, model.Description, model.CreatedAt)
	if err != nil {
		return fmt.Errorf(errFailedToCreateExpense, err)
	}

	return nil
}

func (r *expensesRepository) GetExpense(ctx context.Context, id string) (*models.Expense, error) {
	var expense models.Expense
	var user models.User
	var category models.Category

	query := `
		SELECT
			e.id, e.amount, e.expense_date, e.description, e.created_at,
			c.id AS category_id, c.name AS category_name,
			u.id AS user_id, u.first_name, u.last_name, u.email, u.password, u.created_at
		FROM expenses e
		JOIN categories c ON e.category_id = c.id
		JOIN users u ON e.user_id = u.id
		WHERE e.id = $1
	`
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&expense.ID, &expense.Amount, &expense.Date, &expense.Description, &expense.CreatedAt,
		&category.ID, &category.Label,
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf(errFailedToGetExpenseById, err)
	}

	expense.Category = &category
	expense.User = &user

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

func (r *expensesRepository) GetAllForUser(ctx context.Context, userModel *models.User) (models.Expenses, error) {
	query := `
		SELECT
			e.id, e.amount, e.expense_date, e.description, e.created_at,
			c.id AS category_id, c.name AS category_name,
			u.id AS user_id, u.first_name, u.last_name, u.email, u.password, u.created_at
		FROM expenses e
		JOIN categories c ON e.category_id = c.id
		JOIN users u ON e.user_id = u.id
		WHERE u.id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, userModel.ID)
	if err != nil {
		return nil, fmt.Errorf(errFailedToGetExpensesForUser, err)
	}

	defer rows.Close()
	var expenses models.Expenses
	for rows.Next() {
		expense := &models.Expense{}
		category := &models.Category{}
		user := &models.User{}

		if err := rows.Scan(
			&expense.ID, &expense.Amount, &expense.Date, &expense.Description, &expense.CreatedAt,
			&category.ID, &category.Label,
			&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf(errFailedToGetExpensesForUser, err)
		}

		expense.Category = category
		expense.User = user

		expenses = append(expenses, expense)
	}

	return expenses, nil
}
