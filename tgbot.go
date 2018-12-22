// Telegram Bot 框架：使用包含模組功能的 Bot 框架，不但能用比
// 原本更短的時間在架設機器人上，還能隨時下載到社群製作的實用
// 模組！
//
// (!) 僅中文註釋 | Only Chinese Comment
package TGBotFramework

// 機器人的基本資料，例如 Token 之類。
type Bot struct {
    // Nickname: 您想要用什麼名稱稱呼您的機器人？
    Nickname string
    // Token: 機器人的 Token，可從 @BotFather 取得。
    Token string
}

// 模組的基本資訊。
type ModuleInfo struct {
    // Name: 模組名稱，例如：「傳訊模組」。
    Name string
    // ID: 模組的 ID (不可與其他模組的 ID 相同)，例如：「sendMsgMod」
    ID string
    // Description: 模組的描述。例如：「可以傳訊給某人的模組」
    Description string
}

// 模組基本介面：使用以下的基礎介面來架構出自己的
// 使用方法：[連結] [參數]機器人模組吧！
//
// (!) 製作模組需要一定的 Go 能力。但預設已內建多項模組可供使用。
type BasicModule interface {
    // Settings(): 設定介面。將會呼叫此函數來設定機器人。
    Settings(Bot)
    // Handler(): 接收訊息的函式。將會將內容丟進去這個函式。您只須負責解析。
    Handler(ChatData)
    // Information(): 回傳這模組的資訊。
    Information() ModuleInfo
}

// 基本函式：傳送訊息函式
//
// 請參考文件：
// https://core.telegram.org/bots/api#sendmessage
func SendMessage(chat_id int, message string, parse_mode string, disable_web_page_preview bool, disable_notification bool, reply_to_message_id int) {}

// 基本函式：轉傳訊息函式
// 請參考文件：
// https://core.telegram.org/bots/api#forwardmessage
func ForwardMessage(chat_id int, from_chat_id int, disable_notification bool, message_id int) {}

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
// url = 文件中的 photo, audio, video, document... 參數
func SendMedia(method string, chat_id int, url string, caption string, disable_notification bool, reply_to_message_id int) {}

// 基本函式：URL GET 工具
// 使用方法：[連結] [參數]
