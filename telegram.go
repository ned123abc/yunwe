package telegram

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// SendMessage sends a message through the Telegram Bot API
func SendMessage(text string) error {
	chatID := "-4112414208"
	url := "https://api.telegram.org/bot6559646803:AAFTKQmzFnx1dzbDT9z3mkqU_RzF2lBw_Fs/sendMessage"

	// 构建消息体
	message := map[string]string{
		"chat_id": chatID,
		"text":    text,
	}
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		return err
	}

	// 发送请求
	_, err = http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
	return err
}
