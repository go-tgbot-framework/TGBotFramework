package main

// settings.json 的資料。
type Settings struct {
    // Token: 機器人的 Token，可從 @BotFather 取得。
    Token string
    // IdentifyName: 您想要用什麼名稱稱呼您的機器人？
    IdentifyName string
    // ModulePath: 欲使用模組的檔案位置。
    ModulePath string
}
 
// 版本號碼
const VERSION = "[*SNAPSHOT*]"

// 貢獻者名單
const CONTRIBUTOR = `pan93412 <pan93412@gmail.com>, 2018`
