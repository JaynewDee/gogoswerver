package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gogoswerver"
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

const threadsListHTML = `
<h1 style="text-align:center;">Threads</h1>
<dl>
{{range .Threads}}
	<dt>
		<strong>{{.Title}}<strong>
	</dt>
	<dd>{{.Description}}<dd>
	<dd>
		<form action="/threads/{{.ID}}/delete" method="POST">
			<button type="submit">DELETE</button>
		</form>
	</dd>
{{end}}
</dl>
<a href="/threads/new">Create thread</a>
`

func (h *Handler) ThreadsList() http.HandlerFunc {
	// Space above returned handler func can be used for handler initialization as it is only run once!
	type data struct {
		Threads []gogoswerver.Thread
	}
	temp := template.Must(template.New("List Threads").Parse(threadsListHTML))
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

const threadCreateHTML = `
<style>
	div {
		text-align: center;
	}
	textarea {
		width: 66vw
	}
</style>

<div>
<h2>Create New Thread</h2>
<form action="/threads" method="POST">
	<table>
		<tr>
			<td>Title</td>
			<td><input type="text" name="title" /></td>
		</tr>
		<tr>
			<td>Description</td>
			<td><textarea name="description"></textarea></td>
		</tr>
	</table>
	<button type="submit">Create!</button>
</form>
</div>
`

func (h *Handler) ThreadsCreate() http.HandlerFunc {
	temp := template.Must(template.New("Create Thread").Parse(threadCreateHTML))

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
