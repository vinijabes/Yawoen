package controllers

import (
	"app/src/models"
	"app/src/util"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

//CompanyController ...
type CompanyController struct {
	Controller,
	model models.CompanyModel
}

//NewCompanyController is a CompanyController constructor
func NewCompanyController() *CompanyController {
	c := new(CompanyController)
	c.model = models.CompanyModel{}
	return c
}

//GetCompanies GET /v1/company application/json
func (c *CompanyController) GetCompanies(w http.ResponseWriter, r *http.Request) {
	companies := c.model.GetCompanies()
	util.RespondJSON(w, http.StatusOK, companies)
	return
}

//GetCompanyByNameAndZip GET /v1/company?name={value}&zip={value} application/json
func (c *CompanyController) GetCompanyByNameAndZip(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	zip := r.URL.Query().Get("zip")

	companie := c.model.FindByNameAndZip(name, zip)
	util.RespondJSON(w, http.StatusOK, companie)

	return
}

//CreateCompany POST /v1/company application/json
func (c *CompanyController) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 128*1024*8)) //128kb

	if err != nil {
		log.Fatal("Failed to add company")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatal("Failed to add company")
	}

	if err := json.Unmarshal(body, &company); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddCompany unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	match, _ := regexp.MatchString("^[A-Z]*$", company.Name)
	if !match {
		util.RespondError(w, 422, "Company name must be Uppercase")
		return
	}

	match, _ = regexp.MatchString("^[0-9]{5}$", company.AddressZip)
	if !match {
		util.RespondError(w, 422, "Company zip must contain five digits")
		return
	}

	var re = regexp.MustCompile(`(?m)^[-a-z0-9@:%._\+~#=]{1,256}\.[a-z0-9()]{1,6}\b([-a-z0-9()@:%_\+.~#?&//=]*)$`)
	match = re.MatchString(company.Website)
	if !match {
		util.RespondError(w, 422, "Company website must be an lower case url.(e.g www.example.com)")
		return
	}

	result := c.model.AddCompany(company)

	if !result {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	return
}

//MergeCompanies POST /v1/company multipart/form-data
func (c *CompanyController) MergeCompanies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/MergeCompanies"))
	return
}
