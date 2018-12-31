// Telegram Bot 框架：使用包含模組功能的 Bot 框架，不但能用比
// 原本更短的時間在架設機器人上，還能隨時下載到社群製作的實用
// 模組！
//
// (!) 僅中文註釋 | Only Chinese Comment
package main

import (
    "fmt"
    "time"
    "io/ioutil"
    "encoding/json"
    "os"
)

// 版本號碼
const VERSION = "[*SNAPSHOT*]"

// 貢獻者名單
const CONTRIBUTOR = `pan93412 <pan93412@gmail.com>, 2018`

// settings.json 的資料。
type Settings struct {
    // Token: 機器人的 Token，可從 @BotFather 取得。
    Token string
    // IdentifyName: 您想要用什麼名稱稱呼您的機器人？
    IdentifyName string
    // EnabledModules: 啟用的模組。
    EnabledModules []string
}

// settings.json 設定檔
var JSONData = new(Settings)

// input: 類似 Python 的 input() 函式
func input(prompt string) string {
    var action string
    
    fmt.Print(prompt)
    fmt.Scanln(&action)
    
    return action
}

// 設定機器人部份
func setUpBot() {
    var usrInput = input(setUpBotIntroTxt)
    
    switch usrInput {
        case "1":
            tokentmp := input(fmt.Sprintf(setUpBotToken, JSONData.Token))
            if tokentmp != "" {
                JSONData.Token = tokentmp
            }
            setUpBot()
            return
        case "2":
            identifytmp := input(fmt.Sprintf(setUpBotIdentify, JSONData.IdentifyName))
            if identifytmp != "" {
                JSONData.IdentifyName = identifytmp
            }
            setUpBot()
            return
        case "3":
            intro()
            return
        case "s":
            fmt.Print(writing)
            // 四個空白縮排
            data, errWhenMarshal := json.MarshalIndent(JSONData, "", "    ")
            if errWhenMarshal != nil {
                panic(errWhenMarshal)
            }
            if errWhenWriting := ioutil.WriteFile("settings.json", data, 0755); errWhenWriting != nil {
                panic(errWhenWriting)
            }
            time.Sleep(1 * time.Second)
            fmt.Println(writeSucceed)
            setUpBot()
            return
        default:
            fmt.Printf(usrInputWrong, usrInput)
            setUpBot()
            return
    }
}

// 主選單
func intro() {
    var usrinput = input(fmt.Sprintf(introTxt, VERSION, CONTRIBUTOR))
    switch usrinput {
        case "1":
            fmt.Println("OAO")
            break
        case "2":
            setUpBot()
            return
        case "3":
            fmt.Println("OUO")
            break
        case "4":
            os.Exit(0)
            break
        default:
            fmt.Printf(usrInputWrong, usrinput)
            intro()
            return
    }
}

func main() {
    fmt.Printf(settingsParsing)
    rawJsonData, err := ioutil.ReadFile("settings.json")
    
    if err == nil {
        if err := json.Unmarshal(rawJsonData, &JSONData); err == nil {
            fmt.Println(settingsParseSuccess)
        } else {
            fmt.Printf(settingsParseFailed, solveWay_settingIsInvaild)
            os.Exit(1)
        }
    } else {
        fmt.Printf(settingsParseFailed, solveWay_settingNotFound)
        os.Exit(1)
    }
    
    intro()
}
