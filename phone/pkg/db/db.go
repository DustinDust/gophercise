package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123123"
	dbname   = "gophercise8"
)

type DBConnection struct {
	Connection *sql.DB
}

func NewConnection() *DBConnection {
	conn := &DBConnection{}
	var err error
	conn.Connection, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		panic(err)
	}
	err = conn.Connection.Ping()
	if err != nil {
		panic(err)
	}
	conn.createPhoneNumberTable()
	return conn
}

func (d *DBConnection) ResetDB(name string) error {
	_, err := d.Connection.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return d.CreateDB(name)
}

func (d *DBConnection) CreateDB(name string) error {
	_, err := d.Connection.Exec("CREATE DATABASE " + name + " template template0")
	if err != nil {
		return err
	}
	return nil
}

func (d *DBConnection) createPhoneNumberTable() error {
	statement := `CREATE TABLE IF NOT EXISTS phone_number (
		id SERIAL,
		value VARCHAR(255)
	)`
	_, err := d.Connection.Exec(statement)
	return err
}

