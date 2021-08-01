package routers

import (
	"html/template"
	"net/http"
	"web-server/sessions"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//当該セッションを削除する
		session.Manager.DeleteSessionFromStore(w, r)

		t := template.Must(template.ParseFiles("./templates/logout.html"))
		t.ExecuteTemplate(w, "logout.html", nil)
	}
}
