package main

import (
    "fmt"
    "plugin"
    "time"
)

/* 相關變數 */

// 傳給 botRunner 的頻道 (channel)wwl
// 當 botRunner 準備關閉時會關閉該頻道 (並傳出一個 `true`)。
var closeChecker = make(chan bool, 1)

// 傳給 botRunner 要關閉的通知。當 botRunner 開始下一輪執行時檢查到此部份，
// 將會進入關閉程序。
var isBotStart = false

// 模組執行部份
func moduleRunner(filename string) {
    theModule, err := plugin.Open(filename)
    
    if err != nil {
        panic(moduleNotFound)
    }
    
    handle, err := theModule.Lookup("Handler")
    
    if err != nil {
        panic(moduleInvaild)
    }

    handle.(func (string))(JSONData.Token)
}

// 啟動機器人函式
func botRunner(isClosed chan bool) {
    isClosed <- false
    for {
        if isBotStart {
            // JSONData 設定 -> TGBot_Main.go
            moduleRunner("modules/" + JSONData.ModuleName)
        } else {
            isClosed <- true
            return
        }
    }
}

// 開關機器人部份
func turnBot() {
    var usrInput = input(fmt.Sprintf(turnBotIntroTxt, isBotStart))

    switch usrInput {
        case "1":
            if isBotStart == true {
                fmt.Println(alreadyStarted)
            } else {
                fmt.Print(botStarting)
                isBotStart = true
                go botRunner(closeChecker)
                for data := range closeChecker {
                    if data == false {
                        fmt.Println(botStarted)
                        break
                    }
                }
            }
            time.Sleep(1 * time.Second)
            // 重新顯示選單
            turnBot()
            return
        case "2":
            if isBotStart == false {
                fmt.Println(alreadyClosed)
            } else {
                fmt.Print(botClosing)
                isBotStart = false
                for data := range closeChecker {
                    if data == true {
                        fmt.Println(botStarted)
                        break
                    }
                }
            }
            time.Sleep(1 * time.Second)
            // 重新顯示選單
            turnBot()
            return
        case "3":
            // intro() 主選單函式 -> TGBot_Main.go
            intro()
            return
        default:
            fmt.Printf(usrInputWrong, usrInput)
            turnBot()
            return
    }
}
