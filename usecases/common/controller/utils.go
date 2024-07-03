package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	OK   = http.StatusOK
	BAD  = http.StatusBadRequest
	DENY = http.StatusUnauthorized
)

var (
	Success = gin.H{"result": "success"}
)

type KeyVal struct {
	Key string
	Val interface{}
}

func KV(key string, val interface{}) KeyVal {
	return KeyVal{
		Key: key,
		Val: val,
	}
}

func Ok(data ...KeyVal) (int, gin.H) {
	response := gin.H{"result": "success"}
	for _, d := range data {
		response[d.Key] = d.Val
	}
	return OK, response
}
func Bad(err any) (int, gin.H) {
	return BAD, gin.H{"error": err}
}
func Deny(err any) (int, gin.H) {
	return DENY, gin.H{"error": err}
}
