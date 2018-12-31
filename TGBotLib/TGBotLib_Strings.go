// 字串區：您可自由編輯此處字串，但每次修改過後必須重新編譯。
package TGBotLib

// HTTPCodeError 字串：若伺服器回傳代碼非 200 則回傳此字串。
// %d (1): HTTP 代碼
// %s (2): 伺服器回傳的相關訊息
const HTTPCodeError = "HTTP %d 錯誤：可能發生了些錯誤，若這不僅發生一次，請透過 Issue Tracker 告訴我們！\n伺服器回傳：\n%s\n"
