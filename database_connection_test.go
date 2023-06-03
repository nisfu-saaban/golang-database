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

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "admin'; #"
	password := "salah"

	ctx := context.Background()
	query := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, query)
	if rows.Next() {
		var username string
		rows.Scan(username)
		fmt.Println("sukses login", username)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Login gagal")
	}

	defer rows.Close()
}

func TestQueryParams(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "admin"
	password := "admin"

	ctx := context.Background()
	query := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, query, username, password)
	if rows.Next() {
		var username string
		rows.Scan(username)
		fmt.Println("sukses login", username)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("Login gagal")
	}
}

func TestQueryInsertSafe(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	username := "Doa'; DROP TABLE; #"
	password := "Sumayah"
	ctx := context.Background()

	query := "INSERT INTO user(username,password) VALUES(?,?)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert Data to Database")
}
