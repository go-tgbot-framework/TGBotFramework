// Telegram Bot 框架：使用包含模組功能的 Bot 框架，不但能用比
// 原本更短的時間在架設機器人上，還能隨時下載到社群製作的實用
// 模組！
package main

import (
    "fmt"
    "os"
    _ "./tgbotframework"
    "io/ioutil"
    "encoding/json"
)

// settings.json 設定檔解析出的資料存放位置
// Settings 建構體 -> TGBot_Consts.go
var JSONData = new(Settings)

// input: 類似 Python 的 input() 函式
func input(prompt string) string {
    var action string

    fmt.Print(prompt)
    fmt.Scanln(&action)

    return action
}

// 主選單
func intro() {
    // VERSION, CONTRIBUTOR 常數 -> TGBot_Consts.go
    // introTxt 字串 -> TGBot_Strings.go
    var usrinput = input(fmt.Sprintf(introTxt, VERSION, CONTRIBUTOR))
    switch usrinput {
    case "1":
        // turnBot() 函式 -> TGBot_BotRunner.go
        turnBot()
        return
    case "2":
        // setUpBot() 函式 -> TGBot_BotSettings.go
        setUpBot()
        return
    case "3":
        fmt.Println("尚未完工～")
        break
    case "4":
        os.Exit(0)
        break
    default:
        // usrInputWrong 字串 -> TGBot_Strings.go
        fmt.Printf(usrInputWrong, usrinput)
        intro()
        return
    }
}

func main() {
    // settingsParsing 字串 -> TGBot_Strings.go
    fmt.Printf(settingsParsing)
    rawJsonData, err := ioutil.ReadFile("settings.json")

    if err == nil {
        if err := json.Unmarshal(rawJsonData, &JSONData); err == nil {
            // settingsParseSuccess 字串 -> TGBot_Strings.go
            fmt.Println(settingsParseSuccess)
        } else {
            // settingsParseFailed, solveWay_settingIsInvaild 字串 -> TGBot_Strings.go
            fmt.Printf(settingsParseFailed, solveWay_settingIsInvaild)
            os.Exit(1)
        }
    } else {
        // settingsParseFailed, solveWay_settingNotFound 字串 -> TGBot_Strings.go
        fmt.Printf(settingsParseFailed, solveWay_settingNotFound)
        os.Exit(1)
    }
    
    // 進入主畫面
    intro()
}
