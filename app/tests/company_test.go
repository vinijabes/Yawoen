package tests

import (
	"testing"

	"app/src/models"
)

func TestAddCompany(t *testing.T) {
	models.DatabaseHost = "localhost:27017"
	testCompany := models.Company{}

	testCompany.Name = "Company teste"
	testCompany.AddressZip = "12345"
	testCompany.Website = "https://teste.com"

	model := models.CompanyModel{}
	if !model.AddCompany(testCompany) {
		t.Errorf("Failed to add company")
	}
}

func TestSearchCompanyByNameAndZip(t *testing.T) {
	models.DatabaseHost = "localhost:27017"

	model := models.CompanyModel{}
	companie := model.FindByNameAndZip("Comp", "12345")
	if companie == nil {
		t.Errorf("Failed to search company by name and zip")
	}
}

func TestSearchCompanyByName(t *testing.T) {
	models.DatabaseHost = "localhost:27017"

	model := models.CompanyModel{}
	companie := model.FindByName("Company teste")
	if companie == nil {
		t.Errorf("Failed to search company by name and zip")
	}
}

func TestUpdateCompany(t *testing.T) {
	models.DatabaseHost = "localhost:27017"

	model := models.CompanyModel{}
	companie := model.FindByName("Company teste")
	companie.Name = "Company Alterada"
	if !model.UpdateCompany(*companie) {
		t.Errorf("Failed to update company")
		return
	}

	result := model.FindByName("Company Alterada")

	if result == nil {
		t.Errorf("Failed to update company")
		return
	}
}