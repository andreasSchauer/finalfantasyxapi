package seeding

import(
	"database/sql"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func queryInTransaction(db *database.Queries, dbConn *sql.DB, f func(*database.Queries) error) error {
	tx, err := dbConn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := db.WithTx(tx)

	if err := f(qtx); err != nil {
		return err
	}

	return tx.Commit()
}