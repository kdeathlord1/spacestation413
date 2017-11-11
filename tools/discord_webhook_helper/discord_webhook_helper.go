// Space Station 13 Discord webhook integration helper script
// by Difarem

package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Printf("Usage: %s [webhook ID] [webhook token] [listen address]\n", os.Args[0])
		os.Exit(1)
	}
	id := os.Args[1]
	token := os.Args[2]
	addr := os.Args[3]

	dg, err := discordgo.New()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/broadcast", func(w http.ResponseWriter, req *http.Request) {
		err := req.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if msg, ok := req.Form["msg"]; ok {
			var params discordgo.WebhookParams
			params.Content = msg[0]

			// send the message via the webhook
			err := dg.WebhookExecute(id, token, false, &params)
			if err != nil {
				// something went wrong when sending the message
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	})

	// listen for incoming connections from the BYOND server
	log.Fatal(http.ListenAndServe(addr, mux))
}
