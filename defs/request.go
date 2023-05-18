package defs

type BasicDataSet struct {
	BasicData          []string `json:"basic_data"`
	BasicDataListYear  []string `json:"basic_data_list_year"`
	BasicDataListMonth []string `json:"basic_data_list_month"`
	BasicDataListDay   []string `json:"basic_data_list_day"`
	BasicDataListHour  []string `json:"basic_data_list_hour"`
	BasicOpcList       []string `json:"basic_opc_list"`
	MapDataListDay     []string `json:"map_data_list_day"`
}

type BasicDataSetRequest struct {
	Data     BasicDataSet `json:"data"`
	DayStr   string       `json:"day_str"`
	HourStr  string       `json:"hour_str"`
	MonthStr string       `json:"month_str"`
	YearStr  string       `json:"year_str"`
}
