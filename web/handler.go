package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/JaynewDee/gogoswerver"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
)

type Handler struct {
	*chi.Mux

	store gogoswerver.Store
}

func NewHandler(store gogoswerver.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)

	// Route Grouping        r is "subrouter"
	h.Route("/threads", func(r chi.Router) {
		r.Get("/", h.ThreadsList())
		r.Get("/new", h.ThreadsCreate())
		r.Post("/", h.ThreadsStore())
		r.Post("/{id}/delete", h.ThreadsDelete())
	})

	return h
}

func (h *Handler) ThreadsList() http.HandlerFunc {
	// Space above returned handler func can be used for handler initialization as it is only run once!
	type data struct {
		Threads []gogoswerver.Thread
	}
	temp := template.Must(ThreadsListTemplate())
	//

	return func(w http.ResponseWriter, r *http.Request) {
		ts, err := h.store.Threads()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		temp.Execute(w, data{Threads: ts})
	}
}

func (h *Handler) ThreadsCreate() http.HandlerFunc {
	temp := template.Must(ThreadCreateTemplate())

	fmt.Print("bruh ...")

	return func(w http.ResponseWriter, r *http.Request) {
		temp.Execute(w, nil)
	}
}

func (h *Handler) ThreadsStore() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		description := r.FormValue("description")

		if err := h.store.CreateThread(&gogoswerver.Thread{
			ID:          uuid.New(),
			Title:       title,
			Description: description,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}

func (h *Handler) ThreadsDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idStr)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := h.store.DeleteThread(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/threads", http.StatusFound)
	}
}
