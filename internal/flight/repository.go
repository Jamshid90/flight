package flight

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Jamshid90/flight/internal/database"
	"github.com/Jamshid90/flight/internal/entity"
	errpkg "github.com/Jamshid90/flight/internal/errors"

	sortpkg "github.com/Jamshid90/flight/internal/sort"

	"github.com/Jamshid90/flight/internal/utils"

	"github.com/Jamshid90/flight/pkg/slice"
	"github.com/jackc/pgx/v4"
)

type flightRepo struct {
	db               *database.PostgresDB
	tableName        string
	searchParameters []string
}

func NewRepository(db *database.PostgresDB) entity.FlightRepository {
	return &flightRepo{db, "flight", []string{
		"city_departure",
		"city_arrival",
	}}
}

func (r *flightRepo) Create(ctx context.Context, flight *entity.Flight) error {
	clauses := map[string]interface{}{
		"number":         flight.Number,
		"city_departure": flight.CityDeparture,
		"city_arrival":   flight.CityArrival,
		"time_departure": flight.TimeDeparture,
		"time_arrival":   flight.TimeArrival,
		"created_at":     flight.CreatedAt,
		"updated_at":     flight.UpdatedAt,
	}

	sqlStr, args, err := r.db.Sq.Builder.
		Insert(r.tableName).
		SetMap(clauses).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return fmt.Errorf("error during sql build, flight create: %w", err)
	}

	if err = r.db.QueryRow(ctx, sqlStr, args...).Scan(&flight.ID); err != nil {
		return fmt.Errorf("error during create flight: %w", err)
	}

	return nil
}

func (r *flightRepo) Read(ctx context.Context, id int64) (*entity.Flight, error) {
	var flight entity.Flight
	sqlStr, args, err := r.db.Sq.Builder.
		Select(
			"id",
			"number",
			"city_departure",
			"city_arrival",
			"time_departure",
			"time_arrival",
			"created_at",
			"updated_at",
		).
		From(r.tableName).
		Where(r.db.Sq.Equal("id", id)).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("error during sql build, flight read: %w", err)
	}

	err = r.db.QueryRow(ctx, sqlStr, args...).Scan(
		&flight.ID,
		&flight.Number,
		&flight.CityDeparture,
		&flight.CityArrival,
		&flight.TimeDeparture,
		&flight.TimeArrival,
		&flight.CreatedAt,
		&flight.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return nil, errpkg.NewErrNotFound("flight")
	}

	if err != nil {
		return nil, fmt.Errorf("error during read flight: %w", err)
	}

	return &flight, nil
}

func (r *flightRepo) Update(ctx context.Context, flight *entity.Flight) error {
	sqlStr, args, err := r.db.Sq.Builder.
		Update(r.tableName).
		SetMap(map[string]interface{}{
			"number":         flight.Number,
			"city_departure": flight.CityDeparture,
			"city_arrival":   flight.CityArrival,
			"time_departure": flight.TimeDeparture,
			"time_arrival":   flight.TimeArrival,
			"updated_at":     flight.UpdatedAt,
		}).
		Where(r.db.Sq.Equal("id", flight.ID)).
		ToSql()

	if err != nil {
		return fmt.Errorf("error during sql build, flight update: %w", err)
	}

	commandTag, err := r.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return fmt.Errorf("error during update flight: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("failed flight update")
	}

	return nil
}

func (r *flightRepo) Delete(ctx context.Context, id int64) error {
	sqlStr, args, err := r.db.Sq.Builder.
		Delete(r.tableName).
		Where(r.db.Sq.Equal("id", id)).
		ToSql()
	if err != nil {
		return fmt.Errorf("error during sql build, chat delete: %w", err)
	}

	commandTag, err := r.db.Exec(ctx, sqlStr, args...)
	if err != nil {
		return fmt.Errorf("error during delete chat: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("failed flight delete")
	}

	return nil
}

func (r *flightRepo) List(ctx context.Context, parameters map[string][]string) ([]*entity.Flight, error) {
	var list []*entity.Flight

	queryParameters := utils.NewQueryParameters(parameters)

	query := r.db.Sq.Builder.
		Select("id",
			"number",
			"city_departure",
			"city_arrival",
			"time_departure",
			"time_arrival",
			"created_at",
			"updated_at").
		From(r.tableName).
		Limit(queryParameters.GetLimit()).
		Offset(queryParameters.GetOffset())

	for k, v := range queryParameters.GetParameters() {
		if !slice.Contains(r.searchParameters, k) {
			continue
		}

		if k == "city_departure" {
			query = query.Where("city_departure ilike '%" + v + "%' ")
			continue
		}

		if k == "city_arrival" {
			query = query.Where("city_arrival ilike '%" + v + "%' ")
			continue
		}

		query = query.Where(r.db.Sq.Equal(k, v))
	}

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error during sql build, flight read all: %w", err)
	}

	rows, err := r.db.Query(ctx, sqlStr, args...)
	if err != nil {
		return nil, fmt.Errorf("error during flight read all sql query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var flight entity.Flight
		err := rows.Scan(
			&flight.ID,
			&flight.Number,
			&flight.CityDeparture,
			&flight.CityArrival,
			&flight.TimeDeparture,
			&flight.TimeArrival,
			&flight.CreatedAt,
			&flight.UpdatedAt,
		)
		if err != nil {
			return list, fmt.Errorf("error during scan flight read all %w", err)
		}
		list = append(list, &flight)
	}

	if sort := queryParameters.GetParameters()["sort"]; sort == "number" {
		list = sortpkg.SortByNumber(list)
	}

	return list, nil
}

func (r *flightRepo) CreateBulk(ctx context.Context, flights []*entity.Flight) error {

	if len(flights) == 0 {
		return nil
	}

	var (
		strBuild    strings.Builder
		args        []interface{}
		cnt         int
		columnCount = 7
		rowCount    = len(flights)
	)

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	strBuild.WriteString(`INSERT INTO flight( number,
											  city_departure,
											  city_arrival,
											  time_departure,
											  time_arrival,
											  created_at,
											  updated_at) VALUES`)

	for i, flight := range flights {

		strBuild.WriteString("(")
		for j := 0; j < columnCount; j++ {
			cnt++
			strBuild.WriteString("$")
			strBuild.WriteString(strconv.Itoa(cnt))
			if j != columnCount-1 {
				strBuild.WriteString(", ")
			}

		}
		strBuild.WriteString(")")
		if i != rowCount-1 {
			strBuild.WriteString(",")
		}

		args = append(args,
			flight.Number,
			flight.CityDeparture,
			flight.CityArrival,
			flight.TimeDeparture,
			flight.TimeArrival,
			flight.CreatedAt,
			flight.UpdatedAt,
		)
	}

	if _, err := tx.Exec(context.Background(), strBuild.String(), args...); err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
