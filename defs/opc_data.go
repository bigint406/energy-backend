package defs

var EnergyRunningTimeDay = []string{"energy_boiler_running_time_day1", //能源站锅炉1当日各小时运行时间（分钟）
							"energy_boiler_running_time_day2", //能源站锅炉2当日各小时运行时间（分钟）
							"energy_boiler_running_time_day3", //能源站锅炉3当日各小时运行时间（分钟）
							"energy_boiler_running_time_day4", //能源站锅炉4当日各小时运行时间（分钟）
							"energy_GRB_running_time_day1",    //能源站供热泵1当日各小时运行时间（分钟）
							"energy_GRB_running_time_day2",    //能源站供热泵2当日各小时运行时间（分钟）
							"energy_GRB_running_time_day3",    //能源站供热泵3当日各小时运行时间（分钟）
							"energy_tank_running_time_day1",   //能源站水箱1当日各小时运行时间（分钟）
							"energy_tank_running_time_day2"}   //能源站水箱2当日各小时运行时间（分钟）
var GroupHallwayTemp = []string{"group_hallway_temp1", //1组团走廊温度
	"group_hallway_temp2",      //2组团走廊温度
	"group_hallway_temp3",      //3组团走廊温度
	"group_hallway_temp4",      //4组团走廊温度
	"group_hallway_temp5",      //5组团走廊温度
	"group_hallway_temp6",      //6组团走廊温度
	"group_hallway_temp_pub_s", //公共组团南区走廊温度
	"group_hallway_temp_pub_n"} //公共组团北区走廊温度

// 页面名
const (
	PageSystemEnergy       = "sys_ener"
	PageSystemRefigeration = "sys_ref"
	PageSystemPump         = "sys_pump"
	PageSystemSolarWater   = "sys_sw"
	PageSystemSolarElec    = "sys_se"
)

