// Copyright (c) 2014 Kelsey Hightower. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.
package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
)

func genisoimageHandler(w http.ResponseWriter, r *http.Request) {
	tmpdir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	configPath := path.Join(tmpdir, "/openstack/latest")
	userData := path.Join(configPath, "user_data")
	err = os.MkdirAll(configPath, 755)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	f, err := os.Create(userData)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer os.RemoveAll(tmpdir)
	io.Copy(f, io.LimitReader(r.Body, 10000))
	f.Close()
	cmd := exec.Command("/usr/bin/genisoimage", "-R", "-V", "-J", "config-2", tmpdir)
	var iso bytes.Buffer
	cmd.Stdout = &iso
	err = cmd.Run()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(iso.Bytes())
}

func main() {
	http.HandleFunc("/genisoimage", genisoimageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
