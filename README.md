# Data-Integration-Api

The objective of this application is to expose a RESTful API to perform operation over companies data.

## Stack
- Go
- MongoDB
- Docker

## Endpoints
After start up, the API will be avaible listening port 5000 for following endpoints.

| Name | Path | Method | Content-Type | Description |
| ------ | ------ | ------ | ------ | ------ |
| List all companies| /v1/companies | GET | application/json | Retrieve all companies stored in the database. |
| Search company by name and zip | /v1/companies/search?name={value}&zip={value} | GET | application/json | Provides information based on query parameters values. Parameter name can be part of the company's name but zip needs to be the entire zip code of the company. Both parameters must be send in the request. |
| Create company | /v1/companies | POST | application/json | Create a new company. See example [here](#post-v1companies)  |
| Merge companies with CSV | /v1/companies/merge | POST | multipart/form-data | Parses a valid CSV file and integrate its data with the existent records(Append website on existing records). If the record doesn't exist, it will be discarded. The key of the file must be named "csv". See example [here](#post-v1companiesmerge) |

## Request & Response Examples

### API Resources
- [GET /companies](#get-v1companies)
- [GET /companies/search](#get-v1companiessearch)
- [POST /companies](#post-v1companies)
- [POST /companies/merge](#post-v1companiesmerge)

### GET /v1/companies

Response body:

    [{
        "ID": "5e6ab36fe5574a0006e920e7",
        "name":"TOLA SALES GROUP",
        "zip":"78229",
        "website":"http://repsources.com"
    }, ...]

### GET /v1/companies/search

Example: /v1/companies/search?name=TOLA&zip=78229

Response body:

    {
        "ID":"5e6ab36fe5574a0006e920e7",
        "name":"TOLA SALES GROUP",
        "zip":"78229",
        "website":"http://repsources.com"
    }

### POST /v1/companies

Request body:

    {
        "name": "TOLA SALES GROUP",
        "zipCode": "78229"        
    }

### POST /v1/companies/merge

CSV format:
    
| Name | Address Zip | Website |
| ------ | ------ | ------ |
| TOLA SALES GROUP | 78229 | http://repsources.com |



## Setup

First, you need to have docker and docker-compose installed. The instructions can be found [here](https://docs.docker.com/install/)

## Container

To run the application execute:

```sh
docker-compose up -d
```
On first time the application will load data in **q1_catalog.csv**

### Tests

To perform tests with go, run:

```sh
cd app
go test ./tests
```