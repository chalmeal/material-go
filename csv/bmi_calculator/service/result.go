package service

import (
	"bmi_calclator/components"
	"bmi_calclator/config"
	"fmt"
	"math"
	"time"
)

type Result struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Bmi    float64 `json:"bmi"`
	Result string  `json:"result"`
}

// Result構造体の初期化
func NewResult() Result {
	return Result{}
}

// ResultServiceインターフェース
type ResultService interface {
	GetBmiResult([]Patient) []Result
	ExportResultCsv() error
}

// BMI計算結果取得
func (r *Result) GetBmiResult(patients []Patient) []Result {
	var results []Result
	for _, p := range patients {
		// BMI計算
		height := float64(components.StringToInt(p.Height)) / 100
		weight := float64(components.StringToInt(p.Weight))
		bmi := weight / math.Pow(height, 2)
		// 小数点以下2桁まで
		bmi = math.Round(bmi*100) / 100

		nr := NewResult()
		nr.ID = p.ID
		nr.Name = p.Name
		nr.Bmi = bmi
		// BMI判定
		nr.Result = nr.judgeBmi()

		results = append(results, nr)
	}

	return results
}

// BMI判定
func (r *Result) judgeBmi() string {
	if r.Bmi < 18.5 {
		return "低体重"
	} else if r.Bmi < 25 {
		return "普通体重"
	} else if r.Bmi < 30 {
		return "肥満(1度)"
	} else if r.Bmi < 35 {
		return "肥満(2度)"
	} else if r.Bmi < 40 {
		return "肥満(3度)"
	} else {
		return "肥満(4度)"
	}
}

// BMI計算結果をCSV出力
func (r *Result) ExportResultCsv(results []Result) error {
	// CSVファイル名
	dateTime := time.Now().Format(config.YYYYMMDDHHMMSS)
	csvName := "bmi_calc_result_" + dateTime + ".csv"

	// CSVファイルを作成
	w, err := components.CreateCSVFile(csvName)
	if err != nil {
		return err
	}

	// CSVファイルにデータを書き込む
	rs := r.toSliceString(results)
	fmt.Println(rs)
	err = components.WriteCSVFile(w, rs)
	if err != nil {
		return err
	}

	return nil
}

// Result構造体を[][]stringに変換
func (r *Result) toSliceString(results []Result) [][]string {
	var rs [][]string
	for _, result := range results {
		// 結果をスライスに変換
		row := []string{
			result.ID,
			result.Name,
			fmt.Sprintf("%.2f", result.Bmi),
			result.Result,
		}
		rs = append(rs, row)
	}

	return rs
}
