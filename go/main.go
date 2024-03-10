package main

import (
	"encoding/json"
	"fmt"
	"net"
	"slices"
	"strings"
	"time"

	"github.com/hugolgst/rich-go/client"
)

var discordAppId = ""
var numberOfClients = 0

func updateDiscordPresence(discordAppId *string, t time.Time, filename *string, gitRepo *string, isRedacted bool) {
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

	if gitRepo != nil && !isRedacted {
		activity.Details = *gitRepo
	}

	if filename != nil {
		if isRedacted {
			activity.State = "[REDACTED]"
		} else {
			activity.State = *filename
		}
	}

	err = client.SetActivity(activity)

	if err != nil {
		panic(err)
	}

}

func handlePanicTCPClient(err error) bool {
	if err != nil {
		numberOfClients = numberOfClients - 1
		fmt.Println("Error reading data from client:", err, numberOfClients)

		if numberOfClients == 0 && discordAppId != "" {
			panic("No more clients")
		}

		return true
	}

	return false
}

func handleInitialClientConnection(message string, startTime time.Time, excludedDirsArray *[]string) bool {
	if strings.Contains(message, "connect") {
		onConnectArguments := strings.Split(message, ":")

		discordAppId = onConnectArguments[1]
		excludedDirs := onConnectArguments[2]

		err := json.Unmarshal([]byte(excludedDirs), &excludedDirsArray)

		if err != nil {
			fmt.Println("Error parsing excluded dirs:", err)
		}

		updateDiscordPresence(&discordAppId, startTime, nil, nil, false)

		return true
	}

	return false
}

func extractStatusParams(message string) (string, string, string) {
	// format: filePath:gitRepo
	dirParts := strings.Split(message, "/")
	filenameAndGitRepo := dirParts[len(dirParts)-1]

	fileParts := strings.Split(filenameAndGitRepo, ":")
	filename := fileParts[0]
	gitRepo := fileParts[1]

	var cleanDirPath = strings.Split(message, strings.TrimSpace(gitRepo))[0] + strings.TrimSpace(gitRepo)
	if strings.Contains(cleanDirPath, "fugitive://") {
		cleanDirPath = strings.Split(cleanDirPath, "fugitive://")[1]
	}

	return cleanDirPath, filename, gitRepo
}

func handleTCPClient(conn net.Conn, startTime time.Time) {
	defer conn.Close()

	// Create a buffer to read data into
	buffer := make([]byte, 1024)

	var excludedDirsArray []string

	for {
		// Read data from the client
		bufferData, err := conn.Read(buffer)

		var isClientClosed = handlePanicTCPClient(err)
		if isClientClosed {
			return
		}

		// format: filePath:gitRepo
		message := string(buffer[:bufferData])

		isInitialConnection := handleInitialClientConnection(message, startTime, &excludedDirsArray)

		if isInitialConnection {
			continue
		}

		cleanDirPath, filename, gitRepo := extractStatusParams(message)

		var isRedacted = slices.Contains(excludedDirsArray, cleanDirPath)

		updateDiscordPresence(&discordAppId, startTime, &filename, &gitRepo, isRedacted)
	}
}

func main() {

	// Listen for incoming connections
	const port = 49069
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server is listening on port %d\n", port)

	t := time.Now()

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		numberOfClients = numberOfClients + 1
		// Handle client connection in a goroutine
		go handleTCPClient(conn, t)
	}
}
