package main

import (
    "time"
    "fmt"
)

/* 相關變數 */

// 傳給 botRunner 的頻道 (channel)ww
// 當 botRunner 準備關閉時會關閉該頻道 (並傳出一個 `true`)。
var closeChecker = make(chan bool, 1)

// 傳給 botRunner 要關閉的通知。當 botRunner 開始下一輪執行時檢查到此部份，
// 將會進入關閉程序。
var isBotStart = false

// 模組執行部份
func moduleRunner(filename string) {
    theModule := plugin.Open(filename)

    handle := theModule.Lookup("Handler")

    handle.(func (string) {})()
}

// 啟動機器人函式
func botRunner(isClosed chan bool) {
    for {
        if isBotStart {
            // JSONData 設定 -> TGBot_Main.go
            moduleRunner(JSONData.ModulePath)
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
            fmt.Print(botStarting)
            time.Sleep(1 * time.Second)
            isBotStart = true
            go botRunner(closeChecker)
            fmt.Println(botStarted)
            // 重新執行一次此函式
            turnBot()
            return
        case "2":
            fmt.Print(botClosing)
            time.Sleep(1 * time.Second)
            isBotStart = false
            for data := range closeChecker {
                if data == true {
                    fmt.Println(botStarted)
                    break
                }
            }
            // 重新執行一次此函式
            turnBot()
            return
        case "3":
            // intro() 主選單函式 -> TGBot_Main.go
            intro()
            return
    }
}
