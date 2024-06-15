package dto

type CreateBusinessReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
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
	Name     string `json:"name"`
	Password string `json:"password"`
}
