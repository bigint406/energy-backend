package energyConfig

import (
	"encoding/json"
	"energy/defs"
	_ "energy/defs"
	"energy/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"strings"
	"time"
)

type EnergyConfigDailyController struct {
}

type TankConfig struct {
	Data [6]float64
}

var (
	Vally_cost_time_start  = 23
	Vally_cost_time_end    = 7
	Flat_cost_time_1_start = 7
	Flat_cost_time_1_end   = 10
	Flat_cost_time_2_start = 15
	Flat_cost_time_2_end   = 18
	Flat_cost_time_3_start = 21
	Flat_cost_time_3_end   = 23
	Peak_cost_time_1_start = 10
	Peak_cost_time_1_end   = 15
	Peak_cost_time_2_start = 18
	Peak_cost_time_2_end   = 21
)

// var loadDaily = [24]float64{206.54, 250.18, 214.85, 167.64, 182.05, 191.49, 211.57, 89.44, 27.73, 14.62, 7.68, 32.10, 32.35, 4.84, 33.30, 50.11, 37.97, 5.39, 22.92, 23.98, 87.57, 79.91, 89.96, 203.82}
//var loadDaily = [24]float64{369.94, 355.52, 324.63, 308.96, 289.64, 191.60, 333.00, 177.77, 215.62, 159.51, 165.07, 168.37, 235.35, 218.12, 337.49, 329.06, 140.63, 213.57, 282.12, 299.11, 373.90, 514.68, 313.60, 410.21}
//var loadDaily = [24]float64{172.993275, 145.30265555555556, 145.32967361111113, 135.0105069444444, 146.40135416666666, 160.7001666666667, 178.58025694444444, 132.61163611111118, 131.68320277777775, 126.19351111111106, 121.09128888888888, 108.72937638888875, 96.90038750000004, 99.32340833333335, 95.2395499999999, 90.11576388888895, 92.47189583333326, 83.36640277777775, 99.87508611111106, 91.50452500000009, 105.88248611111115, 91.03844166666666, 87.35436111111117, 83.53618194444446}

var loadDaily = [24]float64{}

var Energy = model.EnergyConfigDaily{
	Qs:                      29768,
	Tank_top_export_temp:    80,
	Tank_bottom_export_temp: 80,
	Vally_cost_time_start:   Vally_cost_time_start,
	Vally_cost_time_end:     Vally_cost_time_end,
	Flat_cost_time_1_start:  Flat_cost_time_1_start,
	Flat_cost_time_1_end:    Flat_cost_time_1_end,
	Flat_cost_time_2_start:  Flat_cost_time_2_start,
	Flat_cost_time_2_end:    Flat_cost_time_2_end,
	Flat_cost_time_3_start:  Flat_cost_time_3_start,
	Flat_cost_time_3_end:    Flat_cost_time_3_end,
	Peak_cost_time_1_start:  Peak_cost_time_1_start,
	Peak_cost_time_1_end:    Peak_cost_time_1_end,
	Peak_cost_time_2_start:  Peak_cost_time_2_start,
	Peak_cost_time_2_end:    Peak_cost_time_2_end,

	Startup_1_boiler_lower_limiting_load_value: 400,
	Startup_2_boiler_lower_limiting_load_value: 3000,
	Startup_3_boiler_lower_limiting_load_value: 7000,
	Startup_4_boiler_lower_limiting_load_value: 12000,

	Daily_load_prediction: loadDaily,
}

func GetPeriod(c *gin.Context) {
	flat := [6]int{Flat_cost_time_1_start, Flat_cost_time_1_end, Flat_cost_time_2_start, Flat_cost_time_2_end, Flat_cost_time_3_start, Flat_cost_time_3_end}
	peak := [4]int{Peak_cost_time_1_start, Peak_cost_time_1_end, Peak_cost_time_2_start, Peak_cost_time_2_end}
	vally := [2]int{Vally_cost_time_start, Vally_cost_time_end}
	c.JSON(http.StatusOK, gin.H{
		"平电价": flat,
		"峰电价": peak,
		"谷电价": vally,
	})
}

