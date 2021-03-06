package main

import (
	"fmt"
	"net/http"
)


func main() {
	url := "http://api.openweathermap.org/data/2.5/find?appid=0a12b8f2f0dd011ed6085cb995ff61b4&lat=-37.81&lon=144.96&cnt=1"

	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting weather: %v\n", err)
		return
	}

	fmt.Printf("Our API url is %s\n", url)
}
