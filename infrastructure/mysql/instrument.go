package mysql_repository

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/hobord/invst-portfolio-backend-golang/domain/entity"
	"github.com/hobord/invst-portfolio-backend-golang/domain/repository"
)

// InstrumentRepository Implements repository.InstrumentRepository
type InstrumentMysqlRepository struct {
	conn *sql.DB
}

// NewInstrumentMysqlRepository returns initialized InstrumentRepositoryImpl
func NewInstrumentMysqlRepository(conn *sql.DB) repository.InstrumentRepository {
	return &InstrumentMysqlRepository{conn: conn}
}

func (r *InstrumentMysqlRepository) queryRow(ctx context.Context, q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := r.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRowContext(ctx, args...), nil
}

func (r *InstrumentMysqlRepository) query(ctx context.Context, q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := r.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(ctx, args...)
}

func (r *InstrumentMysqlRepository) GetByID(ctx context.Context, id int) (*entity.Instrument, error) {
	querySTR := "SELECT id, name, symbol, instrumentType FROM instrument WHERE id=?"
	row, err := r.queryRow(ctx, querySTR, id)
	if err != nil {
		return nil, err
	}

	entity := &entity.Instrument{}
	err = row.Scan(&entity.ID, &entity.Name, &entity.Symbol, &entity.Type)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func parseRows(rows *sql.Rows) ([]*entity.Instrument, error) {
	results := make([]*entity.Instrument, 0)

	for rows.Next() {
		entity := &entity.Instrument{}
		err := rows.Scan(&entity.ID, &entity.Name, &entity.Symbol, &entity.Type)
		if err != nil {
			return nil, err
		}
		results = append(results, entity)
	}
	return results, nil
}

func (r *InstrumentMysqlRepository) GetAll(ctx context.Context) ([]*entity.Instrument, error) {
	querySTR := "SELECT id, name, symbol, instrumentType FROM instrument"
	rows, err := r.query(ctx, querySTR)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results, err := parseRows(rows)

	return results, err
}

// SEARCH
func (r *InstrumentMysqlRepository) Search(ctx context.Context, keyword string) ([]*entity.Instrument, error) {
	querySTR := "SELECT id, name, symbol, instrumentType FROM instrument WHERE symbol LIKE concat('%',?,'%')"
	rows, err := r.query(ctx, querySTR, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results, err := parseRows(rows)

	return results, err
}

// TODO:
func (r *InstrumentMysqlRepository) Save(ctx context.Context, entity *entity.Instrument) error {
	var querySTR string
	isExists, err := r.GetByID(ctx, entity.ID)
	if isExists == nil || entity.ID == 0 {
		querySTR = "INSERT INTO instrument (name, symbol, instrumentType) VALUES (?, ?, ?)"
	} else {
		querySTR = "UPDATE instrument SET name = ?, symbol = ?, instrumentType = ? WHERE ID = ?"
	}
	stmt, err := r.conn.Prepare(querySTR)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if isExists == nil {
		_, err = stmt.ExecContext(ctx, entity.ID, entity.Name, entity.Symbol, entity.Type)
	} else {
		_, err = stmt.ExecContext(ctx, entity.Name, entity.Symbol, entity.Type, entity.ID)
	}
	return err
}

func (r *InstrumentMysqlRepository) Delete(ctx context.Context, id int) error {
	querySTR := "DELETE FROM entity WHERE id=?"
	stmt, err := r.conn.Prepare(querySTR)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}
