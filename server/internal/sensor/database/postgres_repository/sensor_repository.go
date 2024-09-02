package postgres_repository

import (
	"github.com/charmingruby/densor/internal/sensor/database/postgres_repository/mapper"
	"github.com/charmingruby/densor/internal/sensor/domain/entity"
	"github.com/jmoiron/sqlx"
)

const (
	createSensor    = "create sensor"
	findManySensors = "find many sensors"
)

func sensorQueries() map[string]string {
	return map[string]string{
		createSensor: `INSERT INTO sensors
		(id, name, description, sensor_category_id, equipment_id, observation, sector_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING *`,
		findManySensors: `SELECT * FROM sensors`,
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

func (r *PostgresSensorRepository) Store(e entity.Sensor) error {
	stmt, err := r.statement(createSensor)
	if err != nil {
		return err
	}

	mappedEntity := mapper.DomainSensorToPostgres(e)

	if _, err := stmt.Exec(
		mappedEntity.ID,
		mappedEntity.Name,
		mappedEntity.Description,
		mappedEntity.SensorCategoryID,
		mappedEntity.EquipmentID,
		mappedEntity.Observation,
		mappedEntity.SectorID,
	); err != nil {
		return err
	}

	return nil
}

func (r *PostgresSensorRepository) FindMany() ([]entity.Sensor, error) {
	stmt, err := r.statement(findManySensors)
	if err != nil {
		return nil, err
	}

	var pgSensors []mapper.PostgresSensor
	if err := stmt.Select(&pgSensors); err != nil {
		return nil, err
	}

	var sensors []entity.Sensor
	for _, s := range pgSensors {
		sensors = append(sensors,
			mapper.PostgresSensorToDomain(s),
		)
	}

	return sensors, nil
}
