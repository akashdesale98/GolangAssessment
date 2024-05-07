package service

import (
	"context"
	"database/sql"
	"sync"

	"github.com/akashdesale98/GolangAssessment/model"
	_ "github.com/lib/pq"
)

// Database interface for performing database operations
type Database interface {
	GetEmployeeByID(ctx context.Context, id string) (string, error)
	UpdateEmployee(ctx context.Context, emp *model.Employee) error
	GetEmployees(ctx context.Context, limit, offset int) (*[]model.Employee, error)
	CreateEmployee(ctx context.Context, emp *model.Employee) error
	DeleteEmployee(ctx context.Context, id string) error
}

// PostgreSQL implementation of the Database interface
type PostgresDB struct {
	db    *sql.DB
	mutex sync.Mutex
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

func (p *PostgresDB) UpdateEmployee(ctx context.Context, emp *model.Employee) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	// Execute SQL UPDATE statement to update employee details
	query := "UPDATE Employee SET name = $1, position = $2, salary = $3 WHERE id = $4"
	_, err := p.db.ExecContext(ctx, query, emp.Name, emp.Position, emp.Salary, emp.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) GetEmployees(ctx context.Context, limit, offset int) (*[]model.Employee, error) {

	rows, err := p.db.QueryContext(ctx, "SELECT id, name, position, salary FROM employee ORDER BY id LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through rows and construct employee slice
	var employees []model.Employee
	for rows.Next() {
		var employee model.Employee
		err := rows.Scan(&employee.ID, &employee.Name, &employee.Position, &employee.Salary)
		if err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	// Check for errors during row iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &employees, nil
}

func (p *PostgresDB) CreateEmployee(ctx context.Context, emp *model.Employee) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	// Execute query to insert new employee
	_, err := p.db.ExecContext(ctx, "INSERT INTO employee (name, position, salary) VALUES ($1, $2, $3)", emp.Name, emp.Position, emp.Salary)
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) DeleteEmployee(ctx context.Context, id string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	// Execute query to delete employee
	_, err := p.db.ExecContext(ctx, "DELETE FROM employee WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
