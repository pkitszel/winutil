package main

import (
	"io"
	"net/http"
	"os"
)

func wget(url, dest string) (errc error) {
	var file *os.File
	if dest == "-" {
		file = os.Stdout
	} else {
		var err error
		file, err = os.Create(dest)
		if err != nil {
			return err
		}
		defer close(file, &errc)
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer close(resp.Body, &errc)

	_, errc = io.Copy(file, resp.Body)
	return
}

func close(c io.Closer, err *error) {
	errc := c.Close()
	if *err == nil {
		*err = errc
	}
}
