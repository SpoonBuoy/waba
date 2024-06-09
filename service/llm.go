package service

type LlmService struct {
	BaseAddr string
}

func NewLlmService() *LlmService {
	return &LlmService{}
}

func (llms *LlmService) GetResponse(context string, details string, query string) string {
	//get response from llm service
	return ""
}

func (llm *LlmService) Ping() {
	//check if llm is alive
}
