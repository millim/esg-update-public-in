package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Body struct {
	Token string `json:"token"`
	IP    string `json:"ip"`
	Cmd   string `json:"cmd"`
}

var configAdd = ""
var objConfig = ""
var token = ""

func main() {

	flag.StringVar(&configAdd, "t", "", "模版文件 -t config.tmp")
	flag.StringVar(&objConfig, "o", "", "目标文件 -o config.conf")
	flag.StringVar(&token, "token", "", "token -token 1234")
	flag.Parse()

	if configAdd == "" || objConfig == "" {
		println("set -t and -o")
		return
	}

	http.Handle("/api/update/ip", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			var body Body
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&body)
			if err != nil {
				println("err--->", err)
				return
			}
			if token != "" && token != body.Token {
				println("token err--->", time.Now().Format("2006-01-02 15:04"))
				return
			}
			temp(body)
		}
	}))
	println(time.Now().Format("2006-01-02 15:04"), " server begin")
	http.ListenAndServe(":8080", nil)
}

func temp(body Body) {
	tmpl, err := template.ParseFiles(configAdd)
	if err != nil {
		println(err)
		return
	}

	f, err := os.Create(objConfig)
	if err != nil {
		log.Println("file: ", err)
		return
	}
	defer f.Close()
	tmpl.Execute(f, body)
	runCommand(body.Cmd)
}

func runCommand(s string) {
	ss := strings.Split(s, " ")
	star := ss[0]
	end := ss[1:(len(ss))]
	cmd := exec.Command(star, end...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
