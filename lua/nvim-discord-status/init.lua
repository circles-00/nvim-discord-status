local NvimDiscordStatusActions = require("nvim-discord-status.actions")

---@class NvimDiscordStatusOptions
---@field discordAppId string
---@field excludedDirs string[]
---@field cmdBinding string

---@class NvimDiscordStatus
---@field actions NvimDiscordStatusActions
local NvimDiscordStatus = {}

NvimDiscordStatus.__index = NvimDiscordStatus

---@return NvimDiscordStatus
function NvimDiscordStatus.new()
  local self = setmetatable({
    actions = NvimDiscordStatusActions.new(),
  }, NvimDiscordStatus)

  return self
end

---@param opts NvimDiscordStatusOptions
function NvimDiscordStatus:setup(opts)
  opts = opts or {}
  opts.excludedDirs = opts.excludedDirs or {}
  opts.cmdBinding = opts.cmdBinding or "<C-x>"

  if (opts.discordAppId == nil) then
    error("Missing required option 'discordAppId'")
  end


  self.opts = opts

  self.actions:connect(opts)
end


--endexcludeOrIncludeDirectory = function ()
--  print("implement")
--end

-- local function registerCommands(binding)
--   vim.cmd('command! Redact lua require("nvim-discord-status").excludeOrIncludeDirectory()')
--   vim.keymap.set('n', binding, M.excludeOrIncludeDirectory, {noremap = true, silent = true})
-- end

return NvimDiscordStatus.new()
