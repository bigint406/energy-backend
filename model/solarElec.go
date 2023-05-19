package model

func CalcSolarElecGenHour(hourStr string) float64 {
	l, ok := GetOpcFloatList("ZLZ.%E6%80%BB%E5%8F%91%E7%94%B5%E9%87%8F1", hourStr) //总发电量1
	if !ok {
		return 0
	}
	l2, ok := GetOpcFloatList("ZLZ.%E6%80%BB%E5%8F%91%E7%94%B5%E9%87%8F", hourStr) //总发电量
	if !ok {
		return 0
	}
	return RightSubLeft(l) + RightSubLeft(l2)
}
