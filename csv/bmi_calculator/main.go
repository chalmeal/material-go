package main

import (
	"bmi_calclator/service"
	"fmt"
)

var (
	patient service.Patient
	result  service.Result
)

func init() {
	// 患者サービスDI
	patient = service.NewPatient()
	// BMI結果サービスDI
	result = service.NewResult()
}

func main() {
	// 患者情報を取得
	patientInfo, err := patient.GetPatientInfo()
	if err != nil {
		fmt.Printf("患者情報の取得に失敗しました。: %v\n", err)
		return
	}

	// BMI計算結果を取得
	bmiResults := result.GetBmiResult(patientInfo)

	// BMI計算結果をCSV出力
	err = result.ExportResultCsv(bmiResults)
	if err != nil {
		fmt.Printf("BMI計算結果のCSV出力に失敗しました。: %v\n", err)
		return
	}
}
