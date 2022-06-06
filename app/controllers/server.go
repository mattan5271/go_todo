package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"
	"todo_app/app/models"
	"todo_app/config"
)

// HTML生成
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...)) // Mustでテンプレートを予めキャッシュしている
	templates.ExecuteTemplate(writer, "layout", data)
}

// ログインチェック
func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{UUID: cookie.Value}
		if ok, _ := sess.CheckSession(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// サーバー起動
func StartMainServer() error {
	// 静的ファイル読み込み
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// ルート一覧
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/todos", index)
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)

	// 第2引数にnilを渡すことで、デフォルト設定(ページが存在しない場合に404を返す)を使用する
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
