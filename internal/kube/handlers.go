package kube

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/mboufous/k-island/pkg/render"
	"github.com/mboufous/k-island/utils/log"
)

func (api *KubeAPI) V1() {
	api.Router.Route("/v1", func(r chi.Router) {
		r.Mount("/kube", api.NewKubeRouter())
	})
	chi.Walk(api.Router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		log.Infof("[%s] %s", method, route)
		return nil
	})
}

func (api *KubeAPI) NewKubeRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, r.RequestURI+"/namespaces", http.StatusMovedPermanently)
	})

	r.Get("/namespaces", api.getNamespaces())

	r.Route("/{namespace}", func(r chi.Router) {
		r.Get("/pods", api.getPods())
		r.Get("/ingresses", api.getIngresses())
	})
	return r
}

func (api *KubeAPI) getNamespaces() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pods, err := api.KubeService.GetNamespaces()
		if err != nil {
			render.Err(w, http.StatusBadRequest, render.WithError(err))
			return
		}
		render.JSON(w, http.StatusOK, pods)
	}
}

func (api *KubeAPI) getPods() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		namespace := chi.URLParam(r, "namespace")
		// TODO: Validate namespace
		pods, err := api.KubeService.GetPods(namespace)
		if err != nil {
			switch err {
			case ErrNoPodFound:
				// render.Err(w, http.StatusNoContent, render.WithError(err))
				render.Err(w, http.StatusBadRequest, render.WithError(err))
			default:
				render.Err(w, http.StatusBadRequest, render.WithError(err))
			}
			return
		}
		render.JSON(w, http.StatusOK, pods)
	}
}

func (api *KubeAPI) getIngresses() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		namespace := chi.URLParam(r, "namespace")
		// TODO: Validate namespace
		ingresses, err := api.KubeService.GetIngresses(namespace)
		if err != nil {
			switch err {
			case ErrNoIngressFound:
				// render.Err(w, http.StatusNoContent, render.WithError(err))
				render.Err(w, http.StatusBadRequest, render.WithError(err))
			default:
				render.Err(w, http.StatusBadRequest, render.WithError(err))
			}
			return
		}
		render.JSON(w, http.StatusOK, ingresses)
	}
}
