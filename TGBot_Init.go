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
    // 檢查 modules 資料夾是否存在。
    fmt.Print("正在檢查 modules 資料夾……")
    modulesFolder, errWhenStat := os.Stat(ModulesPath)
    if errWhenStat != nil {
        fmt.Println("\rmodules 資料夾不存在。將自動建立。   ")
        errMkdir := os.Mkdir(ModulesPath, 0755)
        if errMkdir != nil {
            panic("請確保此程式所在資料夾可供寫入。")
            fmt.Println("請重新啟動程式套用設定。")
            os.Exit(1)
        }
    } else if !modulesFolder.IsDir() {
        log.Println("\rmodules 不是個資料夾。將移除現有的 modules 檔案後重新建立。   ")
        errRM := os.Remove(ModulesPath)
        errMkdir := os.Mkdir(ModulesPath, 0755)
        if errRM != nil || errMkdir != nil {
            panic("請確保此程式所在資料夾可供寫入。")
        }
        fmt.Println("請重新啟動程式套用設定。")
        os.Exit(1)
    } else {
        fmt.Println("\rmodules 資料夾存在。        ")
    }
    
    // 檢查 settings.json 是否存在
    // settingsParsing 字串 -> TGBot_Strings.go
    fmt.Printf(settingsParsing)
    rawJsonData, errReadFile := ioutil.ReadFile(SettingsFilename)
    
    if errReadFile != nil {
        fmt.Println(settingsParseFailed)
        if !createSettings() {
            panic("請確保此程式所在資料夾可供寫入。")
        }
        fmt.Println("請重新啟動程式套用設定。")
        os.Exit(1)
    } else if errUnmarshal := json.Unmarshal(rawJsonData, &JSONData); errUnmarshal != nil {
        fmt.Println(settingsParseFailed)
        if !createSettings() {
            panic("請確保此程式所在資料夾可供寫入。")
        }
        fmt.Println("請重新啟動程式套用設定。")
        os.Exit(1)
    } else {
        fmt.Println(settingsParseSuccess)
    }
    
    return // 檢查完成。
}