func GetDeviceWorkTime(c *gin.Context) {

	var result [9]int
	for i := 0; i < 9; i++ {
		data, _ := model.GetResultFloatList(defs.EnergyRunningTimeDay[i], model.UnixToString(int(time.Now().Unix())))
		//data, _ := model.GetResultFloatList(defs.EnergyRunningTimeDay[i], "2023/02/20")
		for j := 0; j < len(data); j++ {
			result[i] += int(data[j])
		}
	}
	for i := 0; i < 9; i++ {
		result[i] = result[i] / 60
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
		//"data": []int{2, 2, 2, 2, 2, 2, 2, 2, 2},
	})
}

func GetLoadDetail(c *gin.Context) {
	//Energy.Daily_load_prediction = model.GetTotalLoad("2023/03/18")
	//fmt.Println(Energy.Daily_load_prediction)

	fullBoilerLoad := Energy.GetBoilerLoad()
	tankHeating := Energy.GetTankHeatingLoad()
	var tankHeating2 [24]int
	var tankStorage [24]int
	var boilerLoad [24]int

	len := 24 - len(tankHeating)

	for i := len; i < 24; i++ {
		tankHeating2[i] = int(tankHeating[i-len])
	}
	for i := 0; i < len; i++ {
		tankStorage[i] = int(fullBoilerLoad[i] - Energy.Daily_load_prediction[i])
		boilerLoad[i] = int(Energy.Daily_load_prediction[i])
	}
	for i := len; i < 24; i++ {
		boilerLoad[i] = int(Energy.Daily_load_prediction[i] - tankHeating[i-(len)])
	}

	c.JSON(http.StatusOK, gin.H{
		"电锅炉负荷":  boilerLoad,
		"水箱蓄热负荷": tankStorage,
		"水箱放热负荷": tankHeating2,
	})
}

func GetBoilerConfigDaily(c *gin.Context) {
	//Energy.Daily_load_prediction = model.GetTotalLoad("2023/03/18")

	//a, _ := model.GetResultFloatList(defs.EnergyBoilerRunningNum, model.UnixToString(int(time.Now().Unix())))
	//a, _ := model.GetResultFloatNoTime(defs.EnergyBoilerRunningNum)

	var array []int
	array = make([]int, 24)

	a, _ := model.GetResultFloatList(defs.EnergyRunningTimeDay[0], model.GetToday())
	b, _ := model.GetResultFloatList(defs.EnergyRunningTimeDay[1], model.GetToday())
	e, _ := model.GetResultFloatList(defs.EnergyRunningTimeDay[2], model.GetToday())
	d, _ := model.GetResultFloatList(defs.EnergyRunningTimeDay[3], model.GetToday())

	for i := 0; i < len(a); i++ {
		if a[i] > 0 {
			array[i]++
		}
		if b[i] > 0 {
			array[i]++
		}
		if e[i] > 0 {
			array[i]++
		}
		if d[i] > 0 {
			array[i]++
		}
	}

	//a, _ := model.GetResultFloatList(defs.EnergyBoilerRunningNum, "2023/02/20")
	c.JSON(http.StatusOK, gin.H{
		"实际": array,
		"建议": Energy.GetBoilerRunningNum(),
	})

	/*
		c.JSON(http.StatusOK, gin.H{
			"实际": []int{1, 2, 2, 1, 2, 2, 2, 2, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 1, 2, 2},
			"建议": []int{2, 2, 2, 1, 2, 2, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 1, 2, 2},
		})

	*/
}

func GetData(c *gin.Context) {
	// fmt.Println("水箱放热：", energy.GetTankHeatingLoad())
	c.JSON(http.StatusOK, gin.H{
		"水箱放热":      Energy.GetTankHeatingLoad(),
		"电锅炉承担逐时负荷": Energy.GetBoilerLoad(),
	})
}

func GetConsumption(c *gin.Context) {
	time := c.Query("time")
	a, b := model.GetResultFloatList(defs.GroupHeatConsumptionHour4, time)
	fmt.Println(a)
	fmt.Println(b)
	c.JSON(http.StatusOK, gin.H{
		"code": a,
	})
}

func GetTankConfigDaily(c *gin.Context) {
	//Energy.Daily_load_prediction = model.GetTotalLoad("2023/03/18")

	var a [6]float64

	val, err := model.RedisClient.Get("tankConfigDaily").Result()
	if err != nil {
		// 如果返回的错误是key不存在
		if errors.Is(err, redis.Nil) {
			a = Energy.GetTankRecommendedHourlyWorkingCondition()

			c.JSON(http.StatusOK, gin.H{
				"data": Energy.GetTankRecommendedHourlyWorkingCondition(),
			})

			tankConfigDailyStruct := TankConfig{}
			tankConfigDailyStruct.Data = a

			result, _ := json.Marshal(&tankConfigDailyStruct)

			model.RedisClient.Set("tankConfigDaily", result, time.Minute)
		}
	} else {
		data := TankConfig{}
		_ = json.Unmarshal([]byte(val), &data)

		c.JSON(http.StatusOK, gin.H{
			"data": data.Data,
		})
	}
}

