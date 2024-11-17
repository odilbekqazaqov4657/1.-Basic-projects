package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

var ctx context.Context = context.Background()

func main() {
	db, err := Conn()
	if err != nil {
		return
	}

	defer db.Close(ctx)

	crud(db)
}

func crud(db *pgx.Conn) {

	// GetUsersById(db, 2)
	//GetUsers(db)
	//CreateUser(db)
	//UpdateUser(db)
	DeleteUser(db)
}

func Conn() (*pgx.Conn, error) {

	// psql-U postgres -h localhost -d postgres
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@0.0.0.0:5432/postgres")
	if err != nil {
		log.Println("error on connecting to database", err)
		return nil, err
	}
	log.Println("succassfully connected to db !!!")

	return conn, nil
}

func CreateUser(db *pgx.Conn) {

	query := `
			INSERT INTO
				uzers(
					username,
					gmail,
					age
				) VALUES (
					$1, $2, $3
					)
				`
	_, err := db.Exec(ctx, query, "Umar", "umar@01gmail.com", 23)
	if err != nil {
		log.Println("error on inserting data to db:", err)
		return
	}
	fmt.Println("Successfully inserted your data !!!")

	defer db.Close(ctx)

}

func GetUsersById(db *pgx.Conn, id int) {

	var (
		user_id  int
		username string
		gmail    string
		age      int
	)

	row := db.QueryRow(ctx, "SELECT * FROM uzers WHERE user_id=$2", id)

	row.Scan(
		&user_id,
		&username,
		&gmail,
		&age,
	)

	fmt.Println(user_id, username, gmail, age)

	defer db.Close(ctx)

}

func GetUsers(db *pgx.Conn) {
	query := `
			SELECT
				*
			FROM
				uzers
			`
	rows, err := db.Query(ctx, query)
	if err != nil {
		log.Println("error on getting all data from database", err.Error())
		return
	}

	for rows.Next() {
		var (
			user_id  int
			username string
			gmail    string
			age      int
		)
		err := rows.Scan(
			&user_id,
			&username,
			&gmail,
			&age,
		)

		if err != nil {
			log.Println("error on scanning data: ", err)
			return
		}
		fmt.Println(user_id, username, gmail, age)
	}

	defer rows.Close()
}

func UpdateUser(db *pgx.Conn) {

	query := `
			UPDATE 
				uzers
			SET
				username=$1
			WHERE
				user_id=$2
			`
	_, err := db.Exec(ctx, query, "Umarali", 5)
	if err != nil {
		log.Println("error on updating data:", err)
		return
	}
	fmt.Println("Successfully updated your data !!!")

	defer db.Close(ctx)

}

func DeleteUser(db *pgx.Conn) {

	query := `
			DELETE FROM
				uzers
			WHERE
				user_id=$1
				`
	_, err := db.Exec(ctx, query, 3)
	if err != nil {
		log.Println("error on deleting data:", err)
		return
	}
	fmt.Println("Successfully deleted your data !!!")

	defer db.Close(ctx)

}
