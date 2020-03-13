package controllers

import (
	"app/src/models"
	"app/src/util"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func isNameValid(name string) bool {
	match, _ := regexp.MatchString("^[A-Z ]*$", name)
	return match
}

func isZipValid(zip string) bool {
	match, _ := regexp.MatchString("^[0-9]{5}$", zip)
	return match
}

func isWebsiteValid(website string) bool {
	var re = regexp.MustCompile(`(?m)^([-a-z0-9@:%._\+~#=]{1,256}\.[a-z0-9()]{1,6}\b([-a-z0-9()@:%_\+.~#?&//=]*))?$`)
	match := re.MatchString(website)
	return match
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

	if !isNameValid(company.Name) {
		util.RespondError(w, 422, "Company name must be Uppercase")
		return
	}

	if !isZipValid(company.AddressZip) {
		util.RespondError(w, 422, "Company zip must contain five digits")
		return
	}

	if !isWebsiteValid(company.Website) {
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
	file, fileHandler, err := r.FormFile("csv")
	if err != nil {
		log.Fatalln("Error MergeCompany", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	f, err := os.OpenFile(fileHandler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln("Error MergeCompany", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	io.Copy(f, file)
	uploadedFile, _ := os.Open(fileHandler.Filename)
	reader := csv.NewReader(bufio.NewReader(uploadedFile))
	reader.Comma = ';'
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatalln("Error MergeCompany", err)
		w.WriteHeader(http.StatusBadRequest)
		//w.Write
		return
	}
	if len(records) == 0 {
		log.Fatalln("Empty file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, record := range records {
		if i == 0 {
			continue
		}
		company := models.Company{Name: record[0], AddressZip: record[1], Website: record[2]}

		retrievedCompany := c.model.FindByName(company.Name)
		if retrievedCompany == nil {
			continue
		}

		retrievedCompany.Website = company.Website
		c.model.UpdateCompany(*retrievedCompany)
	}

	w.WriteHeader(http.StatusOK)
	return
}

//LoadCompanies load companies from a csv file
func (c *CompanyController) LoadCompanies(filename string) {
	file, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = ';'
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatalln("Error importing companies ", err)
		return
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		company := models.Company{Name: record[0], AddressZip: record[1], Website: ""}

		if !isNameValid(company.Name) || !isZipValid(company.AddressZip) {
			continue
		}

		c.model.AddCompany(company)
	}
}
