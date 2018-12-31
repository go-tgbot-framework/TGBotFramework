package main

import (
    "fmt"
    "plugin"
    "io/ioutil"
    "strings"
    "os"
)

// 取得模組資訊
func getModInfo(filename string) map[string]string {
    theModule, err := plugin.Open(filename)
    
    if err != nil {
        panic(moduleNotFound)
    }
    
    info, err := theModule.Lookup("Info")
    
    if err != nil {
        panic(moduleInvaild)
    }

    return info.(func () map[string]string)()
}

// 模組設定區塊
func setMod(filename string) {
    theModule, err := plugin.Open(filename)
    
    if err != nil {
        panic(moduleNotFound)
    }
    
    info, err := theModule.Lookup("Settings")
    
    if err != nil {
        panic(moduleInvaild)
    }

    info.(func ())()
}

// 變更模組
func modifyModule() {
    fmt.Println(setUpBotIntroTxt)
    folderInf, err := ioutil.ReadDir("./modules")
    
    if err != nil {
        panic("\nmodules 資料夾不存在，請在執行本程式的地方建立該資料夾，之後下載您所需的模組。")
    }
    
    for _, folderInfo := range folderInf {
        if strings.Contains(folderInfo.Name(), ".so") {
            fmt.Println(folderInfo.Name())
        }
    }
    
    // JSONData -> TGBot_Main.go
    fmt.Printf("目前使用的模組：%s\n", JSONData.ModuleName)
    moduleName := input("請輸入要使用的模組名稱 (留空代表不變更)：")
    
    if moduleName != "" {
        if _, err := os.Stat("modules/" + moduleName); err != nil {
            fmt.Printf("模組「%s」不存在。\n", moduleName)
            setUpModule()
            return
        }
        
        confirm := input(fmt.Sprintf("確定是這個模組：%s？(Y/n)：", getModInfo("modules/" + moduleName)["Name"]))
        if confirm == "Y" || confirm == "" {
            // JSONData -> TGBot_Main.go
            JSONData.ModuleName = moduleName
            // writeSettings() -> TGBot_BotSettings.go
            writeSettings()
        }
    }
    
    setUpModule()
    return
}

// 設定模組部份
func setUpModule() {
    usrInput := input(fmt.Sprintf(setUpModuleIntroTxt, JSONData.ModuleName))
    
    switch usrInput {
        case "1":
            modifyModule()
            setUpModule()
            return
        case "2":
            setMod("modules/" + JSONData.ModuleName)
            setUpModule()
            return
        case "3":
            modInfo := getModInfo("modules/" + JSONData.ModuleName)
            fmt.Printf("模組名稱：%s\n模組版本：%s\n模組作者：%s\n模組描述：%s\n\n", modInfo["Name"], modInfo["Version"],
                       modInfo["Author"], modInfo["Description"])
            setUpModule()
            return
        case "4":
            intro()
            return
        default:
            fmt.Printf(usrInputWrong, usrInput)
            setUpModule()
            return
    }
}
