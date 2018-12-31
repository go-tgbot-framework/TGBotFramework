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
package tgbotframework

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    sc "strconv"
    "strings"
)

// Telegram 的 API 網址。
//
// %s(1): Token
//
// %s(2): 方法
//
const APIURL = "https://api.telegram.org/bot%s/%s"

// 基本函式：接收訊息函式
//
// 請參考文件：
// https://core.telegram.org/bots/api#getting-updates
//
// token: 機器人 (從 @botFather 取得的) Token
//
// offset: 若不打算設定 offset，請傳入 -1。
//
// limit: 若不打算設定 limit，請傳入 -1。
//
// timeout: 若不打算設定 timeout，請傳入 -1。
//
// clean_prev_msg: 是否透過設定上一個接收的 offset 來防止已抓取訊息再次出現。
// 可參閱：https://core.telegram.org/bots/faq#long-polling-gives-me-the-same-updates-again-and-again
//
// 回傳內容：伺服器收到回應後傳回訊息。
func GetUpdates(token string, offset, limit, timeout int, clean_prev_msg bool) *Update {
    var update_obj = new(Update)
    var params = make(url.Values)
    if offset != -1 {
        params.Add("offset", sc.Itoa(offset))
    }
    if limit != -1 {
        params.Add("limit", sc.Itoa(limit))
    }
    if timeout != -1 {
        params.Add("timeout", sc.Itoa(timeout))
    }

    err := json.Unmarshal([]byte(URLGet(fmt.Sprintf(APIURL, token, "getUpdates"), params)), &update_obj)
    if err != nil {
        panic(err)
    }

    // 清除先前訊息吧 :D
    if clean_prev_msg && len(update_obj.Result) > 0 {
        GetUpdates(token, update_obj.Result[len(update_obj.Result)-1].UpdateID+1, -1, -1, false)
    }

    return update_obj
}

// 基本函式：傳送訊息函式
//
// 請參考文件：
// https://core.telegram.org/bots/api#sendmessage
//
// token: 機器人 (從 @botFather 取得的) Token
//
// parse_mode 可以僅傳入 ""，代表維持預設值
//
// reply_to_message_id 可以僅傳入 -1，代表不回覆任何人
//
// 回傳內容：伺服器收到回應後傳回訊息。
func SendMessage(
    token string, chat_id int, text string, parse_mode string,
    disable_web_page_preview, disable_notification bool, reply_to_message_id int) string {
    var params = make(url.Values)
    params.Add("chat_id", sc.Itoa(chat_id))
    params.Add("text", text)
    if parse_mode != "" {
        params.Add("parse_mode", parse_mode)
    }
    params.Add("disable_web_page_preview", sc.FormatBool(disable_web_page_preview))
    params.Add("disable_notification", sc.FormatBool(disable_notification))
    if reply_to_message_id != -1 {
        params.Add("reply_to_message_id", sc.Itoa(reply_to_message_id))
    }

    return URLGet(fmt.Sprintf(APIURL, token, "sendMessage"), params)
}

// 基本函式：轉傳訊息函式
//
// 請參考文件：
// https://core.telegram.org/bots/api#forwardmessage
//
// token: 機器人 (從 @botFather 取得的) Token
//
// 回傳內容：伺服器收到回應後傳回訊息。
func ForwardMessage(token string, chat_id, from_chat_id int, disable_notification bool, message_id int) string {
    var params = make(url.Values)
    params.Add("chat_id", sc.Itoa(chat_id))
    params.Add("from_chat_id", sc.Itoa(from_chat_id))
    params.Add("disable_notification", sc.FormatBool(disable_notification))
    params.Add("message_id", sc.Itoa(message_id))

    return URLGet(fmt.Sprintf(APIURL, token, "forwardMessage"), params)
}

// 基本函式：傳送媒體函式
//
// (!) 因為官方提供之媒體方法之多，因此我們僅提供這些方法共同的
//     功能，因此功能無法如此齊全。
//     但您仍能自己實現自己的媒體函式。
//
// 請參考文件：
// https://core.telegram.org/bots/api#send(The below method)
//
// 可用的方法：
// 請傳入您想要的功能進 `method` 引數。
//
// 傳送圖片：photo
//
// 傳送音訊：audio
//
// 傳送影片：video
//
// 傳送文件：document
//
// fileurl: 文件中的 photo, audio, video, document... 參數 (檔案網址)
//
// token: 機器人 (從 @botFather 取得的) Token
//
// caption 可以僅傳入 ""，代表不要傳送說明文字
//
// reply_to_message_id 可以僅傳入 -1，代表不回覆任何人
//
// 回傳內容：伺服器收到回應後傳回訊息。
func SendMedia(token, method string, chat_id int, fileurl, caption string,
    disable_notification bool, reply_to_message_id int) string {
    var params = make(url.Values)
    // urlmethod 為 method 參數的第一字大寫加上其剩餘小寫字。
    var urlMethod string = strings.ToUpper(string(method[0])) + string(method[1:])
    params.Add("chat_id", sc.Itoa(chat_id))
    params.Add(method, fileurl)
    params.Add("caption", caption)
    params.Add("disable_notification", sc.FormatBool(disable_notification))
    if reply_to_message_id != -1 {
        params.Add("reply_to_message_id", sc.Itoa(reply_to_message_id))
    }

    return URLGet(fmt.Sprintf(APIURL, token, "send"+urlMethod), params)
}

// 基本函式：URL GET 工具
//
// 使用方法：[連結] [參數]
//
// 回傳：得到內容
func URLGet(url string, param url.Values) string {
    var urlParams = param.Encode()
    // URL Format: (URL)?(Params)
    resp, err := http.Get(url + "?" + urlParams)

    if err != nil {
        log.Fatal(err)
    }

    if resp.StatusCode != 200 {
        rawcontent := resp.Body
        content, err_statuscode := ioutil.ReadAll(rawcontent)
        if err_statuscode != nil {
            log.Fatal(err_statuscode)
        }
        rawcontent.Close()
        log.Fatalf(HTTPCodeError, resp.StatusCode, content)
    }

    data := resp.Body
    cont, err2 := ioutil.ReadAll(data)

    if err2 != nil {
        log.Fatal(err2)
    }

    data.Close()
    return string(cont)
}
