// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
)

var addr = flag.String("addr", "81.177.135.38:80", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/api/branches/ws/get-nearest-branches-with-talons", RawQuery: "longitude=37.767536&&latitude=55.419290"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	//var id string
	//var talon model.Talon
	go func() {
		defer close(done)

		_, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		var data []model.BusinessResponse
		err = json.Unmarshal(message, &data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Получено первое сообщение: %+v\n", data)

		for {

			_, message, err := c.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}

			var data model.BusinessResponse
			err = json.Unmarshal(message, &data)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("Получено сообщение: %+v\n", data)

			//err := c.ReadJSON(&id)
			//if err != nil {
			//	log.Println("read:", err)
			//	return
			//}
			//log.Printf("recv: %s", id)

			//message := map[string]string{}
			//
			//err := c.ReadMessage(&message)
			//if err != nil {
			//	log.Println("read:", err)
			//	return
			//}
			//log.Println(message)
		}
	}()

	for {
	}
	//ticker := time.NewTicker(time.Second)
	//defer ticker.Stop()
	//
	//for {
	//	select {
	//	case <-done:
	//		return
	//	case t := <-ticker.C:
	//		err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
	//		if err != nil {
	//			log.Println("write:", err)
	//			return
	//		}
	//	case <-interrupt:
	//		log.Println("interrupt")
	//
	//		// Cleanly close the connection by sending a close message and then
	//		// waiting (with timeout) for the server to close the connection.
	//		err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	//		if err != nil {
	//			log.Println("write close:", err)
	//			return
	//		}
	//		select {
	//		case <-done:
	//		case <-time.After(time.Second):
	//		}
	//		return
	//	}
	//}
}
