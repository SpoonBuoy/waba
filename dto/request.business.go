package dto

type CreateBusinessReq struct {
	Name                string `json:"businessName"`
	BusinessPhoneNumber string `json:"waNumber"`
	Email               string `json:"email"`
	Password            string `json:"password"`
	Phone               string `json:"phone"`
}
type CreateCtxReq struct {
	Content string `json:"content"`
}
type SwitchActiveCtxReq struct {
	ContextId uint `json:"context_id"`
}
type AddWabaCredsReq struct {
	WhatsappNumber string `json:"wa_number"`
}
type LoginReq struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}
