package postgres_repository

import (
	"github.com/charmingruby/densor/internal/sensor/domain/entity"
	"github.com/jmoiron/sqlx"
)

const (
	findManySensorCategories = "find many sensor categories"
	findSensorCategoryByName = "find one sensor category by name"
)

func sensorcategoryQueries() map[string]string {
	return map[string]string{
		findManySensorCategories: `SELECT * FROM sensor_categories`,
		findSensorCategoryByName: `SELECT * FROM sensor_categories
		WHERE name = $1
		`,
	}
}

func NewPostgresSensorCategoryRepository(db *sqlx.DB) (*PostgresSensorCategoryRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range sensorcategoryQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				NewPreparationErr(queryName, "sensor category", err)
		}

		stmts[queryName] = stmt
	}

	return &PostgresSensorCategoryRepository{
		db:    db,
		stmts: stmts,
	}, nil
}

type PostgresSensorCategoryRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func (r *PostgresSensorCategoryRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			NewStatementNotPreparedErr(queryName, "sensor category")
	}

	return stmt, nil
}

func (r *PostgresSensorCategoryRepository) FindMany() ([]entity.SensorCategory, error) {
	stmt, err := r.statement(findManySensorCategories)
	if err != nil {
		return nil, err
	}

	var categories []entity.SensorCategory
	if err := stmt.Select(&categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *PostgresSensorCategoryRepository) FindByName(name string) (entity.SensorCategory, error) {
	stmt, err := r.statement(findSensorCategoryByName)
	if err != nil {
		return entity.SensorCategory{}, err
	}

	var category entity.SensorCategory
	if err := stmt.Get(&category, name); err != nil {
		return entity.SensorCategory{}, err
	}

	return category, nil
}
