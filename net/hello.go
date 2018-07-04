package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
)

type PostServer struct {
}

//	構造体にSeveHTTPを登録することでサーバーとなる
func (v PostServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("[Server Log]")
	log.Println("[URL]", r.URL)
	log.Println("[UserAgent]", r.UserAgent())

	//	fmt.Fprintln(w, r.PostFormValue("key1"))
	//	フォームの値を取得するにはPaseFormまたはPostFormValueなどを事前に行う必要がある
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	}
	//	POST Value
	fmt.Fprintln(w, r.PostForm)
	//	POST & GET Value
	fmt.Fprintln(w, r.Form)
}

func StartHttpServer() {
	http.Handle("/", &PostServer{})
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	urlStr := "http://localhost:9999/"
	go StartHttpServer()
	PostSample(urlStr)
	GetSample(urlStr)
	PrintHostToIp("localhost")
	TCPIPSample()
}

func GetSample(urlStr string) {
	fmt.Println("[GET]")
	res, err := http.Get(urlStr)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("[status]")
	fmt.Println(res.Status)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("[header]")
	fmt.Println(res.Header)
	fmt.Println("[body]")
	fmt.Println(string(body))
	fmt.Println()
}

func PostSample(urlStr string) {
	fmt.Println("[POST]")
	res, err := http.PostForm(urlStr,
		url.Values{
			"key1": {"value1", "value2"},
			"key2": {"value"},
		})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("[status]")
	fmt.Println(res.Status)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("[header]")
	fmt.Println(res.Header)
	fmt.Println("[body]")
	fmt.Println(string(body))
	fmt.Println()
}

//	ホスト名 -> IPAddress
func PrintHostToIp(host string) {
	fmt.Println("[host to ip]")
	addrs, _ := net.LookupHost(host)
	fmt.Println(host, "->", addrs)
	fmt.Println()
}

func TCPIPSample() {
	fmt.Println("[TCP/IP sample]")
	conn, err := net.Dial("tcp", "golang.jp:80")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	//	HTTPリクエストの送信
	fmt.Fprintf(conn, "GET /hello.html HTTP/1.1\r\n")
	fmt.Fprintf(conn, "HOST: golang.jp\r\n")
	fmt.Fprintf(conn, "\r\n")

	//	HTTPレスポンスの受信
	response, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	//	バイトスライス出力
	fmt.Println(response)
	//	文字列出力
	fmt.Println(string(response))
	fmt.Println()
}
