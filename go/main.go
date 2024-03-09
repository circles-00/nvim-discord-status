package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/hugolgst/rich-go/client"
	"github.com/joho/godotenv"
)

func updateDiscordPresence(t time.Time, filename *string, gitRepo *string) {
  dotEnvError := godotenv.Load()
  if dotEnvError != nil {
    panic(dotEnvError)
  }

  appIdFromEnv := os.Getenv("DISCORD_APP_ID")
  fmt.Printf("App ID: %s", appIdFromEnv)
	err := client.Login(appIdFromEnv)

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
		message, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		dir := string(buffer[:message])
		dirParts := strings.Split(dir, "/")
		filenameAndGitRepo := dirParts[len(dirParts)-1]
    filename := strings.Split(filenameAndGitRepo, ":")[0]
    gitRepo := strings.Split(filenameAndGitRepo, ":")[1]

		if len(filename) > 0 {
			updateDiscordPresence(startTime, &filename, &gitRepo)
		}

		// Process and use the data (here, we'll just print it)
		fmt.Printf("Received: %s\n", filename)
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

	updateDiscordPresence(t, nil, nil)

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
