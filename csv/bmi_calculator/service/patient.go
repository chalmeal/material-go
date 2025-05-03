package service

import "bmi_calclator/components"

type Patient struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Height string `json:"height"`
	Weight string `json:"weight"`
}

// Patient構造体の初期化
func NewPatient() Patient {
	return Patient{}
}

// PatientServiceインターフェース
type PatientService interface {
	GetPatientInfo() (Patient, error)
}

// 患者情報を取得
func (p *Patient) GetPatientInfo() ([]Patient, error) {
	// CSVファイルを読み込む
	data, err := components.ReadCSV()
	if err != nil {
		return []Patient{}, err
	}

	// Patient構造体にデータを格納
	var patients []Patient
	for _, record := range data {
		np := NewPatient()
		// CSVの各列をPatient構造体のフィールドにマッピング
		np.ID = record[0]
		np.Name = record[1]
		np.Gender = record[2]
		np.Height = record[3]
		np.Weight = record[4]

		patients = append(patients, np)
	}

	return patients, nil
}
