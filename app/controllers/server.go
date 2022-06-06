package controllers

import (
	"fmt"
	"net/http"
	"text/template"
	"todo_app/config"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...)) // Mustでテンプレートを予めキャッシュしている
	templates.ExecuteTemplate(writer, "layout", data)
}

func StartMainServer() error {
	// 静的ファイル読み込み
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// ルート一覧
	http.HandleFunc("/", top)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)

	// 第2引数にnilを渡すことで、デフォルト設定(ページが存在しない場合に404を返す)を使用する
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
