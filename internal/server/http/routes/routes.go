package routes

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/vsPEach/Framework/internal/entity"
	"github.com/vsPEach/Framework/internal/repository/sql"
	"net/http"
)

type Storage interface {
	Create(ctx context.Context, model sql.Model)
	Update(ctx context.Context, model sql.Model)
	Delete(ctx context.Context, model sql.Model)
	FindAll(ctx context.Context, model sql.Model)
}

type HTTPHandler struct {
	s Storage
}

func (h *HTTPHandler) Routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", h.welcomeHandler)

	r.HandleFunc("/login", h.signIn).Methods(http.MethodPost)
	r.HandleFunc("/register", h.signUp).Methods(http.MethodPost)

	r.HandleFunc("/articles/article", h.createArticle).Methods(http.MethodPost)
	r.HandleFunc("/articles/{id}", h.getArticle).Methods(http.MethodGet)
	r.HandleFunc("/articles", h.getArticles).Methods(http.MethodGet)
	r.HandleFunc("/articles/{id}", h.updateArticle).Methods(http.MethodPost)
	r.HandleFunc("/articles/{id}", h.deleteArticle).Methods(http.MethodDelete)

	r.HandleFunc("/comments/comment", h.createComment).Methods(http.MethodPost)
	r.HandleFunc("/comments/{id}", h.getComment).Methods(http.MethodGet)
	r.HandleFunc("/comments", h.getComments).Methods(http.MethodGet)
	r.HandleFunc("/comments/comment/{id}", h.updateComment).Methods(http.MethodPost)
	r.HandleFunc("/comments/comment/{id}", h.deleteComment).Methods(http.MethodDelete)

	return r
}

func (h *HTTPHandler) signIn(w http.ResponseWriter, r *http.Request) {

}
func (h *HTTPHandler) signUp(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHandler) createArticle(w http.ResponseWriter, r *http.Request) {
}

func (h *HTTPHandler) getArticle(w http.ResponseWriter, r *http.Request) {
}

func (h *HTTPHandler) updateArticle(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHandler) getArticles(w http.ResponseWriter, r *http.Request) {
	h.s.FindAll(context.Background(), entity.Comment{})
}

func (h *HTTPHandler) deleteArticle(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHandler) createComment(w http.ResponseWriter, r *http.Request) {
}

func (h *HTTPHandler) updateComment(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHandler) deleteComment(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHandler) getComment(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHandler) getComments(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHandler) welcomeHandler(w http.ResponseWriter, r *http.Request) {

}
