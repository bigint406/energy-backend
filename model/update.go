package model

import (
	"context"
	"energy/defs"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func loopTime(t time.Duration, callback func(time.Time, bool)) {
	if t == 0 {
		return
	}
	for {
		now := time.Now().Add(-time.Minute * 10)
		nano := time.Duration(now.UnixNano())
		deltaT := nano/t*t + t - nano
		next := now.Add(deltaT)
		time.Sleep(deltaT) //前往下一个整点

		go callback(next, true)
	}
}

// 更新本月的数据，传入的是本月最后一天的t
func updateMonth(t time.Time) {
	month := int(t.Month())
	yearStr := fmt.Sprintf("%04d", t.Year())
	monthStr := fmt.Sprintf("%s/%02d", yearStr, month)

	data := CalcEnergyCarbonMonth(monthStr) //能源站碳排
	MongoUpdateList(yearStr, month, defs.EnergyCarbonYear, data)
	data = CalcEnergyPayloadMonth(monthStr) //能源站锅炉负载
	MongoUpdateList(yearStr, month, defs.EnergyBoilerPayloadYear, data)

	data = CalcColdCarbonMonth(monthStr) //制冷站碳排
	MongoUpdateList(yearStr, month, defs.ColdCarbonYear, data)

	//二次泵站
	data = CalcPumpCarbonMonth(monthStr)
	MongoUpdateList(yearStr, month, defs.PumpCarbonYear, data)

	data = CalcSolarWaterHeatCollectionMonth(monthStr) //太阳能热水集热量
	MongoUpdateList(yearStr, month, defs.SolarWaterHeatCollectionYear, data)
	data = CalcSolarWaterHeatEfficiencyMonth(monthStr) //集热效率
	MongoUpdateList(yearStr, month, defs.SolarWaterHeatEfficiencyYear, data)
	data = CalcSolarWaterGuaranteeRateMonth(monthStr) //保证率
	MongoUpdateList(yearStr, month, defs.SolarWaterGuaranteeRateYear, data)

	//太阳能发电
	data = SumOpcResultList(defs.SolarElecGenMonth, monthStr) //太阳能发电
	MongoUpdateList(yearStr, month, defs.SolarElecGenYear, data)
}

// 更新本日的数据
func updateDay(t time.Time, upsert bool) {
	day := t.Day()
	monthStr := fmt.Sprintf("%04d/%02d", t.Year(), t.Month())
	dayStr := fmt.Sprintf("%s/%02d", monthStr, day)

	data := CalcEnergyCarbonDay(dayStr) //能源站碳排
	MongoUpdateList(monthStr, day, defs.EnergyCarbonMonth, data)
	data = CalcEnergyPayloadDay(dayStr) //能源站锅炉负载
	MongoUpdateList(monthStr, day, defs.EnergyBoilerPayloadMonth, data)

	data = CalcColdCarbonDay(dayStr) //制冷站碳排
	MongoUpdateList(monthStr, day, defs.ColdCarbonMonth, data)

	//二次泵站
	data = CalcPumpCarbonDay(dayStr)
	MongoUpdateList(monthStr, day, defs.PumpCarbonMonth, data)

	//太阳能热水
	data = CalcSolarWaterHeatCollectionDay(dayStr) //集热量
	MongoUpdateList(monthStr, day, defs.SolarWaterHeatCollectionMonth, data)
	data = CalcSolarWaterHeatEfficiencyDay(dayStr) //集热效率
	MongoUpdateList(monthStr, day, defs.SolarWaterHeatEfficiencyMonth, data)
	data = CalcSolarWaterGuaranteeRateDay(dayStr) //保证率
	MongoUpdateList(monthStr, day, defs.SolarWaterGuaranteeRateMonth, data)
	//太阳能发电
	data = SumOpcResultList(defs.SolarElecGenDay, dayStr) //发电量
	MongoUpdateList(monthStr, day, defs.SolarElecGenMonth, data)

	if t.Add(time.Hour*24).Month() != t.Month() {
		updateMonth(t)
	}
}

// 更新本小时的数据
func updateHour(t time.Time, upsert bool) {
	hour := t.Hour()
	dayStr := fmt.Sprintf("%04d/%02d/%02d", t.Year(), t.Month(), t.Day())
	hourStr := fmt.Sprintf("%s %02d", dayStr, hour)
	var data float64
	q3 := CalcEnergyHeatStorageAndRelease(hourStr) //蓄放热量统计，正值蓄热，负值放热
	MongoUpdateList(dayStr, hour, defs.EnergyHeatStorageAndRelease, q3)
	q1 := CalcEnergyBoilerHeatSupply(hourStr)     //能源站锅炉供热量
	q2List := CalcEnergyBoilerEnergyCost(hourStr) //各锅炉能耗(单位kW·h)
	q2 := q2List[0] + q2List[1] + q2List[2] + q2List[3]
	MongoUpdateList(dayStr, hour, defs.EnergyBoilerPowerConsumptionDay1, q2List[0])
	MongoUpdateList(dayStr, hour, defs.EnergyBoilerPowerConsumptionDay2, q2List[1])
	MongoUpdateList(dayStr, hour, defs.EnergyBoilerPowerConsumptionDay3, q2List[2])
	MongoUpdateList(dayStr, hour, defs.EnergyBoilerPowerConsumptionDay4, q2List[3])
	MongoUpdateList(dayStr, hour, defs.EnergyBoilerEnergyCost, q2)
	data = CalcEnergyBoilerEfficiency(q1, q2) //能源站锅炉效率
	MongoUpdateList(dayStr, hour, defs.EnergyBoilerEfficiencyDay, data)
	data = CalcWatertankEfficiency(q3, hourStr) //能源站蓄热水箱效率
	MongoUpdateList(dayStr, hour, defs.EnergyWatertankEfficiencyDay, data)
	data = CalcEnergyEfficiency(hourStr) //能源站效率
	MongoUpdateList(dayStr, hour, defs.EnergyEfficiencyDay, data)
	data = CalcEnergyCarbonHour(hourStr, q2) //能源站碳排
	MongoUpdateList(dayStr, hour, defs.EnergyCarbonDay, data)
	data = CalcEnergyPayloadHour(q1) //能源站锅炉负载率
	MongoUpdateList(dayStr, hour, defs.EnergyBoilerPayloadDay, data)
	dataList := CalcEnergyRunningTimeHour(hourStr) //设备运行时间（分钟）
	for i := 0; i < 9; i++ {
		MongoUpdateList(dayStr, hour, defs.EnergyRunningTimeDay[i], dataList[i])
	}

	//制冷中心
	q1 = CalcColdEnergyCost(hourStr, defs.ColdMachine1)
	q2 = CalcColdEnergyCost(hourStr, defs.ColdMachine2)
	q3 = CalcColdEnergyCost(hourStr, defs.ColdMachine3)
	q := q1 + q2 + q3
	MongoUpdateList(dayStr, hour, defs.ColdEnergyCostDay, q) //耗能
	//制冷效率（流量没拿到）
	//data = CalcColdEfficiency(hourStr, q)

	//碳排
	data = CalcColdCarbonHour(q)
	MongoUpdateList(dayStr, hour, defs.ColdCarbonDay, data)
	//负载率（流量没拿到）

	//二次泵站
	data = CalcPumpEnergyCostHour(hourStr) //耗能
	MongoUpdateList(dayStr, hour, defs.PumpEnergyCostDay, data)
	dataList = CalcPumpEHR(hourStr) //输热比
	MongoUpdateList(dayStr, hour, defs.PumpEHR1, dataList[0])
	MongoUpdateList(dayStr, hour, defs.PumpEHR2, dataList[1])

	dataList = CalcHeatConsumptionHour(hourStr) //耗热统计
	MongoUpdateList(dayStr, hour, defs.GroupHeatConsumptionDay1, dataList[0])
	MongoUpdateList(dayStr, hour, defs.GroupHeatConsumptionDay2, dataList[1])
	MongoUpdateList(dayStr, hour, defs.GroupHeatConsumptionDay3, dataList[2])
	MongoUpdateList(dayStr, hour, defs.GroupHeatConsumptionDay4, dataList[3])
	MongoUpdateList(dayStr, hour, defs.GroupHeatConsumptionDay5, dataList[4])
	MongoUpdateList(dayStr, hour, defs.GroupHeatConsumptionDay6, dataList[5])
	MongoUpdateList(dayStr, hour, defs.GroupHeatConsumptionDayPubS, dataList[6])

	//太阳能热水
	q1 = CalcSolarWaterHeatCollectionHour(hourStr) //集热量
	MongoUpdateList(dayStr, hour, defs.SolarWaterHeatCollectionDay, q1)
	data = CalcSolarWaterHeatEfficiency(t, q1) //集热效率
	MongoUpdateList(dayStr, hour, defs.SolarWaterHeatEfficiencyDay, data)
	q2 = CalcSolarWaterBoilerPowerConsumptionHour(hourStr) //电加热器耗电
	MongoUpdateList(dayStr, hour, defs.SolarWaterBoilerPowerConsumptionDay, q2)
	data = CalcSolarWaterGuaranteeRate(q1, q2) //保证率
	MongoUpdateList(dayStr, hour, defs.SolarWaterGuaranteeRateDay, data)
	//太阳能发电
	data = CalcSolarElecGenHour(hourStr) //本小时发电量
	MongoUpdateList(dayStr, hour, defs.SolarElecGenDay, data)

	if hour == 23 {
		updateDay(t, upsert)
	}
}

// 更新上一分钟的数据。如果仅仅是导入过去数据upsert设为false，需要更新页面设为true
func updateMinute(t time.Time, upsert bool) {
	lastMinTime := t
	lastMin := lastMinTime.Minute()
	month := int(t.Month())
	lastMinYearStr := fmt.Sprintf("%04d", t.Year())
	lastMinLastYearStr := fmt.Sprintf("%04d", t.Year()-1)
	lastMinMonthStr := fmt.Sprintf("%s/%02d", lastMinYearStr, month)
	lastMinDayStr := fmt.Sprintf("%04d/%02d/%02d", lastMinTime.Year(), lastMinTime.Month(), lastMinTime.Day())
	lastMinHourStr := fmt.Sprintf("%s %02d", lastMinDayStr, lastMinTime.Hour())
	lastMinStr := fmt.Sprintf("%s:%02d", lastMinHourStr, lastMinTime.Minute())
	log.Println(lastMinStr)
	var data float64
	var page Pages
	if upsert {
		page.PageInit()
	}
	//报警数据
	data = UpdateEnergyAlarm(lastMinHourStr, lastMin, lastMinTime) //能源站
	if upsert {
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyAlarmNumToday, data)
		page.PageUpdate(defs.PageBasicMap, defs.EnergyAlarmNumToday, data)
	}
	data = UpdateColdAlarm(lastMinHourStr, lastMin, lastMinTime) //制冷中心
	if upsert {
		page.PageUpdate(defs.PageSystemRefigeration, defs.ColdAlarmNumToday, data)
		page.PageUpdate(defs.PageBasicMap, defs.PageSystemRefigeration, data)
	}
	data = UpdatePumpAlarm(lastMinHourStr, lastMin, lastMinTime) //二次泵站
	if upsert {
		page.PageUpdate(defs.PageSystemPump, defs.PumpAlarmNumToday, data)
		page.PageUpdate(defs.PageBasicMap, defs.PageSystemPump, data)
	}
	//之后计算要用的数据

	// exampleTime := "2022/05/01 08:05"
	//二次泵站
	var HeatData defs.LouHeatList
	err := MongoLoukong.FindOne(context.TODO(), bson.D{{"time", lastMinStr}, {"name", "heat"}}).Decode(&HeatData)
	// err := MongoLoukong.FindOne(context.TODO(), bson.D{{"time", exampleTime}, {"name", "heat"}}).Decode(&HeatData)
	if err == nil {
		dataList := CalcPumpHeat(&HeatData) //统计输热量
		MongoUpdateList(lastMinHourStr, lastMin, defs.GroupHeatConsumptionHour1, dataList[0])
		MongoUpdateList(lastMinHourStr, lastMin, defs.GroupHeatConsumptionHour2, dataList[1])
		MongoUpdateList(lastMinHourStr, lastMin, defs.GroupHeatConsumptionHour3, dataList[2])
		MongoUpdateList(lastMinHourStr, lastMin, defs.GroupHeatConsumptionHour4, dataList[3])
		MongoUpdateList(lastMinHourStr, lastMin, defs.GroupHeatConsumptionHour5, dataList[4])
		MongoUpdateList(lastMinHourStr, lastMin, defs.GroupHeatConsumptionHour6, dataList[5])
		MongoUpdateList(lastMinHourStr, lastMin, defs.GroupHeatConsumptionHourPubS, dataList[6])
		MongoUpdateList(lastMinHourStr, lastMin, defs.PumpHeatHour1, dataList[0]+dataList[1])
		MongoUpdateList(lastMinHourStr, lastMin, defs.PumpHeatHour2, dataList[2]+dataList[3]+dataList[4]+dataList[5])
	} else {
		log.Println(lastMinStr)
		log.Println("heat data miss")
		log.Println(err)
	}
	//太阳能热水
	var GAData defs.LouSolarWaterList
	GAerr := MongoLoukong.FindOne(context.TODO(), bson.D{{"time", lastMinStr}, {"name", "GA"}}).Decode(&GAData)
	// GAerr := MongoLoukong.FindOne(context.TODO(), bson.D{{"time", exampleTime}, {"name", "GA"}}).Decode(&GAData)
	if GAerr == nil {
		data = CalcSolarWaterHeatCollectionMin(&GAData.Info) //集热量
		MongoUpdateList(lastMinHourStr, lastMin, defs.SolarWaterHeatCollectionHour, data)
		data = CalcSolarWaterBoilerPowerConsumptionMin(&GAData.Info) //电加热器耗电
		MongoUpdateList(lastMinHourStr, lastMin, defs.SolarWaterBoilerPowerConsumptionHour, data)
	} else {
		log.Println(lastMinStr)
		log.Println("GA data miss")
		log.Println(err)
	}

	var HData defs.LouH
	Herr := MongoLoukong.FindOne(context.TODO(), bson.D{{"time", lastMinStr}, {"name", "H"}}).Decode(&HData)
	if Herr == nil {

	} else {
		log.Println(lastMinStr)
		log.Println("H data miss")
		log.Println(err)
	}

	//实时展示数据
	if upsert {
		var dataList []float64
		//基础设施地图
		if Herr == nil {
			dataList = CalcBasicMapHallwayTemp(&HData.Info)
			for i := 0; i < len(dataList); i++ {
				page.PageUpdate(defs.PageBasicMap, defs.GroupHallwayTemp[i], dataList[i])
			}
		}
		//能源站
		data, _ = CalcEnergyOnlineRate(lastMinHourStr) //能源站设备在线率
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyOnlineRate, data)
		page.PageUpdate(defs.PageBasicMap, defs.EnergyOnlineRate, data)

		data = CalcEnergyBoilerPower(lastMinHourStr, lastMin) //能源站锅炉总功率
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyBoilerPower, data)

		dataList = CalcEnergyBoilerEnergyCost(lastMinHourStr) //本小时各锅炉能耗(单位kW·h)
		q23 := dataList[0] + dataList[1] + dataList[2] + dataList[3]
		dataList = CalcEnergyBoilerEnergyCostToday(lastMinDayStr, dataList) //今日各锅炉能耗
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyBoilerPowerConsumptionToday1, dataList[0])
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyBoilerPowerConsumptionToday2, dataList[1])
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyBoilerPowerConsumptionToday3, dataList[2])
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyBoilerPowerConsumptionToday4, dataList[3])

		data = CalcEnergyPowerConsumptionToday(lastMinTime, q23) //能源站今日能耗
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyPowerConsumptionToday, data)

		data = CalcEnergyBoilerRunningNum(lastMinHourStr, lastMin) //能源站锅炉运行数目
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyBoilerRunningNum, data)

		data = CalcEnergyTankRunningNum(lastMinHourStr, lastMin) //蓄热水箱运行台数
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyTankRunningNum, data)

		dataList = CalcEnergyRunningTimeToday(lastMinTime) //设备今日运行时长
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyRunningTimeToday, dataList)

		data = CalcEnergyHeatSupplyToday(t) //总供热量
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyHeatSupplyToday, data)

		//制冷中心
		q1 := CalcColdMachinePower(lastMinHourStr, lastMin, defs.ColdMachine1)
		q2 := CalcColdMachinePower(lastMinHourStr, lastMin, defs.ColdMachine2)
		q3 := CalcColdMachinePower(lastMinHourStr, lastMin, defs.ColdMachine3)
		q4 := CalcColdCabinetPower(lastMinHourStr, lastMin)
		data = q1 + q2 + q3 + q4 //总功率
		page.PageUpdate(defs.PageSystemRefigeration, defs.ColdPowerMin, data)

		data = q1 + q2 + q3 //制冷机功率
		page.PageUpdate(defs.PageSystemRefigeration, defs.ColdMachinePowerMin, data)

		data = CalcColdEnergyCostToday(lastMinTime) //今日耗能
		page.PageUpdate(defs.PageSystemRefigeration, defs.ColdEnergyCostToday, data)

		data = CalcColdMachineRunningNum(lastMinHourStr, lastMin) //制冷机运行数目
		page.PageUpdate(defs.PageSystemRefigeration, defs.ColdMachineRunningNum, data)

		data = CalcColdCoolingWaterInT(lastMinHourStr, lastMin) //冷却进水温度
		page.PageUpdate(defs.PageSystemRefigeration, defs.ColdCoolingWaterInT, data)

		data = CalcColdCoolingWaterOutT(lastMinHourStr, lastMin) //冷却出水温度
		page.PageUpdate(defs.PageSystemRefigeration, defs.ColdCoolingWaterOutT, data)

		data = CalcColdRefrigeratedWaterInT(lastMinHourStr, lastMin) //冷冻进水温度
		page.PageUpdate(defs.PageSystemRefigeration, defs.ColdRefrigeratedWaterInT, data)

		data = CalcColdRefrigeratedWaterOutT(lastMinHourStr, lastMin) //冷冻出水温度
		page.PageUpdate(defs.PageSystemRefigeration, defs.ColdRefrigeratedWaterOutT, data)

		//二次泵站
		data = CalcPumpPowerMin(lastMinHourStr, lastMin) //总功率
		page.PageUpdate(defs.PageSystemPump, defs.PumpPowerMin, data)

		data = CalcPumpEnergyCostToday(lastMinTime) //今日耗电量
		page.PageUpdate(defs.PageSystemPump, defs.PumpPowerToday, data)

		data = CalcPumpRunningState(lastMinHourStr, lastMin, 1) //泵运行状态
		page.PageUpdate(defs.PageSystemPump, defs.PumpRunningState1, data)

		data = CalcPumpRunningState(lastMinHourStr, lastMin, 2) //泵运行状态
		page.PageUpdate(defs.PageSystemPump, defs.PumpRunningState2, data)

		data = CalcPumpRunningState(lastMinHourStr, lastMin, 3) //泵运行状态
		page.PageUpdate(defs.PageSystemPump, defs.PumpRunningState3, data)

		data = CalcPumpRunningState(lastMinHourStr, lastMin, 4) //泵运行状态
		page.PageUpdate(defs.PageSystemPump, defs.PumpRunningState4, data)

		data = CalcPumpRunningState(lastMinHourStr, lastMin, 5) //泵运行状态
		page.PageUpdate(defs.PageSystemPump, defs.PumpRunningState5, data)

		data = CalcPumpRunningState(lastMinHourStr, lastMin, 6) //泵运行状态
		page.PageUpdate(defs.PageSystemPump, defs.PumpRunningState6, data)

		//太阳能热水
		if GAerr == nil {
			data = CalcSolarWaterBoilerPowerConsumptionToday(t) //电加热器今日总耗电量
			page.PageUpdate(defs.PageSystemSolarWater, defs.SolarWaterBoilerPowerConsumptionToday, data)

			data = CalcSolarWaterHeatCollecterInT(&GAData.Info) //集热器进口温度
			page.PageUpdate(defs.PageSystemSolarWater, defs.SolarWaterHeatCollecterInT, data)

			data = CalcSolarWaterHeatCollecterOutT(&GAData.Info) //集热器出口温度
			page.PageUpdate(defs.PageSystemSolarWater, defs.SolarWaterHeatCollecterOutT, data)

			data = CalcSolarWaterJRQT(&GAData.Info) //锅炉温度
			page.PageUpdate(defs.PageSystemSolarWater, defs.SolarWaterJRQT, data)

			data = CalcSolarWaterHeatCollectionToday(t) //今日总集热量
			page.PageUpdate(defs.PageSystemSolarWater, defs.SolarWaterHeatCollectionToday, data)

			data = CalcSolarWaterPumpRunningNum(&GAData.Info) //水泵运行数目
			page.PageUpdate(defs.PageSystemSolarWater, defs.SolarWaterPumpRunningNum, data)
		}
	}

	if lastMinTime.Minute() == 59 {
		updateHour(lastMinTime, upsert)
	}

	//更新页面数据
	if upsert {
		//基础设施地图
		datalist, _ := GetResultFloatList(defs.GroupHeatConsumptionHour1, lastMinHourStr)
		page.PageUpdate(defs.PageBasicMap, defs.GroupHeatConsumptionHour1, datalist)
		datalist, _ = GetResultFloatList(defs.GroupHeatConsumptionHour2, lastMinHourStr)
		page.PageUpdate(defs.PageBasicMap, defs.GroupHeatConsumptionHour2, datalist)
		datalist, _ = GetResultFloatList(defs.GroupHeatConsumptionHour3, lastMinHourStr)
		page.PageUpdate(defs.PageBasicMap, defs.GroupHeatConsumptionHour3, datalist)
		datalist, _ = GetResultFloatList(defs.GroupHeatConsumptionHour4, lastMinHourStr)
		page.PageUpdate(defs.PageBasicMap, defs.GroupHeatConsumptionHour4, datalist)
		datalist, _ = GetResultFloatList(defs.GroupHeatConsumptionHour5, lastMinHourStr)
		page.PageUpdate(defs.PageBasicMap, defs.GroupHeatConsumptionHour5, datalist)
		datalist, _ = GetResultFloatList(defs.GroupHeatConsumptionHour6, lastMinHourStr)
		page.PageUpdate(defs.PageBasicMap, defs.GroupHeatConsumptionHour6, datalist)
		datalist, _ = GetResultFloatList(defs.GroupHeatConsumptionHourPubS, lastMinHourStr)
		page.PageUpdate(defs.PageBasicMap, defs.GroupHeatConsumptionHourPubS, datalist)
		//——————————————————系统层——————————————————
		//能源站
		datalist, _ = GetResultFloatList(defs.EnergyHeatStorageAndRelease, lastMinDayStr)
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyHeatStorageAndRelease, datalist)
		datalist, _ = GetResultFloatList(defs.EnergyBoilerEnergyCost, lastMinDayStr)
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyBoilerEnergyCost, datalist)
		maplist, _ := GetResultInterfaceList(defs.EnergyAlarmToday, lastMinDayStr)
		page.PageUpdate(defs.PageSystemEnergy, defs.EnergyAlarmToday, maplist)
		for _, v := range defs.OpcItemEnergy {
			datalist, _ = GetOpcFloatList(v, lastMinHourStr)
			page.PageUpdate(defs.PageSystemEnergy, v, datalist)
		}
		//制冷中心
		datalist, _ = GetResultFloatList(defs.ColdEnergyCostDay, lastMinDayStr) //耗能
		page.PageUpdate(defs.PageSystemRefigeration, defs.ColdEnergyCostDay, datalist)
		maplist, _ = GetResultInterfaceList(defs.ColdAlarmToday, lastMinDayStr)
		page.PageUpdate(defs.PageSystemRefigeration, defs.ColdAlarmToday, maplist)
		for _, v := range defs.OpcItemRef {
			datalist, _ = GetOpcFloatList(v, lastMinHourStr)
			page.PageUpdate(defs.PageSystemRefigeration, v, datalist)
		}
		//二次泵站
		maplist, _ = GetResultInterfaceList(defs.PumpAlarmToday, lastMinDayStr)
		page.PageUpdate(defs.PageSystemPump, defs.PumpAlarmToday, maplist)
		//太阳能热水
		datalist, _ = GetResultFloatList(defs.SolarWaterHeatCollectionDay, lastMinDayStr)
		page.PageUpdate(defs.PageSystemSolarWater, defs.SolarWaterHeatCollectionDay, datalist)
		datalist, _ = GetResultFloatList(defs.SolarWaterHeatCollectionMonth, lastMinMonthStr)
		page.PageUpdate(defs.PageSystemSolarWater, defs.SolarWaterHeatCollectionMonth, datalist)
		datalist, _ = GetResultFloatList(defs.SolarWaterHeatCollectionYear, lastMinYearStr)
		page.PageUpdate(defs.PageSystemSolarWater, defs.SolarWaterHeatCollectionYear, datalist)
		datalist, _ = GetResultFloatList(defs.SolarWaterBoilerPowerConsumptionDay, lastMinDayStr)
		page.PageUpdate(defs.PageSystemSolarWater, defs.SolarWaterBoilerPowerConsumptionDay, datalist)
		//太阳能发电
		datalist, _ = GetResultFloatList(defs.SolarElecGenMonth, lastMinMonthStr)
		page.PageUpdate(defs.PageSystemSolarElec, defs.SolarElecGenMonth, datalist)
		datalist, _ = GetResultFloatList(defs.SolarElecGenYear, lastMinYearStr)
		page.PageUpdate(defs.PageSystemSolarElec, defs.SolarElecGenYear, datalist)
		for _, v := range defs.OpcItemSolarElec {
			datalist, _ = GetOpcFloatList(v, lastMinHourStr)
			page.PageUpdate(defs.PageSystemSolarElec, v, datalist)
		}
		for _, v := range defs.OpcItemGreenAndSolarElec {
			datalist, _ = GetOpcFloatList(v, lastMinHourStr)
			page.PageUpdate(defs.PageSystemSolarElec, v, datalist)
			page.PageUpdate(defs.PageGreenPower, v, datalist)
		}
		//——————————————————能效分析——————————————————
		//能源站
		datalist, _ = GetResultFloatList(defs.EnergyBoilerEfficiencyDay, lastMinDayStr)
		page.PageUpdate(defs.PageAnalyseEnergy, defs.EnergyBoilerEfficiencyDay, datalist)
		datalist, _ = GetResultFloatList(defs.EnergyEfficiencyDay, lastMinDayStr)
		page.PageUpdate(defs.PageAnalyseEnergy, defs.EnergyEfficiencyDay, datalist)
		datalist, _ = GetResultFloatList(defs.EnergyCarbonDay, lastMinDayStr)
		page.PageUpdate(defs.PageAnalyseEnergy, defs.EnergyCarbonDay, datalist)
		datalist, _ = GetResultFloatList(defs.EnergyCarbonMonth, lastMinMonthStr)
		page.PageUpdate(defs.PageAnalyseEnergy, defs.EnergyCarbonMonth, datalist)
		datalist, _ = GetResultFloatList(defs.EnergyCarbonYear, lastMinYearStr)
		page.PageUpdate(defs.PageAnalyseEnergy, defs.EnergyCarbonYear, datalist)
		datalist, _ = GetResultFloatList(defs.EnergyCarbonYear, lastMinLastYearStr)
		page.PageUpdate(defs.PageAnalyseEnergy, defs.EnergyCarbonLastYear, datalist)
		datalist, _ = GetResultFloatList(defs.EnergyBoilerPayloadDay, lastMinDayStr)
		page.PageUpdate(defs.PageAnalyseEnergy, defs.EnergyBoilerPayloadDay, datalist)
		datalist, _ = GetResultFloatList(defs.EnergyBoilerPayloadMonth, lastMinMonthStr)
		page.PageUpdate(defs.PageAnalyseEnergy, defs.EnergyBoilerPayloadMonth, datalist)
		datalist, _ = GetResultFloatList(defs.EnergyBoilerPayloadYear, lastMinYearStr)
		page.PageUpdate(defs.PageAnalyseEnergy, defs.EnergyBoilerPayloadYear, datalist)
		//制冷中心
		datalist, _ = GetResultFloatList(defs.ColdEfficientDay, lastMinDayStr)
		page.PageUpdate(defs.PageAnalyseRefigeration, defs.ColdEfficientDay, datalist)
		datalist, _ = GetResultFloatList(defs.ColdPayLoadDay, lastMinDayStr)
		page.PageUpdate(defs.PageAnalyseRefigeration, defs.ColdPayLoadDay, datalist)
		datalist, _ = GetResultFloatList(defs.ColdPayLoadMonth, lastMinMonthStr)
		page.PageUpdate(defs.PageAnalyseRefigeration, defs.ColdPayLoadMonth, datalist)
		datalist, _ = GetResultFloatList(defs.ColdPayLoadYear, lastMinYearStr)
		page.PageUpdate(defs.PageAnalyseRefigeration, defs.ColdPayLoadYear, datalist)
		datalist, _ = GetResultFloatList(defs.ColdCarbonDay, lastMinDayStr)
		page.PageUpdate(defs.PageAnalyseRefigeration, defs.ColdCarbonDay, datalist)
		datalist, _ = GetResultFloatList(defs.ColdCarbonMonth, lastMinMonthStr)
		page.PageUpdate(defs.PageAnalyseRefigeration, defs.ColdCarbonMonth, datalist)
		datalist, _ = GetResultFloatList(defs.ColdCarbonYear, lastMinYearStr)
		page.PageUpdate(defs.PageAnalyseRefigeration, defs.ColdCarbonYear, datalist)
		datalist, _ = GetResultFloatList(defs.ColdCarbonYear, lastMinLastYearStr)
		page.PageUpdate(defs.PageAnalyseRefigeration, defs.ColdCarbonLastYear, datalist)
		//二次泵站
		datalist, _ = GetResultFloatList(defs.PumpEHR1, lastMinDayStr)
		page.PageUpdate(defs.PageAnalysePump, defs.PumpEHR1, datalist)
		datalist, _ = GetResultFloatList(defs.PumpEHR2, lastMinDayStr)
		page.PageUpdate(defs.PageAnalysePump, defs.PumpEHR2, datalist)
		datalist, _ = GetResultFloatList(defs.PumpCarbonYear, lastMinYearStr)
		page.PageUpdate(defs.PageAnalysePump, defs.PumpCarbonYear, datalist)
		datalist, _ = GetResultFloatList(defs.PumpCarbonYear, lastMinLastYearStr)
		page.PageUpdate(defs.PageAnalysePump, defs.PumpCarbonLastYear, datalist)
		//太阳能热水
		datalist, _ = GetResultFloatList(defs.SolarWaterHeatEfficiencyDay, lastMinDayStr)
		page.PageUpdate(defs.PageAnalyseSolarWater, defs.SolarWaterHeatEfficiencyDay, datalist)
		datalist, _ = GetResultFloatList(defs.SolarWaterHeatEfficiencyMonth, lastMinMonthStr)
		page.PageUpdate(defs.PageAnalyseSolarWater, defs.SolarWaterHeatEfficiencyMonth, datalist)
		datalist, _ = GetResultFloatList(defs.SolarWaterHeatEfficiencyYear, lastMinYearStr)
		page.PageUpdate(defs.PageAnalyseSolarWater, defs.SolarWaterHeatEfficiencyYear, datalist)
		datalist, _ = GetResultFloatList(defs.SolarWaterGuaranteeRateDay, lastMinDayStr)
		page.PageUpdate(defs.PageAnalyseSolarWater, defs.SolarWaterGuaranteeRateDay, datalist)
		datalist, _ = GetResultFloatList(defs.SolarWaterGuaranteeRateMonth, lastMinMonthStr)
		page.PageUpdate(defs.PageAnalyseSolarWater, defs.SolarWaterGuaranteeRateMonth, datalist)
		datalist, _ = GetResultFloatList(defs.SolarWaterGuaranteeRateYear, lastMinYearStr)
		page.PageUpdate(defs.PageAnalyseSolarWater, defs.SolarWaterGuaranteeRateYear, datalist)
		//——————————————————绿电——————————————————
		data = CalcEnergyCarbonMonth(lastMinMonthStr)
		page.PageUpdate(defs.PageGreenPower, defs.TotalPowerThisMonth, data)
		for _, v := range defs.OpcItemGreenPower {
			datalist, _ = GetOpcFloatList(v, lastMinHourStr)
			page.PageUpdate(defs.PageGreenPower, v, datalist)
		}
		page.PageUpload()
	}
}

// 定时更新
func LoopQueryUpdate() {
	go loopTime(time.Minute, updateMinute)
}

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s\n", msg, err)
	}
}

// 从t1到t2计算数据，每分钟更新一次
func UpdateData(t1 time.Time, t2 time.Time) {
	for t := t1; t.Before(t2); t = t.Add(time.Minute) {
		updateMinute(t, false)
	}
	log.Print("Update Complete")
}
