package controllers

import (
	"net/http"

	"github.com/SpoonBuoy/waba/models"
	"github.com/SpoonBuoy/waba/service"
	"github.com/gin-gonic/gin"
)

const systemAccessToken = "EAAMRwudSuNsBOwUijKcd1HmrCsL9twIZA5pFE8ASRB5S3yIdMiWVRRvOw2l5ZCumitduU4foUofA6CZBbEe6LcCYAiw4nsYZB1Ub6KsWihDBdGzg8m4E3UOBXW4KKC9bokSxV2inh9PvQMhIKxocBmY5Ja8uIucg5NIi3ZAxFst5xl2G24jFFA5pU082S19rp"
const token = "EAAnrp3AqNGwBOz2ltZCOaeWBQ5OdDm04TH1mnIibbH8aZC0AoZBVB0g1BuLNwq96DsCfe5Ub9KWvKPGUkZAmvO8REmPVUAbKfcyNMTuhN4Jq9sLfwd29SuZBRH5ZBhMa8KfDnZAMRgxbg4wCYuJOYjrB435nL64YHIt3KuGc1LhqlZAu2AN3HXjNRPrCys581QgbA6YZCLZA4cko8tXA8giZAAZD"

type ChatController struct {
	service *service.WabaService
}

func NewChatController(service *service.WabaService) *ChatController {
	return &ChatController{
		service: service,
	}
}

// verify api
func (wbc *ChatController) Verify(c *gin.Context) {
	// "/verify?hub.mode=subscribe&hub.challenge=1291346053&hub.verify_token=bot-test"
	mode := c.Query("hub.mode")
	challenge := c.Query("hub.challenge")
	verifyToken := c.Query("hub.verify_token")

	if mode == "subscribe" && verifyToken == "bot-test" {
		c.String(http.StatusOK, challenge)
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// webhook listening new message events
func (wbc *ChatController) Listen(c *gin.Context) {
	var body models.WAMessagePayload
	if err := c.BindJSON(&body); err != nil {
		c.Error(err)
		return
	}

	if len(body.Entry) > 0 && len(body.Entry[0].Changes) > 0 && len(body.Entry[0].Changes[0].Value.Messages) > 0 && body.Entry[0].Changes[0].Field == "messages" {
		if err := wbc.service.Listen(c, &body); err != nil {
			c.Error(err)
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Payload type not supported",
		})
		return
	}
}
