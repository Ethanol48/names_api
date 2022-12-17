package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type body struct {
	Male     []string
	Female   []string
	Surnames []string
}

type name struct {
	Name    string
	Surname string
}

func getFromCountry(r map[string]body, c string, g string) (name, error) {
	// c = is the Country string, Spain, Saudi Arabia, Germany ...
	// g = gender of desired name
	// r = response from getWholeJson

	switch g {
	case "Male":
		vn := rand.Intn(len(r[c].Male) - 1)
		vsn := rand.Intn(len(r[c].Surnames) - 1)

		result := name{Name: r[c].Male[vn], Surname: r[c].Surnames[vsn]}

		return result, nil

	case "Female":
		vn := rand.Intn(len(r[c].Female) - 1)
		vsn := rand.Intn(len(r[c].Surnames) - 1)

		result := name{Name: r[c].Female[vn], Surname: r[c].Surnames[vsn]}

		return result, nil
	}

	return name{}, nil
}

func getWholeJson(path string) map[string]body {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Error during reading json file: ", err)
	}

	var payload map[string]body

	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload
}

func getNameFromCountry(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	c := vars["country"]
	g := vars["gender"]

	result, _ := getFromCountry(getWholeJson("names_new.json"), c, g)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func randomName(w http.ResponseWriter, r *http.Request) {

	resultado := getWholeJson("names_new.json")

	keys := make([]string, 0, len(resultado))
	for k := range resultado {
		keys = append(keys, k)
	}

	var gdr string

	tmp := rand.Intn(3)
	if tmp >= 2 {
		gdr = "Male"
	} else {
		gdr = "Female"
	}
	// random number to select the country
	vc := rand.Intn(len(keys) - 1)

	var nm string
	if gdr == "Male" {
		vname := rand.Intn(len(resultado[keys[vc]].Male) - 1)
		nm = resultado[keys[vc]].Male[vname]

	} else if gdr == "Female" {
		vname := rand.Intn(len(resultado[keys[vc]].Female) - 1)
		nm = resultado[keys[vc]].Female[vname]
	}

	// random number to select the country
	vcnm := rand.Intn(len(keys) - 1)

	// random number to select a Surname
	vsnm := rand.Intn(len(resultado[keys[vcnm]].Surnames) - 1)
	snm := resultado[keys[vcnm]].Surnames[vsnm]

	result := name{Name: nm, Surname: snm}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {

	rand.Seed(time.Now().UnixNano())

	r := mux.NewRouter()

	r.HandleFunc("/name/{country}/{gender}", getNameFromCountry)
	r.HandleFunc("/random", randomName)
	http.ListenAndServe(":4000", r)

}
