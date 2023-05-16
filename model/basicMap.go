package model

import (
	"energy/defs"
)

func CalcBasicMapHallwayTemp(data *defs.LouH) []float64 {
	ans := make([]float64, 8)
	zutuan := make([][]*string, 8)
	zutuan[0] = []*string{&data.D1_XRH_L1_2.HF_T, &data.D1_XRH_L2_1.HF_T, &data.D1_XRH_L2_2.HF_T, &data.D1_XRH_L2_3.HF_T}
	zutuan[1] = []*string{&data.D2_XRH_B1_1.HF_T, &data.D2_XRH_B1_2.HF_T, &data.D2_XRH_B1_3.HF_T, &data.D2_XRH_B1_4.HF_T,
		&data.D2_XRH_B1_5.HF_T, &data.D2_XRH_B1_6.HF_T, &data.D2_XRH_B1_7.HF_T, &data.D2_XRH_B1_8.HF_T}
	zutuan[2] = []*string{&data.D3_XHR_L2_1.HF_T, &data.D3_XRH_L2_2.HF_T, &data.D3_XRH_L3_1.HF_T}
	//没有D4组团的数据源，所以zutuan[3]是空的
	zutuan[4] = []*string{&data.D5_XRH_B1_1.HF_T, &data.D5_XRH_L2_1.HF_T, &data.D5_XRH_L3_1.HF_T}
	//没有D6组团的数据源，所以zutuan[5]是空的
	zutuan[6] = []*string{&data.GS_XRH_S_B1_2.HF_T, &data.GS_XRH_S_B2_1.HF_T, &data.GS_XRH_S_L3_1.HF_T, &data.GS_XRH_S_L3_2.HF_T, &data.GS_XRH_S_L4_1.HF_T} //南区
	zutuan[7] = []*string{&data.GN_XRH_N_L2_1.HF_T}                                                                                                         //北区
	// HT, err := strconv.ParseFloat(data., 64)
	// if err != nil {
	// 	return 0
	// }
	return ans
}