// 计算结果表名
const (
	GroupHeatConsumptionHour1    = "group_heat_consumption_hour_1"     //1组团耗热量当小时每分钟
	GroupHeatConsumptionHour2    = "group_heat_consumption_hour_2"     //2组团耗热量当小时每分钟
	GroupHeatConsumptionHour3    = "group_heat_consumption_hour_3"     //3组团耗热量当小时每分钟
	GroupHeatConsumptionHour4    = "group_heat_consumption_hour_4"     //4组团耗热量当小时每分钟
	GroupHeatConsumptionHour5    = "group_heat_consumption_hour_5"     //5组团耗热量当小时每分钟
	GroupHeatConsumptionHour6    = "group_heat_consumption_hour_6"     //6组团耗热量当小时每分钟
	GroupHeatConsumptionHourPubS = "group_heat_consumption_hour_pub_s" //公共组团南区耗热量当小时每分钟

	GroupHeatConsumptionDay1    = "group_heat_consumption_day_1"     //1组团耗热量每日各小时
	GroupHeatConsumptionDay2    = "group_heat_consumption_day_2"     //2组团耗热量每日各小时
	GroupHeatConsumptionDay3    = "group_heat_consumption_day_3"     //3组团耗热量每日各小时
	GroupHeatConsumptionDay4    = "group_heat_consumption_day_4"     //4组团耗热量每日各小时
	GroupHeatConsumptionDay5    = "group_heat_consumption_day_5"     //5组团耗热量每日各小时
	GroupHeatConsumptionDay6    = "group_heat_consumption_day_6"     //6组团耗热量每日各小时
	GroupHeatConsumptionDayPubS = "group_heat_consumption_day_pub_s" //公共组团南区耗热量每日各小时

	EnergyOnlineRate                   = "energy_online_rate"                     //能源站设备在线率
	EnergyBoilerPower                  = "energy_boiler_power"                    //能源站锅炉功率
	EnergyPowerConsumptionToday        = "energy_power_consumption_today"         //能源站今日总耗能
	EnergyBoilerRunningNum             = "energy_boiler_running_num"              //能源站锅炉运行数目
	EnergyTankRunningNum               = "energy_tank_running_num"                //能源站蓄热水箱运行数目
	EnergyHeatSupplyToday              = "energy_heat_supply_today"               //能源站今日总供热量
	EnergyHeatStorageAndRelease        = "energy_heat_storage_and_release"        //能源站每日各小时水箱蓄放热量
	EnergyBoilerEnergyCost             = "energy_boiler_energy_cost"              //能源站每日各小时锅炉能耗
	EnergyBoilerEfficiencyDay          = "energy_boiler_efficiency_day"           //能源站每日各小时锅炉效率
	EnergyWatertankEfficiencyDay       = "energy_watertank_efficiency_day"        //能源站每日各小时蓄热水箱效率
	EnergyEfficiencyDay                = "energy_efficiency_day"                  //能源站每日各小时效率
	EnergyCarbonDay                    = "energy_carbon_day"                      //能源站每日各小时碳排
	EnergyCarbonMonth                  = "energy_carbon_month"                    //能源站每月各天碳排总和
	EnergyCarbonYear                   = "energy_carbon_year"                     //能源站每年各月碳排总和
	EnergyBoilerPayloadDay             = "energy_boiler_payload_day"              //能源站每日各小时锅炉负载
	EnergyBoilerPayloadMonth           = "energy_boiler_payload_month"            //能源站每月各天平均锅炉负载
	EnergyBoilerPayloadYear            = "energy_boiler_payload_year"             //能源站每年各月平均锅炉负载
	EnergyAlarmToday                   = "energy_alarm_today"                     //能源站当日告警
	EnergyAlarmNumToday                = "energy_alarm_num_today"                 //能源站今日告警次数
	EnergyBoilerPowerConsumptionDay1   = "energy_boiler_power_consumption_day1"   //能源站锅炉1每日各小时耗能
	EnergyBoilerPowerConsumptionDay2   = "energy_boiler_power_consumption_day2"   //能源站锅炉2每日各小时耗能
	EnergyBoilerPowerConsumptionDay3   = "energy_boiler_power_consumption_day3"   //能源站锅炉3每日各小时耗能
	EnergyBoilerPowerConsumptionDay4   = "energy_boiler_power_consumption_day4"   //能源站锅炉4每日各小时耗能
	EnergyBoilerPowerConsumptionToday1 = "energy_boiler_power_consumption_today1" //能源站锅炉1今日耗能
	EnergyBoilerPowerConsumptionToday2 = "energy_boiler_power_consumption_today2" //能源站锅炉2今日耗能
	EnergyBoilerPowerConsumptionToday3 = "energy_boiler_power_consumption_today3" //能源站锅炉3今日耗能
	EnergyBoilerPowerConsumptionToday4 = "energy_boiler_power_consumption_today4" //能源站锅炉4今日耗能
	EnergyRunningTimeToday             = "energy_running_time_today"              //能源站设备今日运行时长

	ColdPowerMin              = "cold_power_min"               //制冷中心当分钟功率
	ColdEnergyCostToday       = "cold_energy_cost_today"       //制冷中心今日能耗
	ColdMachineRunningNum     = "cold_machine_running_num"     //制冷中心制冷机运行台数
	ColdCoolingWaterInT       = "cold_cooling_water_InT"       //制冷中心冷却水进水温度
	ColdCoolingWaterOutT      = "cold_cooling_water_OutT"      //制冷中心冷却水出水温度
	ColdRefrigeratedWaterInT  = "cold_refrigerated_water_InT"  //制冷中心冷冻水进水温度
	ColdRefrigeratedWaterOutT = "cold_refrigerated_water_OutT" //制冷中心冷冻水出水温度
	ColdMachinePowerMin       = "cold_machine_power_min"       //制冷中心制冷机当分钟功率
	ColdMachinePowerMin1      = "cold_machine_power_min1"      //制冷中心制冷机1实时功率
	ColdMachinePowerMin2      = "cold_machine_power_min2"      //制冷中心制冷机2实时功率
	ColdMachinePowerMin3      = "cold_machine_power_min3"      //制冷中心制冷机3实时功率
	ColdEnergyCostDay         = "cold_energy_cost_day"         //制冷中心每日各小时能耗
	ColdCarbonDay             = "cold_carbon_day"              //制冷中心每日各小时碳排
	ColdCarbonMonth           = "cold_carbon_month"            //制冷中心每月各天碳排总和
	ColdCarbonYear            = "cold_carbon_year"             //制冷中心每年各月碳排总和
	ColdAlarmToday            = "cold_alarm_today"             //制冷中心今日告警
	ColdAlarmNumToday         = "cold_alarm_num_today"         //制冷中心今日告警次数

	PumpPowerMin      = "pump_power_min"       //二次泵站功率
	PumpPowerToday    = "pump_power_today"     //二次泵站今日能耗
	PumpEnergyCostDay = "pump_energy_cost_day" //二次泵站每日各小时能耗
	PumpCarbonMonth   = "pump_carbon_month"    //二次泵站每月各天碳排总和
	PumpCarbonYear    = "pump_carbon_year"     //二次泵站每年各月碳排总和
	PumpRunningState1 = "pump_running_state1"  //二次泵站泵运行状态
	PumpRunningState2 = "pump_running_state2"  //二次泵站泵运行状态
	PumpRunningState3 = "pump_running_state3"  //二次泵站泵运行状态
	PumpRunningState4 = "pump_running_state4"  //二次泵站泵运行状态
	PumpRunningState5 = "pump_running_state5"  //二次泵站泵运行状态
	PumpRunningState6 = "pump_running_state6"  //二次泵站泵运行状态
	PumpHeatHour1     = "pump_heat_hour1"      //二次泵站当小时每分钟环路1输热量
	PumpHeatHour2     = "pump_heat_hour2"      //二次泵站当小时每分钟环路2输热量
	PumpEHR1          = "pump_EHR1"            //二次泵站环路1每日EHR
	PumpEHR2          = "pump_EHR2"            //二次泵站环路2每日EHR
	PumpAlarmToday    = "pump_alarm_today"     //二次泵站今日告警
	PumpAlarmNumToday = "pump_alarm_num_today" //二次泵站今日告警次数

	SolarWaterBoilerPowerConsumptionToday = "solar_water_boiler_power_comsumption_today" //太阳能热水电加热器今日总耗电量
	SolarWaterHeatCollecterInT            = "solar_water_heat_collecter_in_temp"         //太阳能热水集热器进口温度
	SolarWaterHeatCollecterOutT           = "solar_water_heat_collecter_out_temp"        //太阳能热水集热器出口温度
	SolarWaterJRQT                        = "solar_water_JRQ_temp"                       //太阳能热水加热器温度
	SolarWaterHeatCollectionToday         = "solar_water_heat_collection_today"          //太阳能热水今日总集热量
	SolarWaterPumpRunningNum              = "solar_water_pump_running_num"               //太阳能热水水泵运行数目
	SolarWaterHeatCollectionHour          = "solar_water_heat_collection_hour"           //太阳能热水集热量当小时每分钟
	SolarWaterHeatCollectionDay           = "solar_water_heat_collection_day"            //太阳能热水集热量当日每小时
	SolarWaterHeatCollectionMonth         = "solar_water_heat_collection_month"          //太阳能热水集热量每月各天总和
	SolarWaterHeatCollectionYear          = "solar_water_heat_collection_year"           //太阳能热水集热量每年各月总和
	SolarWaterBoilerPowerConsumptionHour  = "solar_water_boiler_power_comsumption_hour"  //太阳能热水电加热器耗电量当小时每分钟
	SolarWaterBoilerPowerConsumptionDay   = "solar_water_boiler_power_comsumption_day"   //太阳能热水电加热器耗电量当日每小时
	SolarWaterHeatEfficiencyDay           = "solar_water_heat_efficiency_day"            //太阳能热水集热效率当日每小时
	SolarWaterHeatEfficiencyMonth         = "solar_water_heat_efficiency_month"          //太阳能热水集热效率每月各天
	SolarWaterHeatEfficiencyYear          = "solar_water_heat_efficiency_year"           //太阳能热水集热效率每年各月
	SolarWaterGuaranteeRateDay            = "solar_water_guarantee_rate"                 //太阳能热水保证率当日每小时
	SolarWaterGuaranteeRateMonth          = "solar_water_guarantee_month"                //太阳能热水保证率每月各天
	SolarWaterGuaranteeRateYear           = "solar_water_guarantee_year"                 //太阳能热水保证率每年各月

	SolarElecGenDay   = "solar_elec_gen_day"   //太阳能发电当日每小时
	SolarElecGenMonth = "solar_elec_gen_month" //太阳能发电当月每日
	SolarElecGenYear  = "solar_elec_gen_year"  //太阳能发电当年每月
)

