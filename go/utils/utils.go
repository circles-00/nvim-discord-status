package utils

import "strings"

func ExtractStatusParams(message string) (string, string, string) {
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

