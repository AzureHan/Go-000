package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
)

type user struct {
	ID   uint
	Name string
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:secret@tcp(127.0.0.1:3306)/go")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	apiGetUser()
}

func apiGetUser() {
	usr, err := svcGetUser(0)
	if err != nil {
		if errors.Is(xerrors.Cause(err), sql.ErrNoRows) {
			fmt.Println("404")
			return
		}
		log.Printf("Error: %+v\n", err)
		fmt.Println("500")
		return
	}
	fmt.Println("200", usr.Name)
}

func svcGetUser(id uint) (*user, error) {
	return daoGetUser(id)
}

func daoGetUser(id uint) (*user, error) {
	var usr user
	err := db.QueryRow("select id, name from user where id = ?", id).Scan(&usr.ID, &usr.Name)
	if err == sql.ErrNoRows {
		// ErrNoRows is returned by Scan when QueryRow doesn't return a row.
		return nil, xerrors.Wrapf(err, "user %d not found", id)
	}
	if err != nil {
		return nil, xerrors.Wrap(err, "get user err")
	}
	return &usr, nil
}
