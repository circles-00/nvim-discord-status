package tcp

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"slices"
	"strings"
	"time"

	"discord_status/discord"
	"discord_status/utils"
)

var (
	DiscordAppId    = ""
	NumberOfClients = 0
	timeout         *time.Timer
)

const TIMEOUT_DURATION = 30 // seconds
var excludedDirsFile = fmt.Sprintf("%s/nvim-discord-status-excludedDirs.txt", os.Getenv("HOME"))

const (
	REDACT_COMMAND   = "redact"
	UNREDACT_COMMAND = "unredact"
)

func handlePanicTCPClient(err error) bool {
	if err != nil {
		NumberOfClients = NumberOfClients - 1
		fmt.Println("Error reading data from client:", err, NumberOfClients)

		if NumberOfClients == 0 && DiscordAppId != "" {
			// panic("No more clients")
			timeout = time.NewTimer(time.Duration(TIMEOUT_DURATION) * time.Second)

			go func() {
				<-timeout.C
				panic("No more clients")
			}()
		}

		return true
	}

	return false
}

func handleInitialClientConnection(message string, startTime time.Time, excludedDirsArray *[]string) bool {
	if strings.Contains(message, "connect") {
		onConnectArguments := strings.Split(message, ":")

		DiscordAppId = onConnectArguments[1]
		excludedDirs := onConnectArguments[2]

		err := json.Unmarshal([]byte(excludedDirs), &excludedDirsArray)
		if err != nil {
			fmt.Println("Error parsing excluded dirs:", err)
		}

		fmt.Printf("Message: %s\n", message)
		discord.UpdateDiscordPresence(&DiscordAppId, startTime, nil, nil, false)

		return true
	}

	return false
}

func isRedacted(excludedDirsArray []string, cleanDirPath string) bool {
	data, err := os.ReadFile(excludedDirsFile)
	if err != nil {
		fmt.Println("Error reading file, creating new file:", err)
	}

	exludedDirsFromFile := strings.Split(string(data), "\n")

	return slices.Contains(exludedDirsFromFile, cleanDirPath) || slices.Contains(excludedDirsArray, cleanDirPath)
}

// Note: We are executing this command even if the redact command is sent, so we can update the presence
func handleStandardConnection(message string, excludedDirsArray []string, startTime time.Time) {
	cleanDirPath, filename, gitRepo := utils.ExtractStatusParams(message)

	isRedacted := isRedacted(excludedDirsArray, cleanDirPath)

	discord.UpdateDiscordPresence(&DiscordAppId, startTime, &filename, &gitRepo, isRedacted)
}

func handleRedactCommand(message string, excludedDirsArray *[]string) string {
	if !strings.Contains(message, "redact") {
		return ""
	}

	cleanDirPath, filename, gitRepo := utils.ExtractStatusParams(message)
	fmt.Println(cleanDirPath, filename, gitRepo)

	data, err := os.ReadFile(excludedDirsFile)
	if err != nil {
		fmt.Println("Error reading file, creating new file:", err)
		os.Create(excludedDirsFile)
	}

	exludedDirsFromFile := strings.Split(string(data), "\n")

	var command string

	if slices.Contains(exludedDirsFromFile, cleanDirPath) || slices.Contains(*excludedDirsArray, cleanDirPath) {
		exludedDirsFromFile = utils.RemoveStringFromSlice(exludedDirsFromFile, cleanDirPath)
		fmt.Println("Removed", exludedDirsFromFile)

		command = UNREDACT_COMMAND
	} else {
		exludedDirsFromFile = append(exludedDirsFromFile, cleanDirPath)

		command = REDACT_COMMAND
	}

	file, err := os.OpenFile(excludedDirsFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0o644)
	for _, dir := range exludedDirsFromFile {
		if strings.TrimSpace(dir) == "" {
			continue
		}

		file.WriteString(dir + "\n")
	}

	file.Close()

	return command
}

func HandleTCPClient(conn net.Conn, startTime time.Time) {
	defer conn.Close()
	if timeout != nil {
		timeout.Stop()
		timeout = nil
	}

	// Create a buffer to read data into
	buffer := make([]byte, 1024)

	var excludedDirsArray []string

	for {
		// Read data from the client
		bufferData, err := conn.Read(buffer)

		isClientClosed := handlePanicTCPClient(err)
		if isClientClosed {
			return
		}

		// format: filePath:gitRepo
		message := string(buffer[:bufferData])

		isInitialConnection := handleInitialClientConnection(message, startTime, &excludedDirsArray)

		if isInitialConnection {
			continue
		}

		redactCommand := handleRedactCommand(message, &excludedDirsArray)
		handleStandardConnection(message, excludedDirsArray, startTime)

		if redactCommand != "" {
			conn.Write([]byte(redactCommand))
		}
	}
}
