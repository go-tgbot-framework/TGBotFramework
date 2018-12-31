package tgbotframework

// 基本建構體：Update
//
// 請參考文件：https://core.telegram.org/bots/api#update
//
// 為讓部份建構體可被導出，因此部份名稱有修改過。請參考 `json:"xxx"`
// 得知原本代表的項目。
type Update struct {
    Result []struct {
        UpdateID int `json:"update_id"`
        Message  Message `json:"message"`
        EditedMessage Message `json:"edited_message"`
        ChannelPost Message`json:"channel_post"`
        EditedChannelPost Message`json:"edited_channel_post`
    }
}

// 基本建構體：Message
//
// 請參考文件：https://core.telegram.org/bots/api#message
//
// 為讓部份建構體可被導出，因此部份名稱有修改過。請參考 `json:"xxx"`
// 得知原本代表的項目。
//
// (!) 因功能繁多，僅實做部份功能。未來將會陸續補齊 [TODO]
type Message struct {
    MessageID int `json:"message_id"`
    From User `json:"from"`
    Date int `json:"date"`
    ForwardFrom Chat `json:"forward_from"`
    ForwardFromChat Chat `json:"forward_from_chat"`
    ForwardDate int `json:"forward_date"`
    Text string `json:"text"`
    // ReplyToMessage Message `json:"reply_to_message"` // [TODO] 不允許重複 Struct。待修復
}

// 基本建構體：User
//
// 請參考文件：https://core.telegram.org/bots/api#user
//
// 為讓部份建構體可被導出，因此部份名稱有修改過。請參考 `json:"xxx"`
// 得知原本代表的項目。
type User struct {
    ID int `json:"id"`
    IsBot bool `json:"is_bot"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Username string `json:"username"`
    LanguageCode string `json:"language_code"`
}

// 基本建構體：Chat
//
// 請參考文件：https://core.telegram.org/bots/api#chat
//
// 為讓部份建構體可被導出，因此部份名稱有修改過。請參考 `json:"xxx"`
// 得知原本代表的項目。
//
// (!) 因功能繁多，僅實做部份功能。未來將會陸續補齊 [TODO]
type Chat struct {
    ID int `json:"id"`
    Type string `json:"type"`
    Title string `json:"title"`
    Username string `json:"username"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    AllMembersAreAdmins bool `json:"all_members_are_administrators"` // 簡化長度
} 
