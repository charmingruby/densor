package postgres_repository

import (
	"github.com/charmingruby/densor/internal/sensor/domain/entity"
	"github.com/jmoiron/sqlx"
)

const (
	findSectorByName = "find one sector by name"
)

func sectorQueries() map[string]string {
	return map[string]string{
		findSectorByName: `SELECT * FROM sectors
		WHERE name = $1
		`,
	}
}

func NewPostgresSectorRepository(db *sqlx.DB) (*PostgresSectorRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range sectorQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				NewPreparationErr(queryName, "sector", err)
		}

		stmts[queryName] = stmt
	}

	return &PostgresSectorRepository{
		db:    db,
		stmts: stmts,
	}, nil
}

type PostgresSectorRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func (r *PostgresSectorRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			NewStatementNotPreparedErr(queryName, "sector")
	}

	return stmt, nil
}

func (r *PostgresSectorRepository) FindByName(name string) (entity.Sector, error) {
	stmt, err := r.statement(findSectorByName)
	if err != nil {
		return entity.Sector{}, err
	}

	var sector entity.Sector
	if err := stmt.Get(&sector, name); err != nil {
		return entity.Sector{}, err
	}

	return sector, nil
}
