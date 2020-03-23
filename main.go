package main

import (
	"flag"
	"log"

	"example.com/covid19/stats"
)

func main() {
	var country = flag.String("country", "India", "Country Name for Covid19 Notification")
	flag.Parse()
	log.Printf("Starting the application with country: [%s]", *country)
	stats.Init(*country)
	stats.StartCovid19Stats()
	return
}
