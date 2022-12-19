# oneplatform
ใช้สำหรับเชื่อมต่อ ONE Service ปัจจุบันมี 2 Service
- One ID
- One Chat


## วิธีการใช้งาน

### ติดตั้ง
```bash
go get https://github.com/benzdeus/oneplatform
```

### Initial ตัวแปร

```go
// สร้าง env ของ one service
options := oneplatform.Options{
    ClientID:      0,
    ClientSecret:  "client_secret",
    OneChatAPIKey: "one_chat_api_key",
}

// service
service := oneplatform.Init(options)
```

### One ID

```go
// Login
resLogin := service.OneID.LoginPWD("username", "password")

// ดึงข้อมูลผู้ใช้งาน
resPersonData := service.OneID.GetAccountData(resLogin.AccessToken)

log.Printf("%+v", resPersonData.FirstNameTh)
```

### One Chat

```go
botID := "xxxxx"
oneID := "xxx"

choices := []one_entities.OneChatChoice{
    {
        Sign:  "true",
        Url:   "https://google.co.th?t=",
        Size:  "full",
        Type:  "link",
        Label: "ตกลง",
    },
    {
        Sign:  "true",
        Url:   "https://google.co.th?t=",
        Size:  "full",
        Type:  "link",
        Label: "ไม่ตกลง",
    },
}

template := []one_entities.OneChatRequestElementSendComponent{
    {
        Image:  "https://animatebkk-online.com/wp-content/uploads/2022/07/4580590170247-3.jpg",
        Title:  "หัวข้อ Component",
        Detail: "กดหน่อยครับ",
        Choice: choices,
    },
}
service.OneChat.SendComponentSelect("หัวข้อแจ้งเตือน", oneID, botID, template)

```