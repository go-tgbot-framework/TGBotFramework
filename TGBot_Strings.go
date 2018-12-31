// Telegram Bot 框架：使用包含模組功能的 Bot 框架，不但能用比
// 原本更短的時間在架設機器人上，還能隨時下載到社群製作的實用
// 模組！
// (!) 您能翻譯此處字串！
package main

/* ========== */
/* 介面選單部份 */
/* ========== */

// introTxt: 主畫面文字。
// 字串使用位置：TGBot_Main.go
// %s (1): 版本號碼
// %s (2): 貢獻者
var introTxt = ` ^TGBot 框架v
版號：%s
貢獻者列表：
%s

(1) 開啟/關閉 機器人
(2) 設定機器人
(3) 模組管理
(4) 關閉此程式

請輸入功能編號 (1-4)：`

// setUpBotIntroTxt: 機器人的設定介面。
// 字串使用位置：TGBot_BotSettings.go
var setUpBotIntroTxt = ` ^TGBot 框架v
[設定介面]

(1) 機器人 Token 設定
(2) 機器人的識別名稱
(3) 返回上一頁
(s) 儲存設定

請輸入功能編號 (1-3,s)：`

// turnBotIntroTxt: 機器人的開關介面。
// %s (1): 機器人目前是開或關？
// 字串使用位置：TGBot_BotRunner.go
var turnBotIntroTxt = ` ^TGBot 框架v
[開關介面]
目前狀態：%v
(true = 開啟 / false = 關閉)

(1) 開啟機器人
(2) 關閉機器人
(3) 返回上一頁

請輸入功能編號 (1-3)：`

/* ========== */
/* 解析 settings.json 部份 */
/* ========== */

// 解析時會顯示的文字。
// (!) 請保留後方空白！
// 字串使用位置：TGBot_Main.go
var settingsParsing = `正在解析 settings.json 的資料…… [解析中]   `

// 解析失敗會顯示的文字。
// (!) 請保證該文字比 settingsParsing 長，可用空白加長。
//     否則 settingsParsing 的未完全覆蓋文字將會繼續留在上方。
// %s (1): 解決方法，例如「請從 template 資料夾複製一個 settings.json 至本程式目錄。」
// 字串使用位置：TGBot_Main.go
var settingsParseFailed = "\r正在解析 settings.json 的資料…… [錯誤]    \n解決方法：%s\n"
var solveWay_settingNotFound = "請從 template 資料夾複製一個 settings.json 至本程式目錄。"
var solveWay_settingIsInvaild = "請檢查您的設定檔案是否正確且有效。"

// 解析成功會顯示的文字。
// (!) 請保證該文字比 settingsParsing 長，可用空白加長。
//     否則 settingsParsing 的未完全覆蓋文字將會繼續留在上方。
// 字串使用位置：TGBot_Main.go
var settingsParseSuccess = "\r正在解析 settings.json 的資料…… [解析成功]    \n"

/* ========== */
/* 設定介面文字部份 */
/* ========== */

// 當使用者要輸入 Token 時出現的畫面。
// %s (1): 目前的 Token
// 字串使用位置：TGBot_BotSettings.go
var setUpBotToken = "目前的 Token：%s\n機器人的 Token (僅按 Enter 代表維持預設值)："

// 當使用者要輸入識別名稱時出現的畫面。
// %s (1): 目前的識別名稱
// 字串使用位置：TGBot_BotSettings.go
var setUpBotIdentify = "目前的識別名稱：%s\n機器人的識別名稱 (僅按 Enter 代表維持預設值)："

// 正在寫入設定時顯示的文字
// 字串使用位置：TGBot_BotSettings.go
var writing = "正在寫入……"

// 寫入成功後出現的文字。
// (!) 請保證該文字比 writing 長，可用空白加長。
// //     否則 writing 的未完全覆蓋文字將會繼續留在上方。
// 字串使用位置：TGBot_BotSettings.go
var writeSucceed = "\r儲存設定值成功，將返回設定介面……"

/* ========== */
/* 機器人開關部份 */
/* ========== */

// 呼叫機器人函式時會顯示的文字。
// 字串使用位置：TGBot_BotRunner.go
var botStarting = "開啟中…"

// 呼叫完成機器人函式會顯示的文字。
// (!) 此 \b\b\b 的意思是將「中…」覆蓋成「完成。」（中文字佔兩個 bytes）
//     您可以選擇不寫三個 \b，這樣結果就是在「中…」之後增加文字。
// 字串使用位置：TGBot_BotRunner.go
var botStarted = "\b\b\b完成。"

// 要求關閉機器人函式時會顯示的文字。
// 字串使用位置：TGBot_BotRunner.go
var botClosing = "關閉中…"

/* ========== */
/* 共用錯誤文字部份 */
/* ========== */

// 若使用者輸入錯選項編號會顯示的文字。
// %s (1): 使用者所輸入錯的內容。
// 字串使用位置：TGBot_Main.go, TGBot_BotSettings.go, TGBot_BotRunner.go
var usrInputWrong = "您輸入錯誤了！您輸入了：「%s」，但並沒有該選項。\n"
