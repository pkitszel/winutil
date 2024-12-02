package main

import (
	"net/http"
)

func serve_static_dir(path, port_str string) error {
	http.Handle("/", http.FileServer(http.Dir(path)))
	return http.ListenAndServe(":"+port_str, nil)
}
