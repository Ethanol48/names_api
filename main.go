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
	prueba  string
}

type errorMessage struct {
	Message string
}

func isInArray(alpha []string, str string) bool {

	// iterate using the for loop
	for i := 0; i < len(alpha); i++ {
		// check
		if alpha[i] == str {
			// return true
			return true
		}
	}
	return false
}

func getFromCountry(r map[string]body, c string, g string) (name, error) {
	// c = is the Country string, Spain, Saudi Arabia, Germany ...
	// g = gender of desired name
	// r = response from getWholeJson
	rand.Seed(time.Now().UnixNano())

	switch g {
	case "Male", "male":
		vn := rand.Intn(len(r[c].Male) - 1)
		vsn := rand.Intn(len(r[c].Surnames) - 1)

		result := name{Name: r[c].Male[vn], Surname: r[c].Surnames[vsn]}

		return result, nil

	case "Female", "female":
		vn := rand.Intn(len(r[c].Female) - 1)
		vsn := rand.Intn(len(r[c].Surnames) - 1)

		result := name{Name: r[c].Female[vn], Surname: r[c].Surnames[vsn]}

		return result, nil

	default:
		return name{}, nil
	}
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

	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	c := vars["country"]
	g := vars["gender"]

	resultado := getWholeJson("names_new.json")

	keys := make([]string, 0, len(resultado))
	for k := range resultado {
		keys = append(keys, k)
	}

	checkcountry := isInArray(keys, c)
	if !checkcountry {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "Error, the country name wasn't right, did you Capitalize it??",
		})

	} else {
		result, _ := getFromCountry(getWholeJson("names_new.json"), c, g)

		empty := name{}
		if result == empty {

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorMessage{
				Message: "Error, the gender wasn't right, the genders are 'Male' or 'Female'",
			})

		} else {

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(result)
		}
	}
}

func randomGender(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	resultado := getWholeJson("names_new.json")

	keys := make([]string, 0, len(resultado))
	for k := range resultado {
		keys = append(keys, k)
	}

	vars := mux.Vars(r)
	gdr := vars["gender"]

	// random number to select the country
	vc := rand.Intn(len(keys) - 1)

	var nm string

	if gdr == "Male" || gdr == "male" {
		vname := rand.Intn(len(resultado[keys[vc]].Male) - 1)
		nm = resultado[keys[vc]].Male[vname]

		// random number to select the country
		vcnm := rand.Intn(len(keys) - 1)

		// random number to select a Surname
		vsnm := rand.Intn(len(resultado[keys[vcnm]].Surnames) - 1)
		snm := resultado[keys[vcnm]].Surnames[vsnm]

		result := name{Name: nm, Surname: snm, prueba: gdr}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)

	} else if gdr == "Female" || gdr == "female" {
		vname := rand.Intn(len(resultado[keys[vc]].Female) - 1)
		nm = resultado[keys[vc]].Female[vname]

		// random number to select the country
		vcnm := rand.Intn(len(keys) - 1)

		// random number to select a Surname
		vsnm := rand.Intn(len(resultado[keys[vcnm]].Surnames) - 1)
		snm := resultado[keys[vcnm]].Surnames[vsnm]

		result := name{Name: nm, Surname: snm, prueba: gdr}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)

	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorMessage{
			Message: "Error, the gender wasn't right, the genders are 'Male' or 'Female'",
		})
	}

}

func randomName(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

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

	r := mux.NewRouter()

	r.HandleFunc("/name/{country}/{gender}", getNameFromCountry)
	r.HandleFunc("/random/{gender}", randomGender)
	r.HandleFunc("/random", randomName)
	http.ListenAndServe(":4000", r)

}
