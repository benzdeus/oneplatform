package one_entities

type OneChatService interface {
	// Method สำหรับการส่งข้อความในรูปแบบของ Card ที่ให้ user ได้เลือกกดใช้งาน
	//
	// title คือ หัวข้อของ template
	//
	// oneID คือ key สำหรับ user เพื่อบ่งบอกว่าเป็น user คนๆนั้น
	//
	// botID คือ key ของ bot ที่จะให้ส่งข้อความไปยัง user
	//
	// templates คือ struct สามารถอ่านเพิ่มเติมได้ที่ package one_entities.OneChatRequestElementSendComponent
	//
	// int ที่ return คือ status code ที่บ่งบอกว่าส่งข้อความสำเร็จหรือไม่
	SendComponentSelect(title string, oneID []string, botID string, templates []OneChatRequestElementSendComponent) int
}

// Struct ของการส่ง Template
//
// Image คือ url ที่เก็บรูปภาพที่นำไปแสดงผลของ template
//
// Title คือ หัวข้อของ template
//
// Detail คือ รายละเอียดของ template
//
// Choice คือ struct ย่อยที่จะบ่งบอกว่า template มีปุ่มอะไรบ้าง
//
// สามารถอ่านรายละเอียดเพิ่มเติมของ Choice ได้ที่ package one_entities.OneChatChoice
type OneChatRequestElementSendComponent struct {
	Image  string          `json:"image"`
	Title  string          `json:"title"`
	Detail string          `json:"detail"`
	Choice []OneChatChoice `json:"choice"`
}

// Struct ย่อยของการส่ง Template
//
// Url คือ link ที่พอ user กดแล้วจะ action ไปยัง api เส้นนั้น
//
// Sign คือ ค่า string ระหว่าง "true" | "false" ถ้าใส่เป็น "true" จะมีในส่วนของ shared token แนบท้ายมา
//
// Type คือ รูปแบบของการกดปุ่มถ้าเป็น "link" จะเด้งไปยังบราวเซอร์ของมือถือ ถ้าเป็น "webview" จะเด้ง popup เป็น webview ขึ้นมา
//
// Size คือขนาดของ Template ระหว่าง "tall" | "full"
//
// Label คือ title ของปุ่มนั้นๆ
type OneChatChoice struct {
	Url   string `json:"url"`
	Sign  string `json:"sign"` // ? true | false
	Size  string `json:"size"` // ? tall | full
	Type  string `json:"type"` // ? webview | link
	Label string `json:"label"`
}

type OneChatComponentResponse struct {
	Status string `json:"status"`
}
