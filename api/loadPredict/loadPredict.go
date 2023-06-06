package loadPredict

import (
	"encoding/json"
	"energy/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"strconv"
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

	val, err := model.RedisClient.Get("loadStatistic_1").Result()
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

			model.RedisClient.Set("loadStatistic_1", result, time.Minute)
		}
	} else {
		data := LoadStatistic{}
		_ = json.Unmarshal([]byte(val), &data)

		c.JSON(http.StatusOK, data)
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

	val, err := model.RedisClient.Get("loadComparison_1").Result()
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

			model.RedisClient.Set("loadComparison_1", result, time.Hour)
		}
	} else {
		data := LoadComparison{}
		_ = json.Unmarshal([]byte(val), &data)

		c.JSON(http.StatusOK, data)
	}

	/*
		if index == "D1组团" {
			a = []int{111, 116, 115, 120, 121, 110, 110, 106, 96, 80, 65, 51, 41, 35, 25, 51, 66, 71, 0, 0, 0, 0, 0, 0, 0, 0}
			b = []int{115, 123, 121, 123, 127, 118, 113, 109, 106, 65, 60, 58, 39, 42, 22, 51, 48, 73, 56, 99, 84, 95, 108, 106, 110, 113}
		}

	*/
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
