package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/SpoonBuoy/waba/dto"
)

type LlmService struct {
	BaseAddr string
}

func NewLlmService(addr string) *LlmService {
	fmt.Printf("llm service with %s uri started", addr)
	return &LlmService{BaseAddr: addr}
}

func (llms *LlmService) Send(query string, context string) string {
	url := fmt.Sprintf("%s", llms.BaseAddr)
	body := map[string]interface{}{
		"prompt":  query,
		"context": context,
	}
	bytesBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(bytesBody))
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		//handle err
		return "ai is down at the moment"
		print(err.Error())
	}

	ans, err := io.ReadAll(res.Body)
	if err != nil {
		print(err.Error)
		return "ai is down at the moment"
	}
	var result dto.LlmResponse
	err = json.Unmarshal(ans, &result)
	if err != nil {
		print(err.Error)
		return "ai could not unmarshall"
	}

	return result.Result
}

func (llms *LlmService) GetResponse(context string, details string, query string) string {
	//get response from llm service
	res := llms.Send(query, context)
	return res
}

func (llm *LlmService) Ping() {
	//check if llm is alive
}
