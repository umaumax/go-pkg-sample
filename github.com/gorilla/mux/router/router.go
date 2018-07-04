package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/goji/httpauth"
	"github.com/gorilla/mux"
)

var (
	port          string
	basicAuthUser string
	basicAuthPass string
)

func init() {
	flag.StringVar(&port, "p", ":5555", "port")

	//	Basic認証
	flag.StringVar(&basicAuthUser, "user", "user", "basic auth user")
	flag.StringVar(&basicAuthPass, "pass", "pass", "basic auth pass")
}

func main() {
	flag.Parse()

	router := mux.NewRouter()
	//	NOTE
	//	as run static file server http://example.com/assets/*** -> root/assets/***
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("root/assets/"))))

	//	NOTE
	//	先にPathPrefixで指定した上限へのアクセスが優先
	//	同一のPathPrefixを指定すると最初のみが適用されるため注意
	router.PathPrefix("/api").Handler(apiRouting())
	//	as run static file server http://example.com/*** -> root/public/***
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("root/public/"))))

	//	NOTE
	//	全てのアクセスに対してかませたい処理
	http.Handle("/", httpauth.SimpleBasicAuth(basicAuthUser, basicAuthPass)(router))

	log.Println("listeing at", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln("http server err:", err)
	}
}
