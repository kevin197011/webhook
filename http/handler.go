package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"webhook/config"
	"webhook/providers"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Write([]byte("ok!"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Write([]byte("Use a POST request to send a message!"))
}

func whatsappAlertsHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Printf("Npt recv request...")
	switch r.Method {
	case http.MethodGet:
		getHandler(w, r)
	case http.MethodPost:
		whatsappPostHandler(w, r)
	default:
		http.Error(w, "Unsupported HTTP method!", 400)
	}
}

func whatsappPostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var emoji string
	dec := json.NewDecoder(r.Body)

	var m HookMessage
	if err := dec.Decode(&m); err != nil {
		log.Printf("error decoding message: %v", err)
		http.Error(w, "invalid request body", 400)
		return
	}
	log.Printf("bind alert json data: %v\n", m)

	if m.Status == "firing" {
		emoji = "❌"
	} else {
		emoji = "✅"
	}

	cf := config.NewConfig()

	for _, alert := range m.Alerts {
		var msg string

		msg = fmt.Sprintf("*状态:* %s %s\\n", m.Status, emoji)
		msg += "*标签:*\\n"

		keys := make([]string, 0, len(alert.Labels))
		for k := range alert.Labels {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		for _, key := range keys {
			msg += fmt.Sprintf("  %s = %s\\n", key, alert.Labels[key])
		}

		msg += fmt.Sprintf("*注解:*\\n  %s\\n", alert.Annotations["message"])
		msg += fmt.Sprintf("*时间:* %s", alert.StartsAt)

		log.Printf("send alert data: %v\n", msg)
		providers.NewGroupOpt("123456", cf.GroupName).Send(msg)
	}

	w.Write([]byte(`{"status": 200, "msg": "ok"}`))
}
