package components

import (
	"bmi_calclator/config"
	"encoding/csv"
	"io"
	"os"
	"strings"
)

// CSVファイル読み込み
func ReadCSV() ([][]string, error) {
	// 最新のCSVファイル名を取得
	data := make([][]string, 0)

	csvName, err := getLatestCSVName()
	if err != nil {
		return data, err
	}

	f, err := os.Open(csvName)
	if err != nil {
		return data, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	i := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return data, err
		}

		// 1行目はヘッダーなのでスキップ
		if i == 0 {
			i++
			continue
		}

		data = append(data, record)
		i++
	}

	return data, nil
}

// 最新のCSVファイル名を取得
func getLatestCSVName() (string, error) {
	// CSVディレクトリパス
	csvDirPath := config.CSV_DIRECTORY_IMPORT_PATH

	// CSVディレクトリ内のファイルを取得
	files, err := os.ReadDir(csvDirPath)
	if err != nil {
		return "", err
	}

	latestCsvFile := ""
	for _, file := range files {
		date := strings.Split(file.Name(), "_")[3]
		if (timeParse(date).After(timeParse(latestCsvFile))) || latestCsvFile == "" {
			latestCsvFile = file.Name()
		}
	}

	csvName := csvDirPath + "/" + latestCsvFile

	return csvName, nil
}

// CSVファイル作成
func CreateCSVFile(fileName string) (*csv.Writer, error) {
	// CSVディレクトリパス
	csvDirPath := config.CSV_DIRECTORY_EXPORT_PATH

	// CSVディレクトリが存在しない場合は作成
	if _, err := os.Stat(csvDirPath); os.IsNotExist(err) {
		if err := os.Mkdir(csvDirPath, 0755); err != nil {
			return nil, err
		}
	}

	// CSVファイルを作成
	f, err := os.Create(csvDirPath + "/" + fileName)
	if err != nil {
		return nil, err
	}

	// CSVファイルにヘッダーを書き込む
	writer := csv.NewWriter(f)
	defer writer.Flush()

	header := []string{"ID", "名前", "BMI", "判定"}
	if err := writer.Write(header); err != nil {
		return nil, err
	}

	return writer, nil
}

// CSVファイルにデータを書き込む
func WriteCSVFile(w *csv.Writer, results [][]string) error {
	// CSVファイルにデータを書き込む
	// 1行ずつ書き込む
	for _, result := range results {
		if err := w.Write(result); err != nil {
			return err
		}
	}
	w.Flush()

	return nil
}
