package health

import (
	"encoding/json"
	"net/http"
)

type statusServer struct {
	statusInfo map[string]string
}

func (ss *statusServer) handleStatus(w http.ResponseWriter, r *http.Request) {

	payload, err := json.Marshal(ss.statusInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func StartHttpListener(listenAddr string, statusInfo map[string]string) {

	// static value at the moment health checks "comming soon"
	statusInfo["status"] = "OK"

	ss := &statusServer{
		statusInfo: statusInfo,
	}

	http.HandleFunc("/status", ss.handleStatus)

	go http.ListenAndServe(listenAddr, nil)
}
