package msgbox

import (
	"github.com/karlmoad/DemoJWTService/authentication"
	"github.com/karlmoad/DemoJWTService/drawgames"
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type secretMessage struct {
	Name		string `json:"name"`
	Message	interface{} `json:"message"`
}



func GetSecretMessage(w http.ResponseWriter, r *http.Request) {
	token := context.Get(r, "TOKEN").(authentication.Token)
	params := mux.Vars(r)

	count := 1

	if val, ok := params["count"]; ok {
		if num, err := strconv.Atoi(val); err == nil {
			count = num
		}
	}

	game, err := drawgames.NewGame("POWERBALL")
	if err != nil {
		log.Fatal(err)
		w.Write([]byte("Failed to initialize " + err.Error()))
		w.WriteHeader(500)
		return
	}

	card, err := game.Draw(count)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		return
	}


	msg := secretMessage{Name:token.GivenName,Message:card}
	json.NewEncoder(w).Encode(msg)
	return
}
