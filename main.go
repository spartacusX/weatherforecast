package main

import (
	"fmt"
	. "github.com/spartacusX/weatherforecast/parser"
	. "github.com/spartacusX/weatherforecast/queryweather"
	"html/template"
	"log"
	"net/http"
	"time"
)

type TodayWeather struct {
	Today       time.Time
	WeatherInfo WeatherDetail
}

var ccMap *CityCodeMap

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method: ", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("static/login.html")
		if err == nil {
			t.Execute(w, nil)
		} else {
			log.Fatal(err)
		}
	} else {
		r.ParseForm()
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])

		//go to city selection page
		SelectCity(w, r)
	}
}

func SelectCity(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("static/selectcity.html")
	if err == nil {
		t.Execute(w, nil)
	} else {
		log.Fatal(err)
	}
}

func QueryWeather(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method: ", r.Method)
	r.ParseForm()
	weather := new(Weather)
	strCode := GetCityCode(r.FormValue("province"), r.FormValue("city"))
	GetWeatherByWeb(strCode, weather)

	tw := TodayWeather{time.Now(), weather.Weatherinfo}

	t, err := template.ParseFiles("static/weatherinfo.html")
	if err == nil {
		t.Execute(w, tw)
	} else {
		log.Fatal(err)
	}
}

func ParseCities() {
	ccMap = Parsecity("cities.txt")
}

func GetCityCode(strProvince string, strCity string) string {
	return ccMap.GetCode(strCity)
}

func main() {
	ParseCities()

	http.HandleFunc("/login", login)
	http.HandleFunc("/selectcity", SelectCity)
	http.HandleFunc("/queryweather", QueryWeather)

	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}
}
