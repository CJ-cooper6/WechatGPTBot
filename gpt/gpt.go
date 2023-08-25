package gpt

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"wechatbot/config"
)

// ResponseData chatGPT响应体
type ResponseData struct {
	Text            string `json:"text"`
	ConversationID  string `json:"conversation_id"`
	ParentMessageID string `json:"parent_message_id"`
}

func Completions(msg string) (string, error) {
	config := config.LoadConfig()
	requestData := url.Values{}
	requestData.Set("access_token", config.AccessToken)
	requestData.Set("prompt", msg)

	// 发送 POST 请求
	response, err := http.Post(config.URL, "application/x-www-form-urlencoded", strings.NewReader(requestData.Encode()))
	if err != nil {
		return "", err
	}
	if response.StatusCode == 500 {
		return "", errors.New("invalid [access_token] or [prompt]")
	}
	defer response.Body.Close()

	var responseData ResponseData
	json.NewDecoder(response.Body).Decode(&responseData)
	return responseData.Text, nil
}
