package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
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
		log.Fatal("Error during Unmarshal(): ", err)
	}

	var payload map[string]body

	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload
}

func main() {

	rand.Seed(time.Now().UnixNano())

	resultado := getWholeJson("names_new.json")
	fmt.Println(getFromCountry(resultado, "Germany", "Male"))

	// fmt.Println(resultado["Spain"].Male)

	// fmt.Println(payload["Spain"])

}
