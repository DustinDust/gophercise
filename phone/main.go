package main

import (
	"fmt"
	"log"
	"phone/pkg/db"
	"phone/pkg/normalizer"
	"phone/pkg/repository"
)

func main() {
	dbConn := db.NewConnection()
	phoneRepo := repository.PhoneRepository{
		DbConn: *dbConn.Connection,
	}
	p, err := phoneRepo.GetAllPhones()
	if err != nil {
		log.Println(err)
	} else {
		for _, row := range p {
			normed := normalizer.NormalizeRegex(row.Value)
			if normed != row.Value {
				fmt.Println("Updating or removing ", normed)
				exisiting, err := phoneRepo.FindPhoneByValue(normed)
				if err != nil {
					panic(err)
				}
				if exisiting != nil {
					phoneRepo.DeletePhone(row.ID)
				} else {
					row.Value = normed
					phoneRepo.UpdatePhone(row)
				}
			} else {
				fmt.Println("No changes required.")
			}
		}
	}
	defer dbConn.Connection.Close()
}
