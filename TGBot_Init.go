// TGBot 框架初始化工具。
package main

import (
    "log"
    "io/ioutil"
    "fmt"
    "os"
    "encoding/json"
)

// Settings 建構體
var JSONData = new(Settings)

// 建立 settings.json，回傳一個 bool，
// 若為 true 代表成功，false 代表失敗。
func createSettings() bool {
    if err := ioutil.WriteFile(SettingsFilename, []byte(SettingsJson), 0644); err != nil {
        return false
    }
    return true
}

// 初始化 Telegram Bot Framework 的函式
func TGBotInit() {
    modulesFolder, errWhenStat := os.Stat(ModulesPath)
    if errWhenStat != nil {
        fmt.Println(init_ModulesFolderNotExists)
        errMkdir := os.Mkdir(ModulesPath, 0755)
        if errMkdir != nil {
            panic(init_ReadOnly)
        } else {
            fmt.Println(init_restartToApply)
            os.Exit(1)
        }
    } else if !modulesFolder.IsDir() {
        log.Println(init_NotAFolder)
        errRM := os.Remove(ModulesPath)
        errMkdir := os.Mkdir(ModulesPath, 0755)
        if errRM != nil || errMkdir != nil {
            panic(init_ReadOnly)
        }
        fmt.Println(init_restartToApply)
        os.Exit(1)
    }
    
    // 檢查 settings.json 是否存在
    // settingsParsing 字串 -> TGBot_Strings.go
    fmt.Printf(settingsParsing)
    rawJsonData, errReadFile := ioutil.ReadFile(SettingsFilename)
    
    if errReadFile != nil {
        fmt.Println(settingsParseFailed)
        if !createSettings() {
            panic(init_ReadOnly)
        }
        fmt.Println(init_restartToApply)
        os.Exit(1)
    } else if errUnmarshal := json.Unmarshal(rawJsonData, &JSONData); errUnmarshal != nil {
        fmt.Println(settingsParseFailed)
        if !createSettings() {
            panic(init_ReadOnly)
        }
        fmt.Println(init_restartToApply)
        os.Exit(1)
    } else {
        fmt.Println(settingsParseSuccess)
    }
    
    return // 檢查完成。
}
