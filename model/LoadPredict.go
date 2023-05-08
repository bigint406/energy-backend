package model

import (
	"encoding/json"
	"energy/defs"
	"energy/utils"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	//getWeatherForecastUrl = "https://api.qweather.com/v7/weather/168h?location=101010800&key=2dee7efdb9a54d06830b1c3af13857db"
	getWeatherForecastUrl = "http://11.10.21.201:7766/api/getData?index=predict"
	getDataUrl            = "http://11.10.21.201:7766/api/getData"
)

type Input struct {
	Date          [168]string  `json:"日期"`
	Temperature   [168]float64 `json:"温度"`
	Humidity      [168]int     `json:"湿度"`
	Radiation     [168]int     `json:"辐射"`
	Wind          [168]float64 `json:"风速"`
	RoomRate      [168]float64 `json:"在室率"`
	OccupancyRate [168]float64 `json:"入住率"`
	Load          [168]float64 `json:"负荷"`
}

type DataStruct struct {
	Data []float64
}

type Output struct {
	Result [168]float64 `json:"result"`
}

type Forecast struct {
	Data []Forecast2
}

type Forecast2 struct {
	Temperature string
	Humidity    string
	Wind        string
}

type Atmosphere struct {
	Time   int         `json:"time"`
	Result Atmosphere2 `json:"data"`
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

//当天凌晨跑
func LoadPredict(index string) Output {
	input := MakeInputBody(index)

	fmt.Println(string(input))

	data := Output{}
	resp, err := http.Post(utils.LoadPredictRouter+"/d1_groups", "application/json", strings.NewReader(string(input)))
	if err != nil {
		log.Println(err)
		return Output{}
	}
	defer resp.Body.Close()
	n, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(n, &data)
	fmt.Println("*")
	fmt.Println(data)
	return data
}

func MakeInputBody(index string) []byte {
	var input Input
	start := FindStart(int(time.Now().Unix()))
	for i := 0; i < 168; i++ {
		input.Date[i] = UnixToString(start + i*3600)
	}

	//前一天数据
	load := GetLoad(index, "yesterday")
	temperature := GetData("temperature", int(time.Now().Unix()-86400), index)
	humidity := GetData("humidity", int(time.Now().Unix()-86400), index)
	radiation := GetData("radiation", int(time.Now().Unix()-86400), index)
	wind := GetData("wind", int(time.Now().Unix()-86400), index)
	roomRate := GetData("roomRate", int(time.Now().Unix()-86400), index)
	occupancyRate := GetData("roomRate", int(time.Now().Unix()-86400), index)

	for i := 0; i < 24; i++ {
		input.Load[i] = load[i]
		input.Temperature[i] = temperature[i]
		input.Humidity[i] = int(humidity[i])
		input.Radiation[i] = int(radiation[i])
		input.Wind[i] = wind[i]
		input.RoomRate[i] = roomRate[i]
		input.OccupancyRate[i] = occupancyRate[i]
	}

	//后六天数据
	forecast := GetForecast()

	for i := 24; i < 168; i++ {
		input.Temperature[i], _ = strconv.ParseFloat(forecast.Data[i-24].Temperature, 64)
		input.Humidity[i], _ = strconv.Atoi(forecast.Data[i-24].Humidity)
		input.Wind[i], _ = strconv.ParseFloat(forecast.Data[i-24].Wind, 64)
		input.Load[i] = 0
		input.Radiation[i] = int(radiation[i%24])
		input.RoomRate[i] = roomRate[i%24]
		input.OccupancyRate[i] = occupancyRate[i%24]
	}

	output, _ := json.Marshal(&input)
	return output
}

// GetData 访问办公网数据库
// index:哪种类型的数据 ||base：unix时间戳 ||zutuan：哪个组团
func GetData(index string, base int, zutuan string) []float64 {
	data := DataStruct{}
	baseString := strconv.Itoa(base)
	resp, err := http.Get(getDataUrl + "?index=" + index + "&base=" + baseString + "&zutuan=" + zutuan)
	if err != nil {
		log.Println(err)
		return []float64{}
	}
	defer resp.Body.Close()
	n, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(n, &data)

	return data.Data
}

func GetLoad(index string, flag string) []float64 {
	var load []float64

	if flag == "today" {
		str := GetToday()
		//str := "2023/05/08"
		switch index {
		case "D1组团":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDay1, str)
		case "D2组团":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDay2, str)
		case "D3组团":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDay3, str)
		case "D4组团":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDay4, str)
		case "D5组团":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDay5, str)
		case "D6组团":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDay6, str)
		case "公共组团南区":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDayPubS, str)
		case "公共组团北区":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDayPubS, str)
		}
	} else if flag == "yesterday" {
		switch index {
		case "D1组团":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDay1, GetYesterday())
		case "D2组团":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDay2, GetYesterday())
		case "D3组团":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDay3, GetYesterday())
		case "D4组团":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDay4, GetYesterday())
		case "D5组团":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDay5, GetYesterday())
		case "D6组团":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDay6, GetYesterday())
		case "公共组团南区":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDayPubS, GetYesterday())
		case "公共组团北区":
			load, _ = GetResultFloatList(defs.GroupHeatConsumptionDayPubS, GetYesterday())
		}
	}

	if len(load) < 24 {
		var load2 []float64
		load2 = make([]float64, 24)
		for i := 0; i < len(load); i++ {
			load2[i] = load[i]
		}
		for i := len(load); i < 24; i++ {
			load2[i] = 0
		}
		return load2
	}

	return load
}

func GetForecast() Forecast {
	data := Forecast{}
	resp, err := http.Get(getWeatherForecastUrl)
	if err != nil {
		log.Println(err)
		return Forecast{}
	}
	defer resp.Body.Close()
	n, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(n, &data)
	return data
}

func UnixToString(unix int) string {
	timeLayout := "2006-01-02 15:04:05"
	timeStr := time.Unix(int64(unix), 0).Format(timeLayout)
	return timeStr
}

func FindStart(value int) int {
	Time := time.Unix(int64(value), 0)
	Time2 := time.Date(Time.Year(), Time.Month(), Time.Day(), 0, 0, 0, 0, Time.Location())
	return int(Time2.Unix())
}

func GetDay(v int64) string {
	timeLayout := "2006-01-02 15:04:05"
	timeStr := time.Unix(v, 0).Format(timeLayout)
	a := strings.Split(timeStr, " ")
	return a[0]
}

func GetToday() string {
	timeLayout := "2006-01-02 15:04:05"
	timeStr := time.Unix(time.Now().Unix(), 0).Format(timeLayout)
	a := strings.Split(timeStr, " ")
	a[0] = strings.Replace(a[0], "-", "/", 2)

	return a[0]
}

func GetYesterday() string {
	timeLayout := "2006-01-02 15:04:05"
	timeStr := time.Unix(time.Now().Unix()-86400, 0).Format(timeLayout)
	a := strings.Split(timeStr, " ")
	a[0] = strings.Replace(a[0], "-", "/", 2)

	return a[0]
}

func FindIntervalDay(value int) (int, int) {
	Time := time.Unix(int64(value), 0)
	Time1 := time.Date(Time.Year(), Time.Month(), Time.Day(), 0, 0, 0, 0, Time.Location())
	Time2 := time.Date(Time.Year(), Time.Month(), Time.Day(), 24, 0, 0, 0, Time.Location())
	return int(Time1.Unix()), int(Time2.Unix())
}
