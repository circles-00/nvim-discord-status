local host, port = "127.0.0.1", 49069
local TCPClient = require('nvim-discord-status.tcp_client')
local utils = require('nvim-discord-status.utils')
local setupOpts = require('nvim-discord-status')
local JSON = require('JSON')

local client = TCPClient:new()

vim.api.nvim_create_autocmd({ "VimEnter" }, {
  pattern = { "*" },
  callback = function()
    local script_path = utils.get_script_path()

    if (script_path == nil) then
      return
    end

    local path_to_script = utils.removeLastThreeParts(script_path)

    local absolute_path = "/" .. path_to_script .. "/" .. "go/discord_status > log.txt 2>&1 &"
    io.popen(absolute_path)

    -- Wait for the TCP server to start
    utils.asyncSleep(0.5, function()
      client:connect(host, port)
      client:send("connect:" .. setupOpts.opts.discordAppId .. ":" .. JSON:encode(setupOpts.opts.excludedDirs))
    end)
  end
})

vim.api.nvim_create_autocmd({ "BufEnter" }, {
  pattern = { "*" },
  callback = function()
    local filename = vim.fn.expand('%:p')
    local handle = io.popen("basename `git rev-parse --show-toplevel`")
    if (handle == nil) then
      return
    end

    local result = handle:read("*a")
    client:send(filename .. ":" .. result);

    handle:close()
  end
})

-- Cleanup the TCP client and the Go program
vim.api.nvim_create_autocmd({ "VimLeavePre" }, {
  pattern = { "*" },
  callback = function()
    client:close()
  end
})
