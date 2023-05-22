package energyConfig

import "energy/model"

func InitConfig() {
	load := model.GetTotalLoad()

	var dayLoad [24]float64
	var weekLoad [7][24]float64

	for i := 0; i < 24; i++ {
		dayLoad[i] = load[i]
	}

	for i := 0; i < 7; i++ {
		for j := 0; j < 24; j++ {
			weekLoad[i][j] = load[i*24+j]
		}
	}

	Energy.Daily_load_prediction = dayLoad
	EnergyWeekly.Week_load_prediction = weekLoad
}
