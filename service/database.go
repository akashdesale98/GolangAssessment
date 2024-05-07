package service

import (
	"context"
	"database/sql"

	"github.com/akashdesale98/GolangAssessment/model"
	_ "github.com/lib/pq"
)

// Database interface for performing database operations
type Database interface {
	GetEmployeeByID(ctx context.Context, id string) (string, error)
}

// PostgreSQL implementation of the Database interface
type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	// Initialize PostgreSQL connection
	db, err := sql.Open("postgres", "postgres://postgres:mysecretpassword@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &PostgresDB{db: db}, nil
}

func (p *PostgresDB) GetEmployeeByID(ctx context.Context, id string) (*model.Employee, error) {

	var emp model.Employee
	err := p.db.QueryRowContext(ctx, "SELECT id, name, salary, position FROM employee WHERE id = $1", id).
		Scan(&emp.ID, &emp.Name, &emp.Salary, &emp.Position)
	if err != nil {
		return nil, err
	}

	return &emp, nil
}
