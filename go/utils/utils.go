package utils

import (
	"strings"
)

func ExtractStatusParams(message string) (string, string, string) {
	var cleanMessage string

	if strings.Contains(message, "redact:") {
		cleanMessage = strings.Split(message, "redact:")[1]
	} else {
		cleanMessage = message
	}

	// format: filePath:gitRepo
	dirParts := strings.Split(cleanMessage, "/")
	filenameAndGitRepo := dirParts[len(dirParts)-1]

	fileParts := strings.Split(filenameAndGitRepo, ":")
	var filename string
	var gitRepo string

	filename = fileParts[0]
	gitRepo = fileParts[1]

	var cleanDirPath = strings.Split(message, strings.TrimSpace(gitRepo))[0] + strings.TrimSpace(gitRepo)
	if strings.Contains(cleanDirPath, "fugitive://") {
		cleanDirPath = strings.Split(cleanDirPath, "fugitive://")[1]
	}

  if strings.Contains(cleanDirPath, "redact:") {
    cleanDirPath = strings.Split(cleanDirPath, "redact:")[1]
  }

	return cleanDirPath, filename, gitRepo
}

func RemoveStringFromSlice(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}

	return slice
}
