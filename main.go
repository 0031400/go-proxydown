package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	port := flag.String("port", "18080", "port of the server")
	host := flag.String("host", "127.0.0.1", "the host of the server")
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		resp, err := http.Get(url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}
		if resp.StatusCode == http.StatusNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
			w.Header().Set("Content-Type", "text/plain")
		}
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		log.Println(url, r.RemoteAddr)
	})
	ListenAddr := fmt.Sprintf("%s:%s", *host, *port)
	fmt.Printf("The server is listen on %s\n", ListenAddr)
	log.Fatal(http.ListenAndServe(ListenAddr, nil))
}
