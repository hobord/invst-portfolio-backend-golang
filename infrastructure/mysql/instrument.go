package mysql_repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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

func (r *InstrumentMysqlRepository) GetByID(ctx context.Context, id int) (*entity.Instrument, error) {
	querySTR := "SELECT instrumentId, name, symbol, instrumentType FROM instrument WHERE instrumentId=?;"
	row, err := queryRow(ctx, r.conn, querySTR, id)
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

func (r *InstrumentMysqlRepository) parseRows(rows *sql.Rows) ([]*entity.Instrument, error) {
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
	querySTR := "SELECT instrumentId, name, symbol, instrumentType FROM instrument;"
	rows, err := query(ctx, r.conn, querySTR)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results, err := r.parseRows(rows)

	return results, err
}

// SEARCH
func (r *InstrumentMysqlRepository) Search(ctx context.Context, keyword string) ([]*entity.Instrument, error) {
	querySTR := "SELECT instrumentId, name, symbol, instrumentType FROM instrument WHERE symbol LIKE concat('%', ?, '%');"
	rows, err := query(ctx, r.conn, querySTR, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results, err := r.parseRows(rows)

	return results, err
}

// List
func (r *InstrumentMysqlRepository) List(ctx context.Context, keyword string, offset, limit int) ([]*entity.Instrument, error) {
	querySTR := "SELECT instrumentId, name, symbol, instrumentType FROM instrument"
	queryParams := make([]interface{}, 0)
	if keyword != "" {
		querySTR = querySTR + " WHERE symbol LIKE concat('%', ?, '%')"
		queryParams = append(queryParams, keyword)
	}
	querySTR = querySTR + " ORDER BY instrumentId DESC"
	if limit > 0 {
		querySTR = fmt.Sprintf("%s LIMIT %d, %d", querySTR, offset, limit)
	}
	querySTR = querySTR + ";"
	log.Println(querySTR)
	rows, err := query(ctx, r.conn, querySTR, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results, err := r.parseRows(rows)

	return results, err
}

func (r *InstrumentMysqlRepository) TotalCountOfList(ctx context.Context, keyword string) (int, error) {
	var cnt int
	querySTR := "SELECT count(instrumentId) FROM instrument"
	queryParams := make([]interface{}, 0)
	if keyword != "" {
		querySTR = querySTR + " WHERE symbol LIKE concat('%', ?, '%')"
		queryParams = append(queryParams, keyword)
	}
	querySTR = querySTR + ";"
	result, err := queryRow(ctx, r.conn, querySTR, queryParams...)
	if err != nil {
		return 0, err
	}
	result.Scan(&cnt)
	return cnt, nil
}

func (r *InstrumentMysqlRepository) Save(ctx context.Context, entity *entity.Instrument) error {
	var querySTR string
	isExists, err := r.GetByID(ctx, entity.ID)
	if isExists == nil || entity.ID == 0 {
		querySTR = "INSERT INTO instrument (name, symbol, instrumentType) VALUES (?, ?, ?);"
	} else {
		querySTR = "UPDATE instrument SET name = ?, symbol = ?, instrumentType = ? WHERE instrumentId = ?;"
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
	querySTR := "DELETE FROM entity WHERE instrumentId=?;"
	stmt, err := r.conn.Prepare(querySTR)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	return err
}
