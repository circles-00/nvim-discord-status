package discord

import (
	"time"

	"github.com/hugolgst/rich-go/client"
)

func UpdateDiscordPresence(discordAppId *string, t time.Time, filename *string, gitRepo *string, isRedacted bool) {
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
