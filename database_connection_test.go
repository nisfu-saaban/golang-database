package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestDatabaseConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}

func TestExecuteQuery(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id,name) VALUES('1','Abdur')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Data to Database")
}

func TestQueryContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id,name FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("id", id)
		fmt.Println("name", name)
	}

}

func TestQueryComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id,name,email,balance,rating,birth_date,married,created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("id", id, "name", name, "email", email, "balance", balance, "rating", rating, "birth date", birthDate, "married", married, "create at", createdAt)

	}
}
