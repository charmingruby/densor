package postgres_repository

import (
	"github.com/charmingruby/densor/internal/sensor/domain/entity"
	"github.com/jmoiron/sqlx"
)

const (
	findEquipmentByName = "find one equipment by name"
)

func equipmentQueries() map[string]string {
	return map[string]string{
		findEquipmentByName: `SELECT * FROM equipments
		WHERE name = $1
		`,
	}
}

func NewPostgresEquipmentRepository(db *sqlx.DB) (*PostgresEquipmentRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range equipmentQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				NewPreparationErr(queryName, "equipment", err)
		}

		stmts[queryName] = stmt
	}

	return &PostgresEquipmentRepository{
		db:    db,
		stmts: stmts,
	}, nil
}

type PostgresEquipmentRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func (r *PostgresEquipmentRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			NewStatementNotPreparedErr(queryName, "equipment")
	}

	return stmt, nil
}

func (r *PostgresEquipmentRepository) FindByName(name string) (entity.Equipment, error) {
	stmt, err := r.statement(findEquipmentByName)
	if err != nil {
		return entity.Equipment{}, err
	}

	var equipment entity.Equipment
	if err := stmt.Get(&equipment, name); err != nil {
		return entity.Equipment{}, err
	}

	return equipment, nil
}
