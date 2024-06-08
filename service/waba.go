package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SpoonBuoy/waba/models"
	"github.com/gin-gonic/gin"
)

const token = "EAAnrp3AqNGwBOz2ltZCOaeWBQ5OdDm04TH1mnIibbH8aZC0AoZBVB0g1BuLNwq96DsCfe5Ub9KWvKPGUkZAmvO8REmPVUAbKfcyNMTuhN4Jq9sLfwd29SuZBRH5ZBhMa8KfDnZAMRgxbg4wCYuJOYjrB435nL64YHIt3KuGc1LhqlZAu2AN3HXjNRPrCys581QgbA6YZCLZA4cko8tXA8giZAAZD"

type WabaService struct {
	business BusinessService
	llm      LlmService
}

func NewWabaService() *WabaService {
	return &WabaService{}
}

func (wbs *WabaService) Listen(c *gin.Context, payload *models.WAMessagePayload) error {
	//listens for incoming messages
	//get sender, reciepient, message

	//need to fetch the active context of the recipient
	// ctx := wbs.business.GetActiveContext()

	//get response from llm
	// res = wbs.llm.GetResponse()

	//prepare whatsapp response as http req with llm res
	fmt.Println("in service layer")
	fmt.Println(payload)
	req := wbs.GetTemplateReq(token, payload.Entry[0].Changes[0].Value.Messages[0].From, payload.Entry[0].Changes[0].Value.Metadata.PhoneNumberID, "Hello from aws")

	//make http request to whatsapp
	wbs.Send(req)
	//return ans
	return nil
}
func (wbs *WabaService) Send(req *http.Request) {
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		//handle err
	}

	fmt.Println(res.Body)
}

func (wbs *WabaService) GetTemplateReq(token string, recipientNum string, senderPhoneNumberId string, message string) *http.Request {
	url := fmt.Sprintf("https://graph.facebook.com/v20.0/%s/messages", senderPhoneNumberId)
	body := map[string]interface{}{
		"messaging_product": "whatsapp",
		"recipient_type":    "individual",
		"to":                recipientNum,
		"type":              "text",
		"text": map[string]interface{}{
			"body": message,
		},
	}
	bytesBody, _ := json.Marshal(body)
	auth := fmt.Sprintf("Bearer %s", token)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bytesBody))
	req.Header.Add("Authorization", auth)
	req.Header.Add("Content-Type", "application/json")

	return req

}