func GetWorkTime(c *gin.Context) {
	a, _ := model.GetResultFloatList(defs.EnergyRunningTimeDay[0], "2022/10/12")
	fmt.Println(a)
	c.JSON(http.StatusOK, gin.H{
		"code": a,
	})
}

func GetDeviceWorkState(c *gin.Context) {
	/*0就是关
	1-4 锅炉
	5-10 泵
	11-18 DV 1~8
	19-20 DVT
	*/
	//var array = [...]string{"ZLZ.系统运行中1", "ZLZ.系统运行中2", "ZLZ.系统运行中3", "ZLZ.系统运行中4", "ZLZ.RUN_P1", "ZLZ.RUN_P2", "ZLZ.RUN_P3", "ZLZ.RUN_P7", "ZLZ.RUN_P8", "ZLZ.RUN_P9", "ZLZ.OPEN_V1", "ZLZ.OPEN_V2", "ZLZ.OPEN_V3", "ZLZ.OPEN_V4", "ZLZ.OPEN_V5", "ZLZ.OPEN_V6", "ZLZ.OPEN_V8", "ZLZ.OPEN_V11", "ZLZ.OUTPUT_T29", "ZLZ.OUTPUT_T30"}
	var array = [...]string{"ZLZ.%E7%B3%BB%E7%BB%9F%E8%BF%90%E8%A1%8C%E4%B8%AD1", "ZLZ.%E7%B3%BB%E7%BB%9F%E8%BF%90%E8%A1%8C%E4%B8%AD2", "ZLZ.%E7%B3%BB%E7%BB%9F%E8%BF%90%E8%A1%8C%E4%B8%AD3", "ZLZ.%E7%B3%BB%E7%BB%9F%E8%BF%90%E8%A1%8C%E4%B8%AD4", "ZLZ.RUN_P1", "ZLZ.RUN_P2", "ZLZ.RUN_P3", "ZLZ.RUN_P7", "ZLZ.RUN_P8", "ZLZ.RUN_P9", "ZLZ.OPEN_V1", "ZLZ.OPEN_V2", "ZLZ.OPEN_V3", "ZLZ.OPEN_V4", "ZLZ.OPEN_V5", "ZLZ.OPEN_V6", "ZLZ.OPEN_V8", "ZLZ.OPEN_V11", "ZLZ.OUTPUT_T29", "ZLZ.OUTPUT_T30"}
	var array2 [20]int
	var result [22]int
	var stringResult [22]string

	for i := 0; i < len(array); i++ {

		//a, _ := model.GetOpcFloatList(array[i], "2023/02/20 13") //ZS
		//a, _ := model.GetOpcFloatList(array[i], "2023/04/24 12")

		a, _ := model.GetOpcFloatList(array[i], MakeTimeStr()) //ZS

		if a[0] == 0 {
			array2[i] = 0
		} else {
			array2[i] = 1
		}
	}

	for i := 0; i < 4; i++ {
		result[i] = array2[i]
	}
	if array2[10] == 0 && array2[15] == 1 {
		result[4] = 1
		result[5] = 1
	}
	for i := 6; i < 21; i++ {
		result[i] = array2[i-2]
	}

	for i := 0; i < 12; i++ {
		if result[i] == 0 {
			stringResult[i] = "不工作"
		} else if result[i] == 1 {
			stringResult[i] = "工作"
		}
	}
	for i := 12; i < 22; i++ {
		if result[i] == 0 {
			stringResult[i] = "关闭"
		} else if result[i] == 1 {
			stringResult[i] = "开通"
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": stringResult,
	})
}

func MakeTimeStr() string {
	timeLayout := "2006-01-02 15:04:05"
	timeStr := time.Unix(time.Now().Unix(), 0).Format(timeLayout)
	a := strings.Split(timeStr, ":")
	b := a[0]
	b = strings.Replace(b, "-", "/", -1)
	return b
}
