package main

import (
    "io/ioutil"
    "encoding/json"
    "time"
    "fmt"
)

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
