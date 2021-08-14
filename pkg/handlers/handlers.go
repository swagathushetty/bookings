package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/swagathushetty/bookings/pkg/config"
	"github.com/swagathushetty/bookings/pkg/models"
	"github.com/swagathushetty/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

//creates new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	//storing the users IP address and stroing it in session
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
	// fmt.Fprintf(w, "this is the home page")
	//or
	// w.Write([]byte("this is the home page"))
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// sum := AddValues(2, 3)
	// fmt.Fprintf(w, fmt.Sprintf("this is the about page and 2+2 is %d", sum))
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again"

	//accessing the users IP address that was added when user was in "/"
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func AddValues(x, y int) int {
	return x + y
}

func (m *Repository) Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)

	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {

	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}

	result := x / y

	return result, nil

}
