package basicMap

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	//getWeatherForecastUrl = "https://api.qweather.com/v7/weather/168h?location=101010800&key=2dee7efdb9a54d06830b1c3af13857db"
	getAtmosphereUrl = "http://11.10.21.201:7766/api/getData?index=atmo"
	getKekongUrl     = "http://11.10.21.201:7766/api/getData?index=kekong"
)

type Atmosphere struct {
	Time int         `json:"time"`
	Data Atmosphere2 `json:"data"`
}

type Atmosphere2 struct {
	Wind        Atmosphere3
	Temperature Atmosphere3
	Humidity    Atmosphere3
	Radiation   Atmosphere3
}

type Atmosphere3 struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Kekong struct {
	Time int     `json:"time"`
	D4   float64 `json:"d4"`
	D5   float64 `json:"d5"`
	D6   float64 `json:"d6"`
}

func GetAtmosphere(c *gin.Context) {
	/*
		c.JSON(http.StatusOK, gin.H{
			"风速":   "2.4",
			"湿度":   "30",
			"温度":   "5",
			"总辐射":  "450",
			"大气压力": "920",
		})

	*/

	c.JSON(http.StatusOK, gin.H{
		//"data": []string{"5", "30", "450", "2.4", "920"},
		"data": GetAtmo(),
	})

}

func GetKekong(c *gin.Context) {

	/*
		c.JSON(http.StatusOK, gin.H{
			"D1": 80,
			"D2": 60,
			"D3": 90,
			"D4": 70,
			"D5": 80,
			"D6": 50,
		})

	*/
	c.JSON(http.StatusOK, gin.H{
		//"data": []int{80, 60, 90, 70, 80, 50},
		"data": GetKK(),
	})
}

func GetAtmo() []string {
	data := Atmosphere{}
	resp, err := http.Get(getAtmosphereUrl)
	if err != nil {
		log.Println(err)
		return []string{}
	}
	defer resp.Body.Close()
	n, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(n, &data)

	return []string{data.Data.Temperature.Value, data.Data.Humidity.Value, data.Data.Radiation.Value, data.Data.Wind.Value, "920"}
}

func GetKK() []int {
	data := Kekong{}
	resp, err := http.Get(getKekongUrl)
	if err != nil {
		log.Println(err)
		return []int{}
	}
	defer resp.Body.Close()
	n, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(n, &data)

	return []int{int(data.D4 * 100), int(data.D5 * 100), int(data.D6 * 100), int(data.D4 * 100), int(data.D5 * 100), int(data.D6 * 100)}
}
