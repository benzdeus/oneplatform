# oneplatform
ใช้สำหรับเชื่อมต่อ ONE Service ปัจจุบันมี 2 Service
- One ID
- One Chat


## วิธีการใช้งาน

### ติดตั้ง
```bash
go get github.com/benzdeus/oneplatform
```

### Initial ตัวแปร

```go
// สร้าง env ของ one service
options := oneplatform.Options{
    // oneid config
    ClientID:      "0",
    ClientSecret:  "client_secret",
    RefCode:       "client_secret_refcode",
    ClientIDSharedToken:      "0",
    ClientSecretSharedToken:  "client_secret",
    ClientRefCodeSharedToken: "client_secret_refcode",

    // onechat config
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

// ดึง access token จาก shared token
resPersonData := service.OneID.GetAccessToken(`xxxx`)

```

```go
//สมัครสมาชิก
requestRegister := one_entities.RequestOneIDRegister{
		AccountTitleTh:  "นาย",
		AccountTitleEng: "Mr",
		FirstNameTh:     "สามาน",
		FirstNameEng:    "saman",
		LastNameTh:      "กกกกกกกก",
		LastNameEng:     "kkkkkkkkkk",
		IdCardType:      "ID_CARD",
		IdCardNum:       "4881917526348",
		Email:           "sakjdksadjsakd@gmail.com",
		MobileNo:        "0999999999",
		BirthDate:       "1995-10-02",
		Username:        "bbkkjsjsjss",
		Password:        "V22344489s",
	}
response := service.OneID.Register(requestRegister)
log.Printf("%#+v", response)

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