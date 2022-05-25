package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/rs/zerolog/log"
	"os"
)

func NewDbConnectionPool() *sql.DB {
	host := os.Getenv("log_saver_db_host")
	port := os.Getenv("log_saver_db_port")
	user := os.Getenv("log_saver_db_user")
	password := os.Getenv("log_saver_db_password")
	dbName := os.Getenv("log_saver_db_name")

	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbName)
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	return db
}

type IDbRepository interface {
	SaveGet(content interface{}) (int64, error)
	SaveCreate(content interface{}) (int64, error)
}
