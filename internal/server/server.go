package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"clusterapi/internal/k8s"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/pods", ListPodsHandler).Methods("GET")

	fmt.Println("Server starting on :8088...")
	http.ListenAndServe(":8088", r)
}

func ListPodsHandler(w http.ResponseWriter, r *http.Request) {
	clientset, err := k8s.GetClientset()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating kubernetes client: %v", err), http.StatusInternalServerError)
		return
	}

	podNames, err := k8s.ListPods(clientset, "default")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error listing pods: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(podNames)

}
