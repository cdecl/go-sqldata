package sqldata

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func assertTest(t *testing.T, valid bool, msg string) {
	if !valid {
		t.Error(msg)
	}
}

func datasetTest(t *testing.T) {
	db, err := sql.Open("mysql", "root:passwd!@tcp(centos:3306)/glass")
	assertTest(t, err == nil, "DB Open Error")
	defer db.Close()

	rows, err := db.Query("SELECT 8 as version, 'glass' as dbname ")
	assertTest(t, err == nil, "DB Query Error")
	defer rows.Close()

	ds := GetDataSet(rows)

	assertTest(t, ds[0]["version"] == "8", "Dataset Error")
	assertTest(t, ds[0]["dbname"] == "glass", "Dataset Error")
}
