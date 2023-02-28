package db

import (
	"context"

	pgxuuid "github.com/jackc/pgx-gofrs-uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pauljamescleary/gomin/pkg/common/config"
	"github.com/rs/zerolog/log"
)

type Database struct {
	Conn *pgxpool.Pool
}

func NewDatabase(c config.Config) *Database {
	pool, _ := NewPgConnectionPool(c)
	return &Database{Conn: pool}
}

func NewPgConnectionPool(c config.Config) (*pgxpool.Pool, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dbUrl := c.DbUrl
	dbConfig, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, err
	}

	// Register UUID support
	dbConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxuuid.Register(conn.TypeMap())
		return nil
	}

	conn, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		return nil, err
	}
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msgf("Error connecting to the database. ")
	}
	return conn, nil
}
