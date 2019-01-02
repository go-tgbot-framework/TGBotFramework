package main

// settings.json 的資料。
type Settings struct {
    // Token: 機器人的 Token，可從 @BotFather 取得。
    Token string
    // ModuleName: 欲使用模組的檔案名稱 (位於 modules 資料夾)。
    ModuleName string
}

// 版本號碼
const VERSION = "0.7.0-beta"

// 貢獻者名單
const CONTRIBUTOR = `pan93412 <pan93412@gmail.com>, 2018`

// settings.json 的預設資料
const SettingsJson = `{
    "Token": "",
    "ModuleName": ""
}`

// 設定檔檔案名稱
const SettingsFilename = "settings.json"

// 模組的存放位置
const ModulesPath = "modules/"
