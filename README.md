# Neovim Discord Status
This plugin utilizes the [Discord Rich Presence API](https://discord.com/developers/docs/rich-presence/how-to) to display your current Neovim status on your Discord profile.

The plugin is written in Lua and GoLang.

The GoLang part is used to create a binary that communicates with Neovim and Discord. The Lua part is used to call the binary and set the status.
The GoLang application is a simple TCP server that listens for connections from Neovim and updates the status accordingly.

The lua part creates a TCP connection to the GoLang server and sends the status to be updated.

# Installation

Using Packer
```lua
use("circles-00/nvim-discord-status")
```

Using Vim-Plug
```vim
Plug 'circles-00/nvim-discord-status'
```

# Setup
You need to create a Discord application and get the client ID. You can do this by going to the [Discord Developer Portal](https://discord.com/developers/applications) and creating a new application.

Once you have the client ID, you can set it in your `init.vim` or `init.lua` file.

```lua
local nvim_discord_status = require("nvim-discord-status")

nvim_discord_status.setup({
   discordAppId  = "YOUR_APP_ID"
})
```


# Usage
The plugin will automatically start when you open Neovim.
The status will be updated when you open a new file or switch buffers.

The status will dissapear when you close all instances of Neovim.

# Screenshots
![image](https://github.com/circles-00/nvim-discord-status/assets/42126548/3afab290-5478-41d8-bd68-9e91a52edbaa)

