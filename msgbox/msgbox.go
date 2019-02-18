package msgbox

import (
	"net/http"
)

func GetInboxMessage(w http.ResponseWriter, r *http.Request) {
/*	user := context.Get(r, "USER").(string)
	params := mux.Vars(r)

	api, err := workspaces.NewApiInstance(user)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}

	ws, err := api.GetWorkspaceContents(params["id"], params["sha"])
	if err == nil {
		json.NewEncoder(w).Encode(ws)
	} else {
		w.WriteHeader(404)
	}*/
}


func GetInboxMessages(w http.ResponseWriter, r *http.Request) {



}

func SendMessage(w http.ResponseWriter, r *http.Request) {



}