// 其它常数
const (
	ColdMachine1 = "Z_LX1"
	ColdMachine2 = "Z_LX2"
	ColdMachine3 = "Z_L"
)

type OpcData struct {
	Time   string        `bson:"time"`
	ItemId string        `bson:"itemid"`
	Value  []interface{} `bson:"value"`
}

type CalculationResultFloatList struct {
	Time  string    `bson:"time"`
	Name  string    `bson:"name"`
	Value []float64 `bson:"value"`
}

type CalculationResultInterfaceList struct {
	Time  string                   `bson:"time"`
	Name  string                   `bson:"name"`
	Value []map[string]interface{} `bson:"value"`
}

type CalculationResultFloat struct {
	Time  string  `bson:"time"`
	Name  string  `bson:"name"`
	Value float64 `bson:"value"`
}
type OpcUpdateTime struct {
	UpdateTime string `bson:"update_time"`
	Group      string `bson:"group"`
}

type OpcAlarm struct {
	Name  string `bson:"name"`
	Type  string `bson:"type"`
	Time  string `bson:"time"`
	State int    `bson:"state"` //0未处理；1已处理
}

type MongoAlarmList struct {
	Time  string     `bson:"time"`
	Name  string     `bson:"name"`
	Value []OpcAlarm `bson:"value"`
}

type Alarm struct {
	Name string
	Type string
}
