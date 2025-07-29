//go:generate goversioninfo -icon=icon.ico -manifest=goversioninfo.exe.manifest
package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	go server()
	SysTray()
}

func server() {
	var data map[string]string = make(map[string]string)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		setHeaders(w)
		switch r.Method {
		case "GET":
			d, e := json.Marshal(data)
			if e != nil {
				http.Error(w, e.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte(d))
		case "POST":
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
	})
	http.ListenAndServe(":8080", nil)
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
