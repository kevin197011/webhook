package http

import (
	"fmt"
	"log"
	"net/http"
	"webhook/config"
)

func InitRouter() {
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/whatsapp", whatsappAlertsHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.NewConfig().Port), nil))
}
