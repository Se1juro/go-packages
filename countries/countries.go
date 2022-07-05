package countries

import (
	"errors"
	"fmt"
)

var myCountry string = " "
var errorNotFound error = errors.New("Country not found")
var countries []string

func New() {

}

func AddCountry(countryName string) {
	countries = append(countries, countryName)
}

func GetCountries() {
	for i, name := range countries {
		myCountryMsg := ""
		if name == myCountry {
			myCountryMsg = "[My country]"
		}
		fmt.Printf("%d. %s %s\n", i+1, name, myCountryMsg)
	}
}

func SetMyCountry(country string) error {
	if !isInCollection(country) {
		return errorNotFound
	}
	myCountry = country
	return nil
}

func isInCollection(country string) bool {
	for _, c := range countries {
		if c == country {
			return true
		}
	}
	return false
}
