// Telegram Bot 框架：使用包含模組功能的 Bot 框架，不但能用比
// 原本更短的時間在架設機器人上，還能隨時下載到社群製作的實用
// 模組！
//
// (!) 僅中文註釋 | Only Chinese Comment
//
// [使用方法]
//
// 1. 將此 API 檔案放置到 [程式目錄]/tgbotframework
//
// 2. 匯入模組：import ./tgbotframework
//
// 3. 您可使用此處提供的函式庫！
//
// 參閱 docs/Modules/Module_Basic[(您熟悉的語言)].md 說明檔案
// 得知更多資訊！:D
package TGBotLib

// 基本函式：接收訊息函式 (最基本模式)
//
// 請參考文件：
// https://core.telegram.org/bots/api#getting-updates
//
// token: 機器人 (從 @botFather 取得的) Token
//
// clean_prev_msg: 是否透過設定上一個接收的 offset 來防止已抓取訊息再次出現。
// 可參閱：https://core.telegram.org/bots/faq#long-polling-gives-me-the-same-updates-again-and-again
//
// 回傳內容：伺服器收到回應後傳回訊息。
func GetUpdatesBasic(token string, clean_prev_msg bool) *Update {
    return GetUpdates(token, -1, -1, -1, clean_prev_msg)
}

// 基本函式：傳送訊息函式 (最基本模式)
//
// 請參考文件：
// https://core.telegram.org/bots/api#sendmessage
//
// token: 機器人 (從 @botFather 取得的) Token
//
// 回傳內容：伺服器收到回應後傳回訊息。
func SendMessageBasic(token string, chat_id int, text string) string {
    return SendMessage(token, chat_id, text, "", false, false, -1)
}
