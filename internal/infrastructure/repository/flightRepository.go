package repository

import (
	"api-app/internal/config"
	"api-app/internal/domain/dto"
	"api-app/internal/domain/entity"
	"api-app/internal/domain/storage"
	"api-app/pkg/db/postgresql"
	"api-app/pkg/object/oid"
	"context"
	"errors"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	log "github.com/sirupsen/logrus"
)

type FlightRepository storage.Storage[entity.Flight, entity.FlightView, dto.FLightCreate]

type flightRepository struct {
	pool postgresql.PGXPool
	cfg  config.PostgresConfig
}

var _ FlightRepository = (*flightRepository)(nil)

func NewFlightRepository(pool postgresql.PGXPool, cfg config.PostgresConfig) FlightRepository {
	return &flightRepository{pool: pool, cfg: cfg}
}

func (fRepo *flightRepository) GetById(ctx context.Context, id oid.Id) (entity.Flight, error) {
	// TODO implement me
	panic("implement me")
}

func (fRepo *flightRepository) GetAll(ctx context.Context) ([]*entity.Flight, error) {
	rows, err := fRepo.pool.Query(ctx, "SELECT * FROM public.flights")
	defer rows.Close()
	if err != nil {
		log.Errorf("error can't get all flights: %s", err.Error())
		return nil, err
	}

	var flights []*entity.Flight
	err = pgxscan.ScanAll(&flights, rows)
	if err != nil {
		log.Errorf("error scanning rows of flights: %s", err.Error())
		return nil, err
	}
	fmt.Println(flights)

	return flights, nil
}

func (fRepo *flightRepository) Store(ctx context.Context, fC dto.FLightCreate) (*entity.Flight, error) {
	query, err := fRepo.pool.Query(
		ctx,
		"INSERT INTO public.flights (airl_code, flt_num, flt_date, orig_iata, dest_iata, departure_time, arrival_time, total_cash, correctly_parsed) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING (id)",
		fC.AirlCode,
		fC.FltNum,
		fC.FltDate,
		fC.OrigIATA,
		fC.DestIATA,
		fC.DepartureTime,
		fC.ArrivalTime,
		fC.TotalCash,
		fC.CorrectlyParsed,
	)
	if err != nil {
		log.Errorf("error inserting new flight: %s\n", err.Error())
		return &entity.Flight{}, err
	}
	defer query.Close()

	newFlight := fC.ToView().ToEntity(oid.Undefined)
	if !query.Next() {
		err := errors.New("error there's no returned id from sql")
		log.Errorln(err.Error())
		return newFlight, err // TODO WHAT TO DO WTF???!!!?
	}
	var id string
	err = query.Scan(&id)
	if err != nil {
		log.Errorf("error scanning new newFlight id: %s\n", err.Error())
		return newFlight, err // TODO WHAT TO DO WTF??!?!?!?!?
	}

	newFlight.Id = oid.ToId(id)
	return newFlight, err
}

func (fRepo *flightRepository) DeleteById(ctx context.Context, id oid.Id) (entity.Flight, error) {
	//TODO implement me
	panic("implement me")
}
