package db

import (
	"ecommerce/config"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // has to be imported
)

/* Library: sqlc, sqlx, ORM
1. sqlc: More speed, not developer freiend, powerful query write complicated
2. sqlx: midium fast
3. ORM: heavy, less power & speed, developer friendly, poweful task
*/

func GetConnectionString(cnf *config.Config) string {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cnf.DbHost,
		cnf.DbPort,
		cnf.DbUser,
		cnf.DbPassword,
		cnf.DbName,
		cnf.DbSslMode,
	)

	return connectionString
}

func NewDBConnection(cnf *config.Config) (*sqlx.DB, error) {
	connectionString := GetConnectionString(cnf)

	client, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
		return nil, err
	}

	return client, nil
}
