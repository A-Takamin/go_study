package main

import (
	"flag"
	"go_study/trace"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// templは1つのテンプレート
type templateHandler struct {
	once     sync.Once // テンプレートを一回だけコンパイル
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

/**
main関数が呼ばれたらサーバがスタートするだけ。
"/"にアクセスすると、chat.htmlを返す。
*/
func main() {
	// flagはコマンドラインの引数に関するパッケージ
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	flag.Parse() //フラグを解釈する。デフォルトは8080。
	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)
	// チャットルームを開始します。バックグラウンドで。
	go r.run()
	// webサーバの開始
	log.Println("webサーバを起動します。ポート：", *addr)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
