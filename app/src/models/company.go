package models

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

//Company Entity
type Company struct {
	ID         bson.ObjectId `bson:"_id"`
	Name       string        `json:"name"`
	AddressZip string        `json:"zipCode"`
	Website    string        `json:"website"`
}

//Companies ...
type Companies []Company

//CompanyModel A model with Company API data
type CompanyModel struct{}

func isNameValid(name string) bool {
	match, _ := regexp.MatchString("^[A-Z&' ]*$", name)
	return match
}

func isZipValid(zip string) bool {
	match, _ := regexp.MatchString("^[0-9]{5}$", zip)
	return match
}

func isWebsiteValid(website string) bool {
	var re = regexp.MustCompile(`(?mi)^((http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/)[a-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?)?$`)
	match := re.MatchString(website)
	return match
}

//GetCompanies Return a list of companies
func (m CompanyModel) GetCompanies() Companies {
	con := Connect()
	if con.db == nil {
		return Companies{}
	}

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
	if con.db == nil {
		return false
	}

	company.Name = strings.ToUpper(company.Name)
	if !isNameValid(company.Name) || !isZipValid(company.AddressZip) || !isWebsiteValid(company.Website) {
		return false
	}

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

//UpdateCompany updates one company already in db
func (m CompanyModel) UpdateCompany(company Company) bool {
	con := Connect()
	if con.db == nil {
		return false
	}

	company.Name = strings.ToUpper(company.Name)
	if !isNameValid(company.Name) || !isZipValid(company.AddressZip) || !isWebsiteValid(company.Website) {
		return false
	}

	companies := con.Collection("companies")
	defer con.Close()

	err := companies.UpdateId(company.ID, company)

	if err != nil {
		log.Fatal("Failed to update company:", err)
		return false
	}

	return true
}

//FindByName returns company if found or nil
func (m CompanyModel) FindByName(name string) *Company {
	con := Connect()
	if con.db == nil {
		return nil
	}

	companies := con.Collection("companies")
	defer con.Close()

	result := Company{}
	err := companies.Find(bson.M{"name": strings.ToUpper(name)}).One(&result)

	if err != nil {
		return nil
	}

	return &result
}

//FindByNameAndZip returns company if found or nil
func (m CompanyModel) FindByNameAndZip(name, zip string) *Company {
	con := Connect()
	if con.db == nil {
		return nil
	}

	companies := con.Collection("companies")
	defer con.Close()

	result := Company{}
	search := bson.RegEx{Pattern: strings.ToUpper(name) + ".*", Options: ""}
	err := companies.Find(bson.M{"name": search, "addresszip": zip}).One(&result)

	if err != nil {
		return nil
	}

	return &result
}
