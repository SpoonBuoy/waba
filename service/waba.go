package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SpoonBuoy/waba/dto"
	"github.com/gin-gonic/gin"
)

const token = "EAAnrp3AqNGwBO2sfhTKhwMa26jzjqGKFVLydCH379RJ7TAGpTKgL9TRZBOlVdOZAGhJ9yAqeN2JVgKKDtYtwO3ul0mzZAKHPc8PWQDpocW9eIlrnDSCt5SIJ0eAXYVoS1Pcdit0eZAJVFZBpmjnujn0Fx835bHZCE3ZCTaP2JmZAu6Q3lKEwwbVvpmnZB5HPf46S2nNCrNZCrgsLqjN7oF2WQZD"

type WabaService struct {
	business BusinessService
	llm      LlmService
}

func NewWabaService(businessSrv BusinessService, llmSrv LlmService) *WabaService {
	return &WabaService{
		business: businessSrv,
		llm:      llmSrv,
	}
}

func (wbs *WabaService) Listen(ctx *gin.Context, payload *dto.WAMessagePayload, bid uint) error {
	//listens for incoming messages
	//get sender, reciepient, message

	//need to fetch the active context of the recipient
	activeContext, err := wbs.business.GetActiveContext(bid)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "no active context found"})
	}
	usrMsg := payload.Entry[0].Changes[0].Value.Messages[0].Text.Body
	//get response from llm
	res := wbs.llm.GetResponse(activeContext.Content, "details", usrMsg)

	//prepare whatsapp response as http req with llm res
	fmt.Println("in service layer")
	fmt.Println(payload)
	req := wbs.GetTemplateReq(token, payload.Entry[0].Changes[0].Value.Messages[0].From, payload.Entry[0].Changes[0].Value.Metadata.PhoneNumberID, res)

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
		print(err.Error())
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
