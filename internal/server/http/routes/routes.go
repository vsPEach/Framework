package routes

import (
	"context"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/vsPEach/Framework/internal/entity"
	"github.com/vsPEach/Framework/internal/repository/sql"
	"github.com/vsPEach/Framework/pkg/utils"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Storage interface {
	Create(ctx context.Context, model sql.Model) error
	Update(ctx context.Context, model sql.Model) error
	Delete(ctx context.Context, model sql.Model) error
	Find(ctx context.Context, model sql.Model) (sql.Model, error)
	FindAll(ctx context.Context, model sql.Model) ([]sql.Model, error)
}

type ArticleView struct {
	ArticleID   uuid.UUID
	Title       string
	ArticleText string
	CreatedAt   time.Time
	Comments    any
}

type ViewData struct {
	Title string
	Data  any
}

type HTTPHandler struct {
	s Storage
}

func NewHTTPHandler(storage Storage) *HTTPHandler {
	return &HTTPHandler{s: storage}
}

func (h *HTTPHandler) Routes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", h.welcomeHandler)

	r.HandleFunc("/login", h.signIn).Methods(http.MethodPost)
	r.HandleFunc("/register", h.signUp).Methods(http.MethodPost)

	r.HandleFunc("/articles/article", h.createArticle).Methods(http.MethodPost, http.MethodGet)
	r.HandleFunc("/articles/{id}", h.getArticle).Methods(http.MethodGet)
	r.HandleFunc("/articles/edit/{id}", h.updateArticle).Methods(http.MethodPost, http.MethodGet)
	r.HandleFunc("/articles/delete/{id}", h.deleteArticle).Methods(http.MethodDelete, http.MethodGet)

	r.HandleFunc("/comments/comment", h.createComment).Methods(http.MethodPost)
	r.HandleFunc("/comments/{id}", h.getComment).Methods(http.MethodGet)
	r.HandleFunc("/comments/edit/{id}", h.updateComment).Methods(http.MethodPost)
	r.HandleFunc("/comments/comment/{id}", h.deleteComment).Methods(http.MethodDelete, http.MethodGet)

	return r
}

func (h *HTTPHandler) signIn(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHandler) signUp(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHandler) createArticle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		executeTemplate(w, "article/create.html", nil)
	case http.MethodPost:
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		request := utils.ArticleToSlice(string(bytes))
		err = h.s.Create(context.Background(), entity.Article{
			ID:        uuid.New(),
			AuthorID:  uuid.UUID{},
			Title:     request[0],
			Text:      request[1],
			CreatedAt: time.Now(),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (h *HTTPHandler) getArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	id, err := utils.StringToUUID(strID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	find, err := h.s.Find(context.Background(), entity.Article{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	obj := find.(entity.Article)
	av := ArticleView{
		ArticleID:   obj.ID,
		Title:       obj.Title,
		ArticleText: obj.Text,
		CreatedAt:   obj.CreatedAt,
		Comments:    h.getComments(),
	}
	executeTemplate(w, "article/show.html", av)
}

func (h *HTTPHandler) updateArticle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		executeTemplate(w, "article/edit.html", strings.TrimPrefix(r.URL.String(), "/"))
	case http.MethodPost:
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		vars := mux.Vars(r)
		strID := vars["id"]
		id, err := utils.StringToUUID(strID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		request := utils.ArticleToSlice(string(bytes))
		err = h.s.Update(context.Background(), entity.Article{
			ID:        id,
			AuthorID:  uuid.UUID{},
			Title:     request[0],
			Text:      request[1],
			CreatedAt: time.Now(),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (h *HTTPHandler) deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	id, err := utils.StringToUUID(strID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = h.s.Delete(context.Background(), entity.Article{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (h *HTTPHandler) createComment(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	request := utils.CommentToString(string(bytes))
	err = h.s.Create(context.Background(), entity.Comment{
		ID:        uuid.New(),
		AuthorID:  uuid.UUID{},
		Text:      request,
		CreatedAt: time.Now(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := r.Header["Referer"]
	paths := strings.Split(url[0], "/")
	http.Redirect(w, r, "/articles/"+paths[len(paths)-1], http.StatusSeeOther)
}

func (h *HTTPHandler) updateComment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		executeTemplate(w, "comment/edit.html", strings.TrimPrefix(r.URL.String(), "/"))
	case http.MethodPost:
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		vars := mux.Vars(r)
		strID := vars["id"]
		id, err := utils.StringToUUID(strID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		request := utils.CommentToString(string(bytes))
		err = h.s.Update(context.Background(), entity.Article{
			ID:        id,
			AuthorID:  uuid.UUID{},
			Text:      request,
			CreatedAt: time.Now(),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (h *HTTPHandler) deleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	id, err := utils.StringToUUID(strID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = h.s.Delete(context.Background(), entity.Comment{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *HTTPHandler) getComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]
	id, err := utils.StringToUUID(strID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	find, err := h.s.Find(context.Background(), entity.Comment{ID: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	executeTemplate(w, "comment/show.html", find)
}

func (h *HTTPHandler) getComments() []entity.Comment {
	var result []entity.Comment
	all, err := h.s.FindAll(context.Background(), entity.Comment{})
	if err != nil {
		return nil
	}
	for _, model := range all {
		result = append(result, model.(entity.Comment))
	}
	return result
}

func (h *HTTPHandler) welcomeHandler(w http.ResponseWriter, r *http.Request) {
	all, err := h.s.FindAll(context.Background(), entity.Article{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view := ViewData{"Articles", all}
	executeTemplate(w, "index.html", view)
}

func executeTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	filename := "./internal/server/http/templates/"
	tmpl, err := template.ParseFiles(filename + tmplName)
	if err != nil {
		log.Println("TMPL ERROR", err)
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("EXECUTE TMPL ERROR", err)
	}
}
