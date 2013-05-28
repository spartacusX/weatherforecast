package cityparser

import (
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
)

type CityCodeMap map[string]string

type city struct {
	Name string `json:"市名"`
	Code string `json:"编码"`
}

type province struct {
	Name   string `json:"省"`
	Cities []city `json:"市"`
}

type CityCode struct {
	AllProvinces []province `json:"城市代码"`
}

func (ccMap *CityCodeMap) GetCode(cityName string) (code string) {
	return (*ccMap)[cityName]
}

// func (cc *CityCode) GetCode(provinceName, cityName string) (code string) {
// 	for _, province := range cc.AllCities {
// 		if province.Name == provinceName {
// 			for _, city := range province.Cities {
// 				if city.Name == cityName {
// 					return city.Code
// 				}
// 			}
// 		}
// 	}
// 	return
// }

// func ParseJson(data interface{}, cc city_code_map) {
// 	switch data.(type) {
// 	case string:
// 		//fmt.Println(data, "is a string")
// 	case int:
// 		//fmt.Println(data, "is a int")
// 	case []interface{}:
// 		array := data.([]interface{})
// 		for _, val := range array {
// 			ParseJson(val, cc)
// 		}
// 	case map[string]interface{}:
// 		mArray := data.(map[string]interface{})
// 		if len(mArray) == 2 && mArray["市名"] != nil {
// 			name := mArray["市名"]
// 			code := mArray["编码"]
// 			//if ok1 && ok2 {
// 			cc[name.(string)] = code.(string)
// 			//}
// 		} else {
// 			for _, val := range mArray {
// 				ParseJson(val, cc)
// 			}
// 		}
// 	default:
// 		fmt.Println(data, "is a type I don't know how to handle")
// 	}
// }

func Parsecity(datafile string) (ccMap *CityCodeMap) {
	buf, err := ioutil.ReadFile(datafile)
	if err != nil {
		log.Fatal("Reading data failed: ", err)
		return
	}

	var cc = new(CityCode)
	err = json.Unmarshal(buf, cc)
	if err != nil {
		log.Fatal("Unmarshal json data failed: ", err)
		return
	}

	var citycode = make(CityCodeMap)
	for _, p := range cc.AllProvinces {
		for _, c := range p.Cities {
			citycode[c.Name] = c.Code
		}
	}
	return &citycode
}
