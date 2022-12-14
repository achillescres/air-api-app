package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/achillescres/saina-api/internal/domain/entity"
	"github.com/achillescres/saina-api/internal/domain/storage"
	"github.com/achillescres/saina-api/internal/domain/storage/dto"
	"github.com/achillescres/saina-api/pkg/db/postgresql"
	"github.com/achillescres/saina-api/pkg/object/oid"
	"github.com/georgysavva/scany/v2/pgxscan"
	log "github.com/sirupsen/logrus"
)

type FlightRepository storage.FlightStorage

type flightRepository struct {
	pool postgresql.PGXPool
}

var _ FlightRepository = (*flightRepository)(nil)

func NewFlightRepository(pool postgresql.PGXPool) FlightRepository {
	return &flightRepository{pool: pool}
}

func (fRepo *flightRepository) GetById(ctx context.Context, id oid.Id) (*entity.Flight, error) {
	// TODO implement me
	panic("implement me")
}

func (fRepo *flightRepository) GetAll(ctx context.Context) ([]*entity.Flight, error) {
	fmt.Println(fRepo)
	fmt.Println(fRepo.pool)
	rows, err := fRepo.pool.Query(ctx, "SELECT * FROM public.flights")
	if err != nil {
		log.Errorf("error can't get all flights: %s", err)
		return nil, err
	}
	defer rows.Close()

	var flights []*entity.Flight
	err = pgxscan.ScanAll(&flights, rows)
	if err != nil {
		log.Errorf("error scanning rows of flights: %s", err)
		return nil, err
	}
	fmt.Println(flights)

	return flights, nil
}

func (fRepo *flightRepository) GetAllInMap(ctx context.Context) (map[oid.Id]*entity.Flight, error) {
	flightsSlice, err := fRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	flightsMap := map[oid.Id]*entity.Flight{}
	for _, flight := range flightsSlice {
		flightsMap[flight.Id] = flight
	}

	return flightsMap, nil
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
		log.Errorf("error inserting new flight: %s\n", err)
		return nil, err
	}
	defer query.Close()

	if !query.Next() {
		err := errors.New("error there's no returned id from sql")
		log.Errorln(err)
		return nil, err // TODO WHAT TO DO WTF???!!!?
	}
	var id string
	err = query.Scan(&id)
	if err != nil {
		log.Errorf("error scanning new newFlight id: %s\n", err)
		return nil, err // TODO WHAT TO DO WTF??!?!?!?!?
	}

	return fC.ToEntity(oid.ToId(id)), err
}

func (fRepo *flightRepository) DeleteById(ctx context.Context, id oid.Id) (*entity.Flight, error) {
	//TODO implement me
	panic("implement me")
}
