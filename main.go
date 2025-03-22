package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	host string
	port uint64
}

func main() {
	banner()
	port := flag.Uint64("p", 8080, "port of the server")
	host := flag.String("host", "0.0.0.0", "the host of the server")
	flag.Parse()
	config := &Config{host: *host, port: *port}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		resp, err := http.Get(url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "")
			myLogger(r, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		if strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "")
			myLogger(r, http.StatusNotFound)
			return
		}
		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "")
			myLogger(r, http.StatusInternalServerError)
			return
		}
		myLogger(r, http.StatusOK)
	})
	fmt.Printf("the server is listen on %s:%d\n", config.host, config.port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.host, config.port), nil)
	if err != nil {
		log.Fatal("error", err)
	}
}
func banner() {
	myBanner := `----------------------
go-proxydown made by 0031400
----------------------
-h help
----------------------`
	fmt.Println(myBanner)
}
func myLogger(r *http.Request, statusCode int) {
	fmt.Printf("%s %s %d %s %s %s\n", r.Method, r.URL.Path, statusCode, time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.UserAgent())
}
