package loadPredict

import (
	"energy/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetLoadStatistic(c *gin.Context) {
	index := c.Query("index")
	a := model.GetLoad(index, "today")
	b := model.GetData("temperature", int(time.Now().Unix()), "")

	//b := []float64{10.2, 9.4, 8.3, 8.3, 9.7, 10.1, 11.5, 12.4, 13.8, 16.2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var x []int
	/*
		var a, b, x []int

		if index == "D1组团" {
			a = []int{111, 116, 115, 120, 121, 110, 110, 106, 96, 80, 65, 51, 41, 35, 25, 51, 66, 71, 0, 0, 0, 0, 0, 0, 0, 0}
			b = []int{-2, -3, -3, -4, -4, -2, -2, -1, 1, 4, 7, 10, 12, 13, 15, 10, 7, 6, 5, 4, 4, 3, 2, 1, 0, -1}
		}
	*/
	for i := 0; i < len(a); i++ {
		a[i], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", a[i]), 64)
	}

	x = make([]int, 24)
	for i := 0; i < 24; i++ {
		x[i] = i
	}

	c.JSON(http.StatusOK, gin.H{
		"x轴":  x,
		"负荷量": a,
		"温度":  b,
	})
}

func GetComparison(c *gin.Context) {
	index := c.Query("index")
	var x []int
	a := model.GetLoad(index, "today")
	data := model.LoadPredict(index)

	var b []float64
	b = make([]float64, 24)
	for i := 0; i < 24; i++ {
		b[i] = data.Result[i]
	}

	for i := 0; i < len(a); i++ {
		a[i], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", a[i]), 64)
		b[i], _ = strconv.ParseFloat(fmt.Sprintf("%.2f", b[i]), 64)
	}

	/*
		if index == "D1组团" {
			a = []int{111, 116, 115, 120, 121, 110, 110, 106, 96, 80, 65, 51, 41, 35, 25, 51, 66, 71, 0, 0, 0, 0, 0, 0, 0, 0}
			b = []int{115, 123, 121, 123, 127, 118, 113, 109, 106, 65, 60, 58, 39, 42, 22, 51, 48, 73, 56, 99, 84, 95, 108, 106, 110, 113}
		}

	*/
	x = make([]int, 24)
	for i := 0; i < 24; i++ {
		x[i] = i
	}

	c.JSON(http.StatusOK, gin.H{
		"x轴":  x,
		"真实值": a,
		"预测值": b,
	})
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
