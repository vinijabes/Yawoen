package routes

import (
	"app/src/controllers"
	"net/http"
)

//Route Object
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var companyController = controllers.NewCompanyController()

//Routes - Route Array
type Routes []Route

var routes = Routes{
	Route{
		"GetCompanies",
		"GET",
		"/v1/companies",
		companyController.GetCompanies,
	},
	Route{
		"SearchCompany",
		"GET",
		"/v1/companies/search",
		companyController.GetCompanyByNameAndZip,
	},
	Route{
		"CreateCompany",
		"POST",
		"/v1/companies",
		companyController.CreateCompany,
	},
	Route{
		"MergeCompany",
		"POST",
		"/v1/companies/merge",
		companyController.MergeCompanies,
	},
}

// Route{
// 	"Index",
// 	"GET",
// 	"/v1/",
// 	teste,
// },
