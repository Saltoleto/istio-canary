package main

import "net/http"

func main()  {

	http.HandleFunc("/saltoleto", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("<h1>Hello Saltoleto2<h1>"))
	})
	http.ListenAndServe(":8080", nil)
}

