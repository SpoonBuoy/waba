package controllers

type ChatController struct{}

func NewChatController() *ChatController {
	return &ChatController{}
}

// verify api
func (wbc *ChatController) Verify() {

}

// webhook listening new message events
func (wbc *ChatController) Listen() {

}
