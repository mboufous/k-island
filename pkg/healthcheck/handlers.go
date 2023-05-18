package healthcheck

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/mboufous/k-island/utils/log"
)

func Setup(r chi.Router) {
	r.Get("/ping", Readiness)
	r.Get("/health", Liveness)
}

func Readiness(w http.ResponseWriter, r *http.Request) {
	_, cancel := context.WithTimeout(r.Context(), time.Second)
	defer cancel()

	status := "ok"
	statusCode := http.StatusOK

	// check k8s status (use ctx)

	responseData := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	if err := jsonResponse(w, statusCode, responseData); err != nil {
		log.Error("raediness error: %s", err)
	}
	log.Infof("readiness [statusCode: %d, status: %s]", statusCode, status)
}

func Liveness(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		host = "unavailable"
	}

	data := struct {
		Status    string `json:"status,omitempty"`
		Host      string `json:"host,omitempty"`
		Pod       string `json:"pod,omitempty"`
		PodIP     string `json:"podIP,omitempty"`
		Node      string `json:"node,omitempty"`
		Namespace string `json:"namespace,omitempty"`
	}{
		Status:    "up",
		Host:      host,
		Pod:       os.Getenv("KUBERNETES_PODNAME"),
		PodIP:     os.Getenv("KUBERNETES_NAMESPACE_POD_IP"),
		Node:      os.Getenv("KUBERNETES_NODENAME"),
		Namespace: os.Getenv("KUBERNETES_NAMESPACE"),
	}

	statusCode := http.StatusOK
	if err := jsonResponse(w, statusCode, data); err != nil {
		log.Error("liveness error: %s", err)
	}
	log.Infof("liveness [statusCode: %d, status: %s]", statusCode, data.Status)
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) error {

	// Convert the response value to JSON.
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Set the content type and headers once we know marshaling has succeeded.
	w.Header().Set("Content-Type", "application/json")

	// Write the status code to the response.
	w.WriteHeader(statusCode)
	// Send the result back to the client.
	if _, err := w.Write(jsonData); err != nil {
		return err
	}

	return nil
}
