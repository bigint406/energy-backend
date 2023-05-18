package defs

type LouHeat struct {
	Name string `bson:"name"`
	CF   string `bson:"CF"`
	InT  string `bson:"inT"`
	OutT string `bson:"outT"`
}

type LouHeatList struct {
	Time string    `bson:"time"`
	Info []LouHeat `bson:"info"`
}

type LouSolarWaterList struct {
	Time string        `bson:"time"`
	Info LouSolarWater `bson:"info"`
}

// GA的数据
type LouSolarWater struct {
	Heater_1_1  LouSolarWaterStatus      `bson:"Heater_1_1"`
	Heater_1_2  LouSolarWaterStatus      `bson:"Heater_1_2"`
	Heater_1_3  LouSolarWaterStatus      `bson:"Heater_1_3"`
	Heater_1_4  LouSolarWaterStatus      `bson:"Heater_1_4"`
	Heater_1_5  LouSolarWaterStatus      `bson:"Heater_1_5"`
	Heater_2_1  LouSolarWaterStatus      `bson:"Heater_2_1"`
	Heater_2_2  LouSolarWaterStatus      `bson:"Heater_2_2"`
	Heater_2_3  LouSolarWaterStatus      `bson:"Heater_2_3"`
	Heater_2_4  LouSolarWaterStatus      `bson:"Heater_2_4"`
	Heater_2_5  LouSolarWaterStatus      `bson:"Heater_2_5"`
	Heater_3_1  LouSolarWaterStatus      `bson:"Heater_3_1"`
	Heater_3_2  LouSolarWaterStatus      `bson:"Heater_3_2"`
	Heater_3_3  LouSolarWaterStatus      `bson:"Heater_3_3"`
	Heater_3_4  LouSolarWaterStatus      `bson:"Heater_3_4"`
	Heater_3_5  LouSolarWaterStatus      `bson:"Heater_3_5"`
	Heater_4_1  LouSolarWaterStatus      `bson:"Heater_4_1"`
	Heater_4_2  LouSolarWaterStatus      `bson:"Heater_4_2"`
	Heater_4_3  LouSolarWaterStatus      `bson:"Heater_4_3"`
	Heater_4_4  LouSolarWaterStatus      `bson:"Heater_4_4"`
	Heater_4_5  LouSolarWaterStatus      `bson:"Heater_4_5"`
	CollectHeat LouSolarWaterCollectHeat `bson:"CollectHeat"`
	HRPump_1    LouSolarWaterStatus      `bson:"HRPump_1"`
	HRPump_2    LouSolarWaterStatus      `bson:"HRPump_2"`
	JRPump_1    LouSolarWaterStatus      `bson:"JRPump_1"`
	JRPump_2    LouSolarWaterStatus      `bson:"JRPump_2"`
	System      LouSolarWaterSystem      `bson:"System"`
}

// 集热器温度
type LouSolarWaterCollectHeat struct {
	HT string `bson:"HT"`
	LT string `bson:"LT"`
}

// 运行状态
type LouSolarWaterStatus struct {
	Sta string `bson:"Sta"`
}

// GA系统信息
type LouSolarWaterSystem struct {
	JRQ_T string `bson:"JRQ_T"`
}

// H的数据，根据回风温度估测走廊温度
type LouH struct {
	Info LouHInfo `bson:"info"`
}

type LouHInfo struct {
	//D1组团
	D1_XRH_L1_2 LouHXRH `bson:"D1_XRH_L1_2"`
	D1_XRH_L2_1 LouHXRH `bson:"D1_XRH_L2_1"`
	D1_XRH_L2_2 LouHXRH `bson:"D1_XRH_L2_2"`
	D1_XRH_L2_3 LouHXRH `bson:"D1_XRH_L2_3"`
	//D2组团
	D2_XRH_B1_1 LouHXRH `bson:"D2_XRH_B1_1"`
	D2_XRH_B1_2 LouHXRH `bson:"D2_XRH_B1_2"`
	D2_XRH_B1_3 LouHXRH `bson:"D2_XRH_B1_3"`
	D2_XRH_B1_4 LouHXRH `bson:"D2_XRH_B1_4"`
	D2_XRH_B1_5 LouHXRH `bson:"D2_XRH_B1_5"`
	D2_XRH_B1_6 LouHXRH `bson:"D2_XRH_B1_6"`
	D2_XRH_B1_7 LouHXRH `bson:"D2_XRH_B1_7"`
	D2_XRH_B1_8 LouHXRH `bson:"D2_XRH_B1_8"`
	//D3组团
	D3_XHR_L2_1 LouHXRH `bson:"D3_XHR_L2_1"`
	D3_XRH_L2_2 LouHXRH `bson:"D3_XRH_L2_2"`
	D3_XRH_L3_1 LouHXRH `bson:"D3_XRH_L3_1"`
	//D5组团
	D5_XRH_L2_1 LouHXRH `bson:"D5_XRH_L2_1"`
	D5_XRH_L3_1 LouHXRH `bson:"D5_XRH_L3_1"`
	D5_XRH_B1_1 LouHXRH `bson:"D5_XRH_B1_1"`
	//南区
	GS_XRH_S_B1_2 LouHXRH `bson:"GS_XRH_S_B1_2"`
	GS_XRH_S_B2_1 LouHXRH `bson:"GS_XRH_S_B2_1"`
	GS_XRH_S_L3_1 LouHXRH `bson:"GS_XRH_S_L3_1"`
	GS_XRH_S_L3_2 LouHXRH `bson:"GS_XRH_S_L3_2"`
	GS_XRH_S_L4_1 LouHXRH `bson:"GS_XRH_S_L4_1"`
	//北区
	GN_XRH_N_L2_1 LouHXRH `bson:"GN_XRH_N_L2_1"`
}

// H中的组团热回收新风机组
type LouHXRH struct {
	HF_T string `bson:"HF_T"`
}
