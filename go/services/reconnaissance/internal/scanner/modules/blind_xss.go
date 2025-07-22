package modules

import (
	"fmt"
	"net/http"
)

func StartBlindXSSServer() {
	http.HandleFunc("/xss", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[!] Blind XSS Callback! Query:", r.URL.RawQuery)
	})
	fmt.Println("Listening for Blind XSS callbacks...")
	go http.ListenAndServe(":8081", nil)
}
