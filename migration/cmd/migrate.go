package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"sort"
	"time"

	"example.com/expenses-tracker/pkg/database"
)

type MigrationHandler struct {
	database     *sql.DB
	sqlDirectory string
}

func NewMigrationHandler(sqlDirectory string) *MigrationHandler {
	db, err := database.NewDatabase()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return &MigrationHandler{
		database:     db,
		sqlDirectory: sqlDirectory,
	}
}

func (h *MigrationHandler) hasMigrationExecuted(migration string) (bool, error) {
	sql := `SELECT script, executed_at FROM database_migrations WHERE script = $1 LIMIT 1`
	rows, err := h.database.Query(sql, migration)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	return rows.Next(), nil
}

func (h *MigrationHandler) Execute() error {
	err := h.createMigrationsTableIfNotExists()
	if err != nil {
		return err
	}

	migrationFiles, err := h.getAllMigrationFiles()
	if err != nil {
		return err
	}

	var migrationsToExecute []string

	for _, file := range *migrationFiles {
		result, err := h.hasMigrationExecuted(file)
		if err != nil {
			return err
		}

		if result {
			fmt.Printf("\033[34mMigration %s has already executed\033[0m\n", file)
			continue
		}

		fmt.Printf("\033[33mWarning: Migration %s to be executed\033[0m\n", file)
		migrationsToExecute = append(migrationsToExecute, file)
	}

	for _, file := range migrationsToExecute {
		result, err := h.executeMigration(file)
		if err != nil {
			return err
		}

		if result {
			fmt.Printf("\033[32mMigration %s executed successfully\033[0m\n", file)
		}
	}

	return nil
}

func (h *MigrationHandler) executeMigration(migrationName string) (bool, error) {
	tx, err := h.database.Begin()
	if err != nil {
		return false, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	filePath := fmt.Sprintf("%s/%s", h.sqlDirectory, migrationName)
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	_, err = tx.Exec(string(fileContent))
	if err != nil {
		return false, err
	}

	_, err = tx.Exec("INSERT INTO database_migrations(script, executed_at) VALUES ($1, $2)", migrationName, time.Now())
	if err != nil {
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (h *MigrationHandler) createMigrationsTableIfNotExists() error {
	sql := `
	CREATE TABLE IF NOT EXISTS database_migrations (
		id SERIAL PRIMARY KEY,
		script VARCHAR(255) NOT NULL,
		executed_at TIMESTAMP NOT NULL
	);
	`

	_, err := h.database.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

func (h *MigrationHandler) getAllMigrationFiles() (*[]string, error) {
	files, err := os.ReadDir(h.sqlDirectory)
	var filenames []string
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			filenames = append(filenames, file.Name())
		}
	}

	sort.Strings(filenames)

	return &filenames, err
}
