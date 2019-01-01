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
func setUpMod(filename string) {
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
    fmt.Println(modifyModuleHelpTxt)
    folderInf, err := ioutil.ReadDir(ModulesPath)

    if err != nil {
        panic(modulesPathNotFound)
    }

    for _, folderInfo := range folderInf {
        if strings.Contains(folderInfo.Name(), ".so") {
            fmt.Printf("模組 | %s (%s)\n", folderInfo.Name(), getModInfo(ModulesPath + folderInfo.Name())["Name"])
        }
    }

    // JSONData -> TGBot_Init.go
    fmt.Printf(currentModule, JSONData.ModuleName)
    moduleName := input(enterModuleName)

    if moduleName != "" {
        if _, err := os.Stat(ModulesPath + moduleName); err != nil {
            fmt.Printf(modNotFound, moduleName)
            moduleControl()
            return
        }

        confirm := input(fmt.Sprintf(confirmModIsCorrect, getModInfo(ModulesPath + moduleName)["Name"]))
        if confirm == "Y" || confirm == "" {
            // JSONData -> TGBot_Main.go
            JSONData.ModuleName = moduleName
            // writeSettings() -> TGBot_BotSettings.go
            writeSettings()
        }
    }

    moduleControl()
    return
}

// 設定模組部份
func moduleControl() {
    usrInput := input(fmt.Sprintf(moduleControlIntroTxt, JSONData.ModuleName))

    switch usrInput {
        case "1":
            modifyModule()
            moduleControl()
            return
        case "2":
            setUpMod((ModulesPath + JSONData.ModuleName))
            moduleControl()
            return
        case "3":
            modInfo := getModInfo(ModulesPath + JSONData.ModuleName)
            fmt.Printf(moduleInfo, JSONData.ModuleName, modInfo["Name"], modInfo["Version"],
                       modInfo["Author"], modInfo["Description"])
            moduleControl()
            return
        case "4":
            intro()
            return
        default:
            fmt.Printf(usrInputWrong, usrInput)
            moduleControl()
            return
    }
}
