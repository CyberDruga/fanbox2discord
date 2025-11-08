package db

import (
	"database/sql"
	"embed"
	"log/slog"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/sqlite"
	"github.com/go-errors/errors"
	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql schema.sql
var fs embed.FS

type customWriter struct{}

func (*customWriter) Write(p []byte) (n int, err error) {
	slog.Info(string(p))
	return len(p), nil
}

func NewClient(dbLocation string) (queries *Queries, err error) {

	if dbLocation == "" {
		dbLocation = ":memory:"
	}

	u, err := url.Parse("sqlite:" + dbLocation)

	if err != nil {
		err = errors.New(err)
		return
	}

	dbmt := dbmate.New(u)
	dbmt.FS = fs
	dbmt.MigrationsDir = []string{"./migrations"}
	dbmt.SchemaFile = ""
	dbmt.Log = &customWriter{}

	err = dbmt.CreateAndMigrate()

	database, err := sql.Open("sqlite", dbLocation)
	if err != nil {
		err = errors.Join(errors.Errorf("Couldn't connect to the database"), err)
		return
	}

	queries = New(database)

	return

}
