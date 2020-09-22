package providers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"webhook/config"
)

type GroupOpt struct {
	GroupAdmin string
	GroupName  string
}

func NewGroupOpt(groupAdmin string, groupName string) *GroupOpt {
	return &GroupOpt{
		GroupAdmin: groupAdmin,
		GroupName:  groupName,
	}
}

func (s *GroupOpt) Send(msg string) bool {
	return whatsapp(msg, s.GroupAdmin, s.GroupName)
}

func whatsapp(msg string, groupAdmin string, groupName string) bool {
	instanceID := config.InstanceID
	clientID := config.ClientID
	clientSecret := config.ClientSecret

	message := fmt.Sprintf(`{ "group_admin":"%s", "group_name": "%s", "message": "%s" }`, groupAdmin, groupName, msg)
	log.Printf("Post data info: %v\n", message)
	client := &http.Client{}
	data := strings.NewReader(message)
	url := fmt.Sprintf("http://api.whatsmate.net/v3/whatsapp/group/text/message/%s", instanceID)
	log.Println(url)
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("X-WM-CLIENT-ID", clientID)
	req.Header.Set("X-WM-CLIENT-SECRET", clientSecret)
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
	return true
}
