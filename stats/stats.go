package stats

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"example.com/covid19/notifications"
)

// CovidAPIResponse stores the data return from the API
type CovidAPIResponse struct {
	Country            string `json:"country,omitempty"`
	Cases              int    `json:"cases"`
	TodayCases         int    `json:"todayCases,omitempty"`
	Deaths             int    `json:"deaths"`
	TodayDeaths        int    `json:"todayDeaths,omitempty"`
	Recovered          int    `json:"recovered"`
	Active             int    `json:"active,omitempty"`
	Critical           int    `json:"critical,omitempty"`
	CasesPerOneMillion int    `json:"casesPerOneMillion,omitempty"`
}

// APIURL stores the API Endpoint to be hit for the stats overall
const APIURL = "https://corona.lmao.ninja/all"

// CountryAPIURL stores the API Endpoint to be hit for the stats per country
const CountryAPIURL = "https://corona.lmao.ninja/countries/"

var totalCases, deaths, recovered int
var country string

// Init initializes the stats
func Init(c string) {
	country = c
	covResp, err := getCovid19Stats()
	if err != nil {
		log.Printf("Error while getting the stats. Error: [%#v]", err)
		return
	}
	totalCases = covResp.Cases
	deaths = covResp.Deaths
	recovered = covResp.Recovered
	notifications.SendNotification(
		"Current Stats are:",
		"Total Cases:"+strconv.Itoa(totalCases)+" Total Casualities: "+strconv.Itoa(deaths)+" Total Recovered: "+strconv.Itoa(recovered),
		"WARNING",
	)
}

// StartCovid19Stats starts the process
func StartCovid19Stats() {
	for {
		covResp, err := getCovid19Stats()
		if err != nil {
			log.Printf("Error while getting the stats. Error: [%#v]", err)
			continue
		}
		if (covResp.Cases - totalCases) > 0 {
			// Send Notification
			notifications.SendNotification(
				strconv.Itoa(covResp.Cases-totalCases)+" more confirmed cases",
				"Total count has reached to "+strconv.Itoa(covResp.Cases)+" in "+country,
				"WARNING",
			)
			totalCases = covResp.Cases
			time.Sleep(5 * time.Second)
		}
		if (covResp.Deaths - deaths) > 0 {
			// Send Notification
			notifications.SendNotification(
				strconv.Itoa(covResp.Deaths-deaths)+" more death case because of Covid19",
				"Total death count has reached to "+strconv.Itoa(covResp.Deaths)+" in "+country,
				"CRITICAL",
			)
			deaths = covResp.Deaths
			time.Sleep(5 * time.Second)
		}
		if (covResp.Recovered - recovered) > 0 {
			// Send Notification
			notifications.SendNotification(
				strconv.Itoa(covResp.Recovered-recovered)+" patient recovered from Covid19",
				"Total recovered cases have reached to "+strconv.Itoa(covResp.Recovered)+" in India",
				"WARNING",
			)
			recovered = covResp.Recovered
			time.Sleep(5 * time.Second)
		}
		time.Sleep(1 * time.Minute)
	}
}

func getCovid19Stats() (covResp CovidAPIResponse, err error) {
	resp, err := http.Get(CountryAPIURL + country)
	if strings.ToLower(country) == "all" {
		resp, err = http.Get(APIURL)
	}
	if err != nil {
		log.Printf("Error while getting the data from the API. Error: [%#v]", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error while reading the body from the response. Error: [%#v]", err)
		return
	}
	json.Unmarshal(respBody, &covResp)
	log.Printf("Got API Response. Response: [%#v]", covResp)
	return
}
