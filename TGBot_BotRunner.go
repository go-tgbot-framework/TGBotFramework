package main

import (
    "time"
    "fmt"
)

/* 相關變數 */

// 傳給 botRunner 的頻道 (channel)
// 當 botRunner 準備關閉時會關閉該頻道 (並傳出一個 `true`)。
var closeChecker = make(chan bool, 1)

// 傳給 botRunner 要關閉的通知。當 botRunner 開始下一輪執行時檢查到此部份，
// 將會進入關閉程序。
var isBotStart = false

/*// 模組執行部份
func moduleRunner(filename string) {
    theModule := plugin.Open(filename)

    handle := theModule.Lookup("Handler")

    handle.(func (string) {})()
}*/

// 啟動機器人函式
func botRunner(modlist []string, isClosed chan bool) {
    for {
        if isBotStart {
            fmt.Println("I am still running!!! :D")
            time.Sleep(2 * time.Second)
        } else {
            fmt.Println("Oh no. :(")
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
            go botRunner(nil, closeChecker)
            fmt.Println(botStarted)
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
            turnBot()
            return
        case "3":
            intro()
            return
    }
}
