package providers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"webhook/config"

	"go.uber.org/zap"
)

type whatsappOpt struct {
	groupAdmin string
	groupName  string
}

func NewWhatsappOpt(groupAdmin string, groupName string) *whatsappOpt {
	return &whatsappOpt{
		groupAdmin: groupAdmin,
		groupName:  groupName,
	}
}

func (w *whatsappOpt) send(msg string) {
	message := fmt.Sprintf(`{ "group_admin":"%s", "group_name": "%s", "message": "%s" }`,
		w.groupAdmin, w.groupName, msg)
	zap.L().Info("whatsapp post alert.", zap.Any("data", message))
	client := &http.Client{}
	data := strings.NewReader(message)
	url := fmt.Sprintf("http://api.whatsmate.net/v3/whatsapp/group/text/message/%s", config.InstanceID)
	zap.L().Info("Request url:", zap.Any("url", url))
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		zap.L().Error("NewRequest whatsapp api fail", zap.Error(err))
	}
	req.Header.Set("X-WM-CLIENT-ID", config.ClientID)
	req.Header.Set("X-WM-CLIENT-SECRET", config.ClientSecret)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error("Client.Do fail", zap.Error(err))
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error("ioutil.ReadAll fail", zap.Error(err))
	}
	zap.L().Info("Request whatsapp response", zap.Any("response", bodyText))
}
