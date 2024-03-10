local TCPClient = require('nvim-discord-status.tcp_client')
local utils = require('nvim-discord-status.utils')
local JSON = require('JSON')

---@class NvimDiscordStatusActions
---@field host string
---@field port number
---@field client TCPClient
local NvimDiscordStatusActions = {}

NvimDiscordStatusActions.__index = NvimDiscordStatusActions

---@return NvimDiscordStatusActions
function NvimDiscordStatusActions.new()
  return setmetatable({
    host = "127.0.0.1",
    port = 49069,
    client = TCPClient:new()
  }, NvimDiscordStatusActions)
end

---@param setupOpts NvimDiscordStatusOptions
function NvimDiscordStatusActions:connect(setupOpts)
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
        self.client:connect(self.host, self.port)
        self.client:send("connect:" .. setupOpts.discordAppId .. ":" .. JSON:encode(setupOpts.excludedDirs))
      end)
    end
  })

  vim.api.nvim_create_autocmd({ "BufEnter" }, {
    pattern = { "*" },
    callback = function()
      local result = utils.getFilePathAndGitRepo()

      if (result == nil) then
        return
      end

      self.client:send(result.filename .. ":" .. result.git_repo);
    end
  })

  -- Cleanup the TCP client and the Go program
  vim.api.nvim_create_autocmd({ "VimLeavePre" }, {
    pattern = { "*" },
    callback = function()
      self.client:close()
    end
  })
end

function NvimDiscordStatusActions:excludeOrIncludeDirectory()
  local result = utils.getFilePathAndGitRepo()

  if (result == nil) then
    return
  end

  self.client:send("redact:" .. result.filename .. ":" .. result.git_repo);
end

function NvimDiscordStatusActions:registerCommands(binding)
  vim.cmd('command! Redact lua require("nvim-discord-status"):excludeOrIncludeDirectory()')
  vim.keymap.set('n', binding, '<cmd>Redact<CR>', { noremap = true, silent = true })
end

return NvimDiscordStatusActions
