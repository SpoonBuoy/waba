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
	llm      LlmService
}

func NewWabaService() *WabaService {
	return &WabaService{}
}

func (wbs *WabaService) Listen(msg models.WabaRequest) {
	//listens for incoming messages
	//get sender, reciepient, message

	//need to fetch the active context of the recipient
	ctx := wbs.business.GetActiveContext()

	//get response from llm
	res = wbs.llm.GetResponse()

	//prepare whatsapp respnse with llm res
	ans = wbs.GetTemplate()

	//return ans
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
