package discord

import (
	"fmt"
	"strings"
	"time"

	"github.com/hugolgst/rich-go/client"
)

func UpdateDiscordPresence(discordAppId *string, t time.Time, filename *string, gitRepo *string, isRedacted bool) {
	if discordAppId == nil || len(*discordAppId) == 0 {
		panic("Discord App ID is required")
	}

	fileExtension := ""
	if filename != nil {
		fileExtensionParts := strings.Split(*filename, ".")
		if len(fileExtensionParts) > 1 {
			fileExtension = fileExtensionParts[len(fileExtensionParts)-1]
		}
	}

	extensionImage := GetLanguageUrl(fmt.Sprintf(".%s", fileExtension), isRedacted)

	err := client.Login(*discordAppId)
	if err != nil {
		panic(err)
	}

	activity := client.Activity{
		LargeImage: extensionImage,
		SmallImage: "https://icons.iconarchive.com/icons/papirus-team/papirus-apps/512/nvim-icon.png",
		SmallText:  "NeoVim the best editor ever",
		Details:    "Idle", // Default
		Timestamps: &client.Timestamps{
			Start: &t,
		},
	}

	if gitRepo != nil && !isRedacted {
		activity.Details = fmt.Sprintf("Workspace: %s", *gitRepo)
	} else if isRedacted {
		activity.Details = ""
	}

	if filename != nil {
		if isRedacted {
			activity.State = "[REDACTED]"
		} else {
			activity.State = fmt.Sprintf("Editing %s", *filename)
		}
	}

	err = client.SetActivity(activity)

	if err != nil {
		panic(err)
	}
}
