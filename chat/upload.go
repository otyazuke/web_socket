package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("userid")
	file, header, err := r.FormFile("avatarFile")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	filename := filepath.Join("avatars", userId+filepath.Ext(header.Filename))
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, "成功")
}
