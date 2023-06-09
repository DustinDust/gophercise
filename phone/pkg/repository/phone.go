package repository

import (
	"database/sql"
	"log"
)

type PhoneRepository struct {
	DbConn sql.DB
}

type Phone struct {
	ID    int    `sql:"id"`
	Value string `sql:"value"`
}

func (p *PhoneRepository) InsertPhone(phstr string) (int, error) {
	statement := "INSERT INTO phone_number(value) VALUES($1) RETURNING id"
	var id int
	err := p.DbConn.QueryRow(statement, phstr).Scan(&id)
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (p *PhoneRepository) GetPhone(id int) (*Phone, error) {
	var phone Phone
	statement := "SELECT * FROM phone_number WHERE id=$1"
	err := p.DbConn.QueryRow(statement, id).Scan(&phone.ID, &phone.Value)
	if err != nil {
		return nil, err
	}
	return &phone, nil
}

func (p *PhoneRepository) FindPhoneByValue(value string) (*Phone, error) {
	var phone Phone
	statement := "SELECT * FROM phone_number WHERE value=$1"
	err := p.DbConn.QueryRow(statement, value).Scan(&phone.ID, &phone.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &phone, nil
}

func (p *PhoneRepository) UpdatePhone(ph Phone) error {
	statement := `UPDATE phone_number SET value=$2 WHERE id=$1`
	_, err := p.DbConn.Exec(statement, ph.ID, ph.Value)
	return err
}

func (p *PhoneRepository) DeletePhone(id int) error {
	statement := `DELETE FROM phone_number WHERE id=$1`
	_, err := p.DbConn.Exec(statement, id)
	return err
}

func (p *PhoneRepository) GetAllPhones() ([]Phone, error) {
	statement := "SELECT * FROM phone_number"
	rows, err := p.DbConn.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	phones := make([]Phone, 0)
	for rows.Next() {
		var id int
		var value string
		if err := rows.Scan(&id, &value); err != nil {
			log.Printf("Error scanning rows: %v", err)
			break
		}
		phones = append(phones, Phone{ID: id, Value: value})
	}
	return phones, nil
}

func (p *PhoneRepository) Seed() {
	_, err := p.InsertPhone("1234567890")
	if err != nil {
		panic(err)
	}
	_, err = p.InsertPhone("123 456 7891")
	if err != nil {
		panic(err)
	}
	_, err = p.InsertPhone("(123) 456 7892")
	if err != nil {
		panic(err)
	}
	_, err = p.InsertPhone("(123) 456-7893")
	if err != nil {
		panic(err)
	}
	_, err = p.InsertPhone("123-456-7894")
	if err != nil {
		panic(err)
	}
	_, err = p.InsertPhone("123-456-7890")
	if err != nil {
		panic(err)
	}
	_, err = p.InsertPhone("1234567892")
	if err != nil {
		panic(err)
	}
	_, err = p.InsertPhone("(123)456-7892")
	if err != nil {
		panic(err)
	}
}
