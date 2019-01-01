package main

import (
    "fmt"
    "plugin"
)

/* 相關變數 */

// 傳給 botRunner 的頻道 (channel)
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

    // 詳閱 docs/Modules/Spec.md
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
            // JSONData 設定 -> TGBot_Init.go
            moduleRunner(ModulesPath + JSONData.ModuleName)
        } else {
            isClosed <- true
            return
        }
    }
}

// 開關機器人部份
func botControl() {
    var usrInput = input(fmt.Sprintf(botControlIntroTxt, isBotStart))

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

            // 重新顯示選單
            botControl()
            return
        case "2":
            if isBotStart == false {
                fmt.Println(alreadyClosed)
            } else {
                fmt.Print(botClosing)
                isBotStart = false

                // 反覆檢查 closeChecker 頻道，若回傳 true 則表示機器人已關閉。
                for data := range closeChecker {
                    if data == true {
                        // 僅為字串複用。應為「botClosed」。
                        fmt.Println(botStarted)
                        break
                    }
                }
            }

            // 重新顯示選單
            botControl()
            return
        case "3":
            // intro() 主選單函式 -> TGBot_Main.go
            intro()
            return
        default:
            fmt.Printf(usrInputWrong, usrInput)
            botControl()
            return
    }
}
