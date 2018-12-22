// Telegram Bot Framework: Build your bot easier way with modules!
//
// (!) 中英文註釋 | Chinese & English Comment
package TGBotFramework

// 機器人的基本資料，例如 Token 之類。
//
// The basic informations about your bot, like token.
//
type Bot struct {
    // Nickname: 您想要用什麼名稱稱呼您的機器人？
    // Nickname: What is the name that you want to call your bot as?
    Nickname string
    // Token: 機器人的 Token，可從 @BotFather 取得。
    // Token: The token of your bot. You can get your token by @BotFather.
    Token string
}

// 模組的基本資訊。
//
// The basic information about the module.
//
type ModuleInfo struct {
    // Name: 模組名稱，例如：「Send Message Module」。
    // Name: Module name, for example: "Send Message Module".
    Name string
    // ID: 模組的 ID (不可與其他模組的 ID 相同)，例如：「sendMsgMod」
    // ID: Module ID (You can't name the same name of other modules), for example: "sendMsgMod"
    ID string
    // Description: 模組的描述。例如：「A module which can send message to someone.」
    // Description: Module description. For example: "A module which can send message to someone."
    Description string
}

// 模組基本介面：使用以下的基礎介面來架構出自己的機器人模組吧！
// 
// (!) 製作模組需要一定的 Go 能力。但預設已內建多項模組可供使用。
//
// The module basic interface: Use the below interface to build your bot module!
//
// (!) Make module need a certain Go capability. But we have provided many
//     module to use.
//
type BasicModule interface {
    // Settings(): 設定介面。將會呼叫此函數來設定機器人。
    // Settings(): Settings interface: We will call the function to set your bot.
    Settings(Bot)
    
    // Handler(): 接收訊息的函式。將會將內容丟進去這個函式。您只須負責解析。
    Handler(ChatData)
    
    // Information(): 回傳這模組的資訊。
    Information() ModuleInfo
}

// 基本函式：傳送訊息的函式
//
// BASIC FUNCTION: A function which can send message.
//
// 請參考文件 Please refer the documentation:
// https://core.telegram.org/bots/api#sendmessage
//
func SendMessage(chat_id int, message string, parse_mode string, disable_web_page_preview bool, disable_notification bool, reply_to_message_id int) {}

// 基本函式：轉傳訊息的函式
//
// BASIC FUNCTION: A function which can forward message.
//
// 請參考文件 Please refer the documentation:
// https://core.telegram.org/bots/api#forwardmessage
//
func ForwardMessage(chat_id int, from_chat_id int, disable_notification bool, message_id int) {}

// 基本函式：傳送媒體的函式。
//
// BASIC FUNCTION: A function which can send media.
//
// (!) 因為官方提供之媒體 method 之多，因此我們僅提供這些 method 共同的
//     功能，因此功能無法如此齊全。
//     但您仍能自己實現自己的媒體函式。
// 
// (!) Because there are so many media method provided by official,
//     we just provided the common feature of these methods.
//     But you still can implement your media functions.
//
// 請參考文件 Please refer the documentation:
// https://core.telegram.org/bots/api#send(The below method)
//
// 可用的方法 AVAILABLE METHODS：
// 請傳入您想要的功能進 `method` 引數。
//
// 傳送圖片 Send Photos: photo
// 傳送音訊 Send Audios: audio
// 傳送影片 Send Videos: video
// 傳送文件 Send Documents: document
//
// url = 文件中的 photo, audio, video, document... 參數
//
// url = The photo, audio, video, document parameter in documentation.
func SendMedia(chat_id int, url string, caption string, disable_notification bool, reply_to_message_id int) {}
