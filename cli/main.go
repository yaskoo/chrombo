package main

import (
	"encoding/json"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"net/http"
	"os"
)

type Page struct {
	DebuggerUrl string `json:"webSocketDebuggerUrl"`
}

func main() {
	if len(os.Args) != 2 {
		println("give me single argument: 'host:port' to connect to")
		os.Exit(1)
	}

	res, err := http.Get("http://" + os.Args[1] + "/json")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var data []byte
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var pages []Page
	json.Unmarshal(data, &pages)

	println(pages[0].DebuggerUrl)

	var ws *websocket.Conn
	ws, err = websocket.Dial(pages[0].DebuggerUrl, "", "http://localhost/")
	if err != nil {
		panic(err)
	}

	_, err = ws.Write([]byte(`
		"id": "ecddba4e-b121-4604-bc73-79c7b72cfeb7",
		"method": "Page.navigate",
		"params": {
			url: "https://www.google.bg"
		}
	`))

	if err != nil {
		panic(err)
	}
}
