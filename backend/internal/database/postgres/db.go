package postgres

import (
	"context"
	"os"

	"github.com/Suhach/fasion-store/backend/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var Pool *pgxpool.Pool

func Init() error {
	dbURL := os.Getenv("DB_URL")
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		logger.Log.Error("fail to connect to database", zap.String("db", "postgres"))
		return err
	}
	Pool = pool
	return nil
}
