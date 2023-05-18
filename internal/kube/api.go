package kube

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type KubeAPI struct {
	KubeService KubeService
	Router      chi.Router
}

func NewKubeAPI(kubeService KubeService) *KubeAPI {
	r := chi.NewRouter()
	r.Use(
		middleware.RequestID,
		middleware.Logger,
		middleware.Recoverer,
	)

	return &KubeAPI{
		KubeService: kubeService,
		Router:      r,
	}
}

func (api *KubeAPI) APIHandler() http.Handler {
	return api.Router
}
