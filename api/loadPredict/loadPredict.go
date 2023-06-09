package loadPredict

import (
	"encoding/json"
	"energy/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"golang.org/x/sync/singleflight"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type LoadStatistic struct {
	Load []float64
	Temp []float64
	Xray []int
}

type LoadComparison struct {
	Real    []float64
	Predict []float64
	Xray    []int
}

func GetLoadStatistic(c *gin.Context) {

	index := c.Query("index")
	var a, b []float64
	var x []int
	var val string
	var err error

	switch index {
	case "D1组团":
		val, err = model.RedisClient.Get("loadStatistic_1").Result()
	case "D2组团":
		val, err = model.RedisClient.Get("loadStatistic_2").Result()
	case "D3组团":
		val, err = model.RedisClient.Get("loadStatistic_3").Result()
	case "D4组团":
		val, err = model.RedisClient.Get("loadStatistic_4").Result()
	case "D5组团":
		val, err = model.RedisClient.Get("loadStatistic_5").Result()
	case "D6组团":
		val, err = model.RedisClient.Get("loadStatistic_6").Result()
	case "公共组团南区":
		val, err = model.RedisClient.Get("loadStatistic_7").Result()
	case "公共组团北区":
		val, err = model.RedisClient.Get("loadStatistic_8").Result()
	}
	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			a = model.GetLoad(index, "today")
			b = model.GetData("temperature", int(time.Now().Unix()), "")
			x = make([]int, 24)
			for i := 0; i < 24; i++ {
				x[i] = i
			}
			for i := 0; i < len(a); i++ {
				a[i], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", a[i]), 64)
			}

			c.JSON(http.StatusOK, gin.H{
				"x轴":  x,
				"负荷量": a,
				"温度":  b,
			})

			loadStaticStruct := LoadStatistic{}
			loadStaticStruct.Load = a
			loadStaticStruct.Temp = b
			loadStaticStruct.Xray = x

			result, _ := json.Marshal(&loadStaticStruct)

			switch index {
			case "D1组团":
				model.RedisClient.Set("loadStatistic_1", result, 10*time.Minute)
			case "D2组团":
				model.RedisClient.Set("loadStatistic_2", result, 10*time.Minute)
			case "D3组团":
				model.RedisClient.Set("loadStatistic_3", result, 10*time.Minute)
			case "D4组团":
				model.RedisClient.Set("loadStatistic_4", result, 10*time.Minute)
			case "D5组团":
				model.RedisClient.Set("loadStatistic_5", result, 10*time.Minute)
			case "D6组团":
				model.RedisClient.Set("loadStatistic_6", result, 10*time.Minute)
			case "公共组团南区":
				model.RedisClient.Set("loadStatistic_7", result, 10*time.Minute)
			case "公共组团北区":
				model.RedisClient.Set("loadStatistic_8", result, 10*time.Minute)
			}
		}
	} else {
		data := LoadStatistic{}
		_ = json.Unmarshal([]byte(val), &data)

		c.JSON(http.StatusOK, gin.H{
			"x轴":  data.Xray,
			"负荷量": data.Load,
			"温度":  data.Temp,
		})
	}

	//b := []float64{10.2, 9.4, 8.3, 8.3, 9.7, 10.1, 11.5, 12.4, 13.8, 16.2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	/*
		var a, b, x []int

		if index == "D1组团" {
			a = []int{111, 116, 115, 120, 121, 110, 110, 106, 96, 80, 65, 51, 41, 35, 25, 51, 66, 71, 0, 0, 0, 0, 0, 0, 0, 0}
			b = []int{-2, -3, -3, -4, -4, -2, -2, -1, 1, 4, 7, 10, 12, 13, 15, 10, 7, 6, 5, 4, 4, 3, 2, 1, 0, -1}
		}
	*/
}

