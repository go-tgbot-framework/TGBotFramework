package main

import (
    "io/ioutil"
    "encoding/json"
    "fmt"
)

// 寫入設定值
func writeSettings() {
    fmt.Print(writing)
    // 四個空白縮排
    data, errWhenMarshal := json.MarshalIndent(JSONData, "", "    ")
    if errWhenMarshal != nil {
        panic(errWhenMarshal)
    }
    if errWhenWriting := ioutil.WriteFile(SettingsFilename, data, 0755); errWhenWriting != nil {
        panic(errWhenWriting)
    }
    fmt.Println(writeSucceed)
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
        intro()
        return
    case "s":
        writeSettings()
        setUpBot()
        return
    default:
        fmt.Printf(usrInputWrong, usrInput)
        setUpBot()
        return
    }
}
