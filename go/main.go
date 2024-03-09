package main

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/hugolgst/rich-go/client"
)

var discordAppId = ""

func updateDiscordPresence(discordAppId *string, t time.Time, filename *string, gitRepo *string) {
	if discordAppId == nil || len(*discordAppId) == 0 {
		panic("Discord App ID is required")
	}

	err := client.Login(*discordAppId)

	if err != nil {
		panic(err)
	}

	var activity = client.Activity{
		LargeImage: "https://static-00.iconduck.com/assets.00/apps-neovim-icon-2048x2048-21jvoi4h.png",
		LargeText:  "Neovim is the best editor ever",
		//	SmallImage: "https://static-00.iconduck.com/assets.00/apps-neovim-icon-2048x2048-21jvoi4h.png",
		//	SmallText:  "NeoVim the best editor ever",
		Timestamps: &client.Timestamps{
			Start: &t,
		},
	}

	if gitRepo != nil {
		activity.Details = *gitRepo
	}

	if filename != nil {
		activity.State = *filename
	}

	err = client.SetActivity(activity)

	if err != nil {
		panic(err)
	}

}

func handleTCPClient(conn net.Conn, startTime time.Time) {
	defer conn.Close()

	// Create a buffer to read data into
	buffer := make([]byte, 1024)

	for {
		// Read data from the client
		bufferData, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		message := string(buffer[:bufferData])

		if strings.Contains(message, "connect") {
			discordAppIdFromSocket := strings.Split(message, ":")[1]
			discordAppId = discordAppIdFromSocket
			updateDiscordPresence(&discordAppId, startTime, nil, nil)

			continue
		}

		dirParts := strings.Split(message, "/")
		filenameAndGitRepo := dirParts[len(dirParts)-1]
		filename := strings.Split(filenameAndGitRepo, ":")[0]
		gitRepo := strings.Split(filenameAndGitRepo, ":")[1]

		if len(filename) > 0 {
			updateDiscordPresence(&discordAppId, startTime, &filename, &gitRepo)
		}
	}
}

func main() {

	// Listen for incoming connections
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	t := time.Now()

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client connection in a goroutine
		go handleTCPClient(conn, t)
	}
}
