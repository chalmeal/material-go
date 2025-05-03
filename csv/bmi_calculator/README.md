**bmi_calclator**

### 概要
患者の身長と体重からBMI値を計算し、値に応じて結果(肥満度)を返す。

### 仕様
* 患者の身長と体重は、.inputフォルダに配置したCSVから一括で読み込む。
    * CSVの命名は「bmi_calc_partient_{YYYYMMDD}.csv」とし、日付が最新のCSVを取得する。
* 入力CSVの項目として以下を指定する。

| 項目論理名 | 項目物理名 |
| --- | ------ |
| ID  | id     | 
| 氏名 | name   | 
| 性別 | gender | 
| 身長 | height | 
| 体重 | weight | 

* 計算して返された結果は、.outputフォルダにCSV形式で保存する。
    * CSVの命名は「bmi_calc_result_{YYYYMMDDhhmmss}.csv」とする。
* 出力CSVの項目として以下を指定する。

| 項目論理名 | 項目物理名 |
| ---- | ------ |
| ID   | id     | 
| 氏名  | name   | 
| BMI値 | bmi   |
| 結果 | result |
