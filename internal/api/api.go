package api

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/Coflnet/homepage/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/hostrouter"
)

type WebServer struct {
	config     *usecase.Config
	translator *usecase.Translator
}

func NewWebServer(config *usecase.Config, translator *usecase.Translator) *WebServer {
	return &WebServer{
		config:     config,
		translator: translator,
	}
}

func (s *WebServer) StartServer() error {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	hr := hostrouter.New()

	hr.Map("consulting.coflnet.com", s.StartConsultingPage())
	hr.Map("*", s.StartHomepage())

	r.Mount("/", hr)

	return http.ListenAndServe(":3000", r)
}

func (s *WebServer) StartHomepage() chi.Router {

	r := chi.NewRouter()
	r.Get("/", s.handleHome)
	r.Get("/impressum", s.handleImprint)
	r.Post("/contact", s.handleContactFormPost)

	fs := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fs))

	return r
}

func (s *WebServer) StartConsultingPage() chi.Router {
	r := chi.NewRouter()
	r.Get("/", s.handleHome)

	return r
}

func (s *WebServer) handleHome(w http.ResponseWriter, r *http.Request) {
	lang := r.Header.Get("Accept-Language")

	tmpl := template.Must(template.ParseGlob("./internal/views/*.html"))
	projects := s.config.ListProjects()
	websiteData := s.translator.RetrieveWebsiteDataWithProjects(lang, projects)

	err := tmpl.ExecuteTemplate(w, "homepage.html", websiteData)
	if err != nil {
		slog.Error("Error while executing template: ", "err", err)
		return
	}
}

func (s *WebServer) handleImprint(w http.ResponseWriter, r *http.Request) {
	lang := r.Header.Get("Accept-Language")

	tmpl := template.Must(template.ParseGlob("./internal/views/*.html"))
	websiteData := s.translator.RetrieveWebsiteData(lang)

	err := tmpl.ExecuteTemplate(w, "impressum.html", websiteData)
	if err != nil {
		slog.Error("Error while executing template: ", "err", err)
		return
	}
}

func (s *WebServer) handleContactFormPost(w http.ResponseWriter, r *http.Request) {
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	message := r.FormValue("message")

	tmpl := template.Must(template.ParseGlob("./internal/views/*.html"))

	err := usecase.SendContactMessage(firstname, lastname, email, message)
	if err != nil {
		slog.Error("Error while sending contact message: ", "err", err)

		err = tmpl.ExecuteTemplate(w, "contact-error.html", nil)
		if err != nil {
			slog.Error("Error while executing template: ", "err", err)
			return
		}
		return
	}

	err = tmpl.ExecuteTemplate(w, "contact-success.html", nil)
	if err != nil {
		slog.Error("Error while executing template: ", "err", err)
	}
}
