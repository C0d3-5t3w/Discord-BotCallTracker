package main

import (
	"bufio"
	"fmt"
	"godiscopwn/bot"
	"godiscopwn/server"
	"net/http"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Read Discord bot token
	fmt.Print("Enter Discord Bot Token: ")
	discordToken, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading token,", err)
		return
	}
	discordToken = discordToken[:len(discordToken)-1] // Remove the newline character

	// Read port number
	fmt.Print("Enter port number: ")
	port, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading port,", err)
		return
	}
	port = port[:len(port)-1]

	go server.StartWebSocketServer(port)
	go bot.StartDiscordBot(discordToken)

	http.HandleFunc("/", serveHome)
	fmt.Printf("Serving index.html on :%s\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
