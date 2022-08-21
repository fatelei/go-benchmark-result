package gobenchmarkresult

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func BenchmarkMySQLSelectWithoutInterpolateParams(t *testing.B) {
	db, err := sql.Open("mysql", "root:12345678@/test")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	for i := 0; i < t.N; i++ {
		db.QueryRow("SELECT * FROM test WHERE id = ?", 1)
	}
}

func BenchmarkMySQLSelectWithInterpolateParams(t *testing.B) {
	db, err := sql.Open("mysql", "root:12345678@/test?interpolateParams=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	for i := 0; i < t.N; i++ {
		db.QueryRow("SELECT * FROM test WHERE id = ?", 1)
	}
}
