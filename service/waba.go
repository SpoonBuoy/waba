package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SpoonBuoy/waba/models"
)

type WabaService struct {
	business BusinessService
}

func NewWabaService() *WabaService {
	return &WabaService{}
}

func (wbs *WabaService) Listen(msg models.WabaMessageReq) {
	//listens for incoming messages
	sender := msg.From
	recipient := msg.To
	query := msg.Message

	//need to fetch the active context of the recipient
	ctx := wbs.business.GetActiveContext()
}
func (wbs *WabaService) Send() {
	//sends back the response to the message

}

func (wbs *WabaService) GetTemplate(token string, to string, from string, msg string) {
	url := fmt.Sprintf("https://graph.facebook.com/v20.0/%s/messages", from)
	body := map[string]interface{}{
		"messaging_product": "whatsapp",
		"recipient_type":    "individual",
		"to":                to,
		"type":              "text",
		"text": map[string]interface{}{
			"body": msg,
		},
	}
	bytesBody, _ := json.Marshal(body)
	auth := fmt.Sprintf("Bearer %s", token)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bytesBody))
	req.Header.Add("Authorization", auth)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		//handle err
	}

	fmt.Println(res.Body)
}
