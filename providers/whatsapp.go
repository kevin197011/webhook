package providers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"webhook/config"
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
	log.Printf("Post data info: %v\n", message)
	client := &http.Client{}
	data := strings.NewReader(message)
	url := fmt.Sprintf("http://api.whatsmate.net/v3/whatsapp/group/text/message/%s", config.InstanceID)
	log.Println(url)
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("X-WM-CLIENT-ID", config.ClientID)
	req.Header.Set("X-WM-CLIENT-SECRET", config.ClientSecret)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%s\n", bodyText)
}
