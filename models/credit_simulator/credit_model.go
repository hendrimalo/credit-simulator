package models

type CreditSimulator struct {
	VehicleType      string  `json:"vehicle_type"`
	VehicleCondition string  `json:"vehicle_condition"`
	Year             int     `json:"year"`
	DownPayment      float32 `json:"down_payment"`
	Total            float32 `json:"total"`
	Tenor            int     `json:"tenor"`
}

type CreditSimulatorRespon struct {
	Tahun              int     `json:"tahun"`
	Rate               float32 `json:"rate"`
	InstallmentPayment float32 `json:"installment_payment"`
}
