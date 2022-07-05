package main

import (
	"fmt"

	"github.com/go-packages/countries"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("Message")
	myUuid := uuid.New()
	fmt.Println(myUuid)
	countries.AddCountry("Colombia")
	countries.AddCountry("Argentina")
	countries.AddCountry("Brazil")
	countries.AddCountry("Japan")

	countries.SetMyCountry("Colombia")

	countries.GetCountries()

}
