package models

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

//Company Entity
type Company struct {
	ID         bson.ObjectId `bson:"_id"`
	Name       string        `json:"name"`
	AddressZip string        `json:"zip"`
	Website    string        `json:"website"`
}

//Companies ...
type Companies []Company

//CompanyModel A model with Company API data
type CompanyModel struct{}

//GetCompanies Return a list of companies
func (m CompanyModel) GetCompanies() Companies {
	con := Connect()
	companies := con.Collection("companies")
	defer con.Close()

	results := Companies{}
	if err := companies.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to get Companies:", err)
	}

	return results
}

//AddCompany inserts a company in the DB
func (m CompanyModel) AddCompany(company Company) bool {
	con := Connect()
	companies := con.Collection("companies")
	defer con.Close()

	company.ID = bson.NewObjectId()
	err := companies.Insert(company)

	if err != nil {
		log.Fatal("Failed to insert company:", err)
		return false
	}

	return true
}
