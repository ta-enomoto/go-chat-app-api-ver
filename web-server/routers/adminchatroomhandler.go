package routers

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goserver/query"
	"goserver/sessions"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func AdminChatroomHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/admin/adminsessionexpired.html"))
			t.ExecuteTemplate(w, "adminsessionexpired.html", nil)
			return
		}

		userCookie, _ := r.Cookie(session.Manager.CookieName)
		userSid, _ := url.QueryUnescape(userCookie.Value)
		userSessionVar := session.Manager.SessionStore[userSid].SessionValue["userId"]

		if userSessionVar != "admin" {
			fmt.Fprintf(w, "管理者以外はアクセスできません。")
			return
		}

		t := template.Must(template.ParseFiles("./templates/admin/chatrooms/adminchatroom.html"))
		t.ExecuteTemplate(w, "adminchatroom.html", nil)
		return

	case "POST":
		if ok := session.Manager.SessionIdCheck(w, r); !ok {
			t := template.Must(template.ParseFiles("./templates/admin/adminsessionexpired.html"))
			t.ExecuteTemplate(w, "adminsessionexpired.html", nil)
			return
		}

		userCookie, _ := r.Cookie(session.Manager.CookieName)
		userSid, _ := url.QueryUnescape(userCookie.Value)
		userSessionVar := session.Manager.SessionStore[userSid].SessionValue["userId"]

		if userSessionVar != "admin" {
			fmt.Fprintf(w, "管理者以外はアクセスできません。")
			return
		}

		if r.FormValue("delete-room") == "このルームを削除する" {
			roomUrl := r.URL.Path
			_roomId := strings.TrimPrefix(roomUrl, "/admin/chatrooms/chatroom")
			roomId, _ := strconv.Atoi(_roomId)

			dbChtrm, err := sql.Open("mysql", query.ConStrChtrm)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer dbChtrm.Close()

			query.DeleteChatroomById(roomId, dbChtrm)

			t := template.Must(template.ParseFiles("./templates/mypage/chatroomdeleted.html"))
			t.ExecuteTemplate(w, "chatroomdeleted.html", nil)
		}
	}
}