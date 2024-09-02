package postgres_repository

import (
	"github.com/charmingruby/densor/internal/sensor/database/postgres_repository/mapper"
	"github.com/charmingruby/densor/internal/sensor/domain/entity"
	"github.com/jmoiron/sqlx"
)

const (
	createSensor   = "create sensor"
	findSensorByID = "find sensor by id"
)

func sensorQueries() map[string]string {
	return map[string]string{
		createSensor: `INSERT INTO sensors
		(id, name)
		VALUES ($1, $2)
		RETURNING *`,
		findSensorByID: `SELECT * FROM sensors 
		WHERE id = $1`,
	}
}

func NewPostgresSensorRepository(db *sqlx.DB) (*PostgresSensorRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range sensorQueries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				NewPreparationErr(queryName, "sensor", err)
		}

		stmts[queryName] = stmt
	}

	return &PostgresSensorRepository{
		db:    db,
		stmts: stmts,
	}, nil
}

type PostgresSensorRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func (r *PostgresSensorRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			NewStatementNotPreparedErr(queryName, "sensor")
	}

	return stmt, nil
}

func (r *PostgresSensorRepository) Store(e *entity.Sensor) error {
	stmt, err := r.statement(createSensor)
	if err != nil {
		return err
	}

	mappedEntity := mapper.DomainSensorToPostgres(*e)

	if _, err := stmt.Exec(
		mappedEntity.ID,
		mappedEntity.Name,
	); err != nil {
		return err
	}

	return nil
}

func (r *PostgresSensorRepository) FindMany(id string) (*entity.Sensor, error) {
	stmt, err := r.statement(findSensorByID)
	if err != nil {
		return nil, err
	}

	var sensor mapper.PostgresSensor
	if err := stmt.Get(&sensor, id); err != nil {
		return nil, err
	}

	mappedSensor := mapper.PostgresSensorToDomain(sensor)

	return &mappedSensor, nil
}
