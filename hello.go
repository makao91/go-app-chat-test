package main

import (
	"encoding/json"
	"fmt"

	// "log"
	// "os"

	// "time"

	"github.com/go-redis/redis"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var (
	rdb *redis.Client
)

type ChatMessage struct {
	From    string
	Message string
	Time    int64
}

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type hello struct {
	app.Compo

	Name            string
	Raw_messages    []string
	Messages        []ChatMessage
	updateAvailable bool
}

func (h *hello) OnAppUpdate(ctx app.Context) {
	h.updateAvailable = ctx.AppUpdateAvailable() // Reports that an app update is available.
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *hello) Render() app.UI {
	// data := h.Messages
	// fmt.Println(h.Messages[0])
	// for i := range data { // Looping over the root slice
	// 	// fmt.Println(data[i])
	// }
	// fmt.Print(h.messages)
	// fmt.Print(h.raw_messages)
	fmt.Println(h.updateAvailable)
	fmt.Println(h.Messages)
	fmt.Println("render")
	return app.Text(h.Messages)
}

// func (h *hello) OnNav(ctx app.Context) {
// 	chatMessages, _ := Rdb.LRange("chat_messages", 0, -1).Result()
// 	for _, chatMessage := range chatMessages {
// 		var message ChatMessage
// 		json.Unmarshal([]byte(chatMessage), &message)
// 		h.Messages = append(h.Messages, message)
// 	}
// 	// fmt.Println(h.Messages)
// 	// fmt.Println(chatMessages)
// 	// fmt.Println("nav")
// 	// fmt.Print(err)
// 	h.Name = "lutas"
// }

func (h *hello) OnPreRender(ctx app.Context) {
	chatMessages, _ := Rdb.LRange("chat_messages", 0, -1).Result()
	for _, chatMessage := range chatMessages {
		var message ChatMessage
		json.Unmarshal([]byte(chatMessage), &message)
		h.Messages = append(h.Messages, message)
	}
	// fmt.Println(h.messages)
	// fmt.Println(h.raw_messages)
}

// func (h *hello) OnMount(ctx app.Context) {
// 	chatMessages, _ := Rdb.LRange("chat_messages", 0, -1).Result()
// 	for _, chatMessage := range chatMessages {
// 		var message ChatMessage
// 		json.Unmarshal([]byte(chatMessage), &message)
// 		h.messages = append(h.messages, message)
// 	}
// 	fmt.Println(h.messages)
// 	fmt.Println(h.raw_messages)
// }
