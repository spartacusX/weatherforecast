package queryweather

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type WeatherDetail struct {
	City     string
	Temp1    string
	Temp2    string
	Temp3    string
	Temp4    string
	Temp5    string
	Temp6    string
	Weather1 string
	Weather2 string
	Weather3 string
	Weather4 string
	Weather5 string
	Weather6 string
	Wind1    string
	Wind2    string
	Wind3    string
	Wind4    string
	Wind5    string
	Wind6    string
}

type Weather struct {
	Weatherinfo WeatherDetail
}

const PROXYURL = "http://web-proxy.rose.hp.com:8080"
const BASEURL = "http://m.weather.com.cn/data/"

func GetWeatherByWeb(cityCode string, weather *Weather) {
	strURL := BASEURL + cityCode + ".html"
	//fmt.Println(url)
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(PROXYURL)
	}

	transport := &http.Transport{Proxy: proxy}

	client := &http.Client{Transport: transport}

	resp, err := client.Get(strURL)
	if err != nil {
		log.Fatal("Get wether data failed: ", err)
		return
	}

	defer resp.Body.Close() //close when exit 

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Reading response body failed: ", err)
		return
	}

	err = json.Unmarshal(body, weather)
	if err != nil {
		log.Fatal("Unmarshal json data failed: ", err)
		return
	}

	//fmt.Printf("%+v\n", f.Weatherinfo)
}
