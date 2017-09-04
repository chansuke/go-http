package main

import (
	"fmt"
	"log"
	"net/http"
	"net/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprit(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dmup))
	fmt.Fprinf(w, "<html><body>hello</body></head>\n")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("start http listening :8443")
	err := http.ListenAndServerTLS(":8443", "server.crt", "server.key", nil)
	log.Println(err)
}
