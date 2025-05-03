**kaonavi-driver**

### 概要
カオナビWebhookを利用するために、ドライバの役割を持つプログラムを作成する。

### 仕様
* 基本的な仕様はカオナビWebhookを参考
    *  [カオナビAPI - Webhook(β)](https://developer.kaonavi.jp/api/v2.0/index.html#section/Webhookb)
* Webhookで利用するイベントはコマンドで切り分ける。
    * member_created
    * member_updated
    * member_deleted
> 例：go run main.go member_created
* Debug Runからそれぞれイベントを選択することも可能
> vscodeを利用