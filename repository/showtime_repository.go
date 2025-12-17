package repository

import (
	"context"
	"errors"
	"technical-test/config"
	"technical-test/database"
	"technical-test/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
)

type ShowtimeRepository interface {
	Create(ctx context.Context, tx pgx.Tx, showtime model.Showtime) (model.Showtime, error)
	GetByID(ctx context.Context, tx pgx.Tx, id int64) (model.Showtime, error)
	GetAllAvailable(ctx context.Context, tx pgx.Tx) ([]model.Showtime, error)
	Update(ctx context.Context, tx pgx.Tx, showtime model.Showtime) error
	Delete(ctx context.Context, tx pgx.Tx, id int64) error
}

type ShowtimeRepositoryImpl struct {
	WrapDB *database.WrapDB
	Env    *config.EnvironmentVariable
}

func NewShowtimeRepository(
	wrapDB *database.WrapDB,
	env *config.EnvironmentVariable,
) ShowtimeRepository {
	return &ShowtimeRepositoryImpl{
		WrapDB: wrapDB,
		Env:    env,
	}
}
func (r *ShowtimeRepositoryImpl) GetByID(ctx context.Context, tx pgx.Tx, id int64) (showtime model.Showtime, err error) {
	query := `
		SELECT 
			id, movie_id, studio_id, show_date, start_time, status
		FROM 
			showtimes 
		WHERE 
			id = $1`

	var conn pgx.Row
	if tx != nil {
		conn = tx.QueryRow(ctx, query, id)
	} else {
		conn = r.WrapDB.Postgres.QueryRow(ctx, query, id)
	}

	err = conn.Scan(
		&showtime.ID,
		&showtime.MovieID,
		&showtime.StudioID,
		&showtime.ShowDate,
		&showtime.StartTime,
		&showtime.Status,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return model.Showtime{}, errors.New("showtime not found")
	} else if err != nil {
		log.Error().Err(err).Int64("id", id).Msg("Failed to get showtime by ID")
		return model.Showtime{}, err
	}

	return showtime, nil
}

func (r *ShowtimeRepositoryImpl) GetAllAvailable(ctx context.Context, tx pgx.Tx) (showtimes []model.Showtime, err error) {
	query := `
		SELECT 
			id, movie_id, studio_id, show_date, start_time, status
		FROM 
			showtimes 
		ORDER BY 
			show_date ASC, start_time ASC`

	var rows pgx.Rows
	if tx != nil {
		rows, err = tx.Query(ctx, query)
	} else {
		rows, err = r.WrapDB.Postgres.Query(ctx, query)
	}

	if err != nil {
		log.Error().Err(err).Msg("Failed to get all available showtimes")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s model.Showtime
		err = rows.Scan(
			&s.ID,
			&s.MovieID,
			&s.StudioID,
			&s.ShowDate,
			&s.StartTime,
			&s.Status,
		)
		if err != nil {
			log.Error().Err(err).Msg("Failed to scan showtime row")
			return nil, err
		}
		showtimes = append(showtimes, s)
	}

	if err = rows.Err(); err != nil {
		log.Error().Err(err).Msg("Error during showtime row iteration")
		return nil, err
	}

	return showtimes, nil
}

func (r *ShowtimeRepositoryImpl) Create(ctx context.Context, tx pgx.Tx, s model.Showtime) (model.Showtime, error) {
	query := `
		INSERT INTO showtimes (movie_id, studio_id, show_date, start_time, status,duration_minutes) 
		VALUES ($1, $2, $3, $4, $5,$6) 
		RETURNING id`

	conn := r.WrapDB.Postgres.QueryRow(ctx, query, s.MovieID, s.StudioID, s.ShowDate, s.StartTime, s.Status, s.DurationMinutes)

	if tx != nil {
		conn = tx.QueryRow(ctx, query, s.MovieID, s.StudioID, s.ShowDate, s.StartTime, s.Status, s.DurationMinutes)
	} else {
		conn = r.WrapDB.Postgres.QueryRow(ctx, query, s.MovieID, s.StudioID, s.ShowDate, s.StartTime, s.Status, s.DurationMinutes)
	}

	err := conn.Scan(&s.ID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create showtime")
		return model.Showtime{}, err
	}

	return s, nil
}

func (r *ShowtimeRepositoryImpl) Update(ctx context.Context, tx pgx.Tx, s model.Showtime) error {
	query := `
		UPDATE showtimes 
		SET movie_id = $2, studio_id = $3, show_date = $4, start_time = $5, status = $6 
		WHERE id = $1`

	var commandTag pgconn.CommandTag
	var err error

	if tx != nil {
		commandTag, err = tx.Exec(ctx, query, s.ID, s.MovieID, s.StudioID, s.ShowDate, s.StartTime, s.Status)
	} else {
		commandTag, err = r.WrapDB.Postgres.Exec(ctx, query, s.ID, s.MovieID, s.StudioID, s.ShowDate, s.StartTime, s.Status)
	}

	if err != nil {
		log.Error().Err(err).Int64("id", s.ID).Msg("Failed to update showtime")
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("showtime update failed: ID not found")
	}

	return nil
}

func (r *ShowtimeRepositoryImpl) Delete(ctx context.Context, tx pgx.Tx, id int64) error {
	query := `UPDATE showtimes SET status = 'CANCELLED' WHERE id = $1`

	var commandTag pgconn.CommandTag
	var err error

	if tx != nil {
		commandTag, err = tx.Exec(ctx, query, id)
	} else {
		commandTag, err = r.WrapDB.Postgres.Exec(ctx, query, id)
	}

	if err != nil {
		log.Error().Err(err).Int64("id", id).Msg("Failed to delete (cancel) showtime")
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return errors.New("showtime cancellation failed: ID not found")
	}

	return nil
}