func GetComparison(c *gin.Context) {

	index := c.Query("index")
	var a, b []float64
	var x []int
	var val string
	var err error

	switch index {
	case "D1组团":
		val, err = model.RedisClient.Get("loadComparison_1").Result()
	case "D2组团":
		val, err = model.RedisClient.Get("loadComparison_2").Result()
	case "D3组团":
		val, err = model.RedisClient.Get("loadComparison_3").Result()
	case "D4组团":
		val, err = model.RedisClient.Get("loadComparison_4").Result()
	case "D5组团":
		val, err = model.RedisClient.Get("loadComparison_5").Result()
	case "D6组团":
		val, err = model.RedisClient.Get("loadComparison_6").Result()
	case "公共组团南区":
		val, err = model.RedisClient.Get("loadComparison_7").Result()
	case "公共组团北区":
		val, err = model.RedisClient.Get("loadComparison_8").Result()
	}

	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			a = model.GetLoad(index, "today")
			predict := model.LoadPredict(index)
			b = make([]float64, 24)
			for i := 0; i < 24; i++ {
				b[i] = predict.Result[i]
			}
			x = make([]int, 24)
			for i := 0; i < 24; i++ {
				x[i] = i
			}
			for i := 0; i < len(a); i++ {
				a[i], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", a[i]), 64)
				b[i], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", b[i]), 64)
			}

			c.JSON(http.StatusOK, gin.H{
				"x轴":  x,
				"真实值": a,
				"预测值": b,
			})

			loadComparisonStruct := LoadComparison{}
			loadComparisonStruct.Real = a
			loadComparisonStruct.Predict = b
			loadComparisonStruct.Xray = x

			result, _ := json.Marshal(&loadComparisonStruct)

			switch index {
			case "D1组团":
				model.RedisClient.Set("loadComparison_1", result, time.Hour)
			case "D2组团":
				model.RedisClient.Set("loadComparison_2", result, time.Hour)
			case "D3组团":
				model.RedisClient.Set("loadComparison_3", result, time.Hour)
			case "D4组团":
				model.RedisClient.Set("loadComparison_4", result, time.Hour)
			case "D5组团":
				model.RedisClient.Set("loadComparison_5", result, time.Hour)
			case "D6组团":
				model.RedisClient.Set("loadComparison_6", result, time.Hour)
			case "公共组团南区":
				model.RedisClient.Set("loadComparison_7", result, time.Hour)
			case "公共组团北区":
				model.RedisClient.Set("loadComparison_8", result, time.Hour)
			}
		}
	} else {
		data := LoadComparison{}
		_ = json.Unmarshal([]byte(val), &data)

		c.JSON(http.StatusOK, gin.H{
			"x轴":  data.Xray,
			"真实值": data.Real,
			"预测值": data.Predict,
		})
	}

	/*
		if index == "D1组团" {
			a = []int{111, 116, 115, 120, 121, 110, 110, 106, 96, 80, 65, 51, 41, 35, 25, 51, 66, 71, 0, 0, 0, 0, 0, 0, 0, 0}
			b = []int{115, 123, 121, 123, 127, 118, 113, 109, 106, 65, 60, 58, 39, 42, 22, 51, 48, 73, 56, 99, 84, 95, 108, 106, 110, 113}
		}

	*/
}

func getComparisonData(group *singleflight.Group, key string, index string) interface{} {
	value, _, _ := group.Do(key, func() (ret interface{}, err error) { //do的入参key，可以直接使用缓存的key，这样同一个缓存，只有一个协程会去读DB

		//fmt.Println("test")
		val, err := model.RedisClient.Get("loadComparison_1").Result()
		if err != nil {
			// 如果返回的错误是key不存在
			if errors.Is(err, redis.Nil) {
				var a, b []float64
				var x []int
				a = model.GetLoad(index, "today")
				predict := model.LoadPredict(index)
				b = make([]float64, 24)
				for i := 0; i < 24; i++ {
					b[i] = predict.Result[i]
				}
				x = make([]int, 24)
				for i := 0; i < 24; i++ {
					x[i] = i
				}
				for i := 0; i < len(a); i++ {
					a[i], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", a[i]), 64)
					b[i], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", b[i]), 64)
				}

				loadComparisonStruct := LoadComparison{}
				loadComparisonStruct.Real = a
				loadComparisonStruct.Predict = b
				loadComparisonStruct.Xray = x

				result, _ := json.Marshal(&loadComparisonStruct)
				model.RedisClient.Set("loadComparison_1", result, 10*time.Minute)
				return loadComparisonStruct, nil
			}
		} else {
			data := LoadComparison{}
			_ = json.Unmarshal([]byte(val), &data)
			return data, nil
		}
		return "err", nil
	})
	return value
}

var (
	group singleflight.Group
	wg    sync.WaitGroup
)

func GetComparison2(c *gin.Context) {
	index := c.Query("index")
	wg.Add(1)
	key := "comparison"

	go func() {
		defer wg.Done()
		value := getComparisonData(&group, key, index)

		fmt.Println(value)

		c.JSON(http.StatusOK, value)
	}()
	wg.Wait()
}

func GetRealLoad(c *gin.Context) {
	index := c.Query("index")
	flag := c.Query("flag")
	a := model.GetLoad(index, flag)

	c.JSON(http.StatusOK, gin.H{
		"真实值": a,
	})
}

func GetLoadPredict(c *gin.Context) {
	index := c.Query("index")
	a := model.LoadPredict(index)

	c.JSON(http.StatusOK, gin.H{
		"预测值": a,
	})
}

func GetTemp(c *gin.Context) {
	a := model.GetData("temperature", int(time.Now().Unix()), "")

	c.JSON(http.StatusOK, gin.H{
		"code": a,
	})
}

func Test(c *gin.Context) {
	a := model.GetForecast()
	fmt.Println(a)
	c.JSON(http.StatusOK, gin.H{
		"code": a,
	})
}
