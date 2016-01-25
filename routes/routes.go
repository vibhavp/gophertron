package routes

import (
	"html/template"
	"net/http"

	_ "github.com/rakyll/gom/http"
	"github.com/vibhavp/gophertron/controllers"
	"github.com/vibhavp/gophertron/models"
)

func Main(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./views/active.html"))
	tmpl.Execute(w, models.GetGames())
}

func InitRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", Main)
	mux.HandleFunc("/create", controllers.Create)
	mux.HandleFunc("/join", controllers.Join)
	mux.HandleFunc("/game", controllers.Game)
	mux.HandleFunc("/websocket", controllers.WebSocket)
	mux.HandleFunc("/game.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./views/game.js")
	})
}
