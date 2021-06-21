package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M) {
	databaseUrl = os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		databaseUrl = "host=localhost dbname=restapi_test sslmode=disable user=oleksandr password=toor"
	}

	os.Exit(m.Run())
}
