package tcp

import (
	"discord_status/discord"
	"discord_status/utils"
	"encoding/json"
	"fmt"
	"net"
	"slices"
	"strings"
	"time"
)

var discordAppId = ""
var NumberOfClients = 0

func handlePanicTCPClient(err error) bool {
	if err != nil {
		NumberOfClients = NumberOfClients - 1
		fmt.Println("Error reading data from client:", err, NumberOfClients)

		if NumberOfClients == 0 && discordAppId != "" {
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

		discord.UpdateDiscordPresence(&discordAppId, startTime, nil, nil, false)

		return true
	}

	return false
}

func handleStandardConnection(message string, excludedDirsArray []string, startTime time.Time) {
	cleanDirPath, filename, gitRepo := utils.ExtractStatusParams(message)

	var isRedacted = slices.Contains(excludedDirsArray, cleanDirPath)

	discord.UpdateDiscordPresence(&discordAppId, startTime, &filename, &gitRepo, isRedacted)

}

func HandleTCPClient(conn net.Conn, startTime time.Time) {
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

		handleStandardConnection(message, excludedDirsArray, startTime)
	}
}


