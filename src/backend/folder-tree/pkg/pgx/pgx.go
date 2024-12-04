package pgx

import (
	"context"
	"fmt"
	"folder-tree/config"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPgx(cfg config.Config) (*pgxpool.Pool, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)

	poolcfg, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, err
	}
	poolcfg.MaxConns = int32(cfg.MaxConnection)
	poolcfg.MaxConnIdleTime = time.Duration(cfg.MaxIdleConnection) * time.Minute
	poolcfg.MaxConnLifetime = time.Duration(cfg.ConnectionLifeTime) * time.Minute

	conn, err := pgxpool.ConnectConfig(context.Background(), poolcfg)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
