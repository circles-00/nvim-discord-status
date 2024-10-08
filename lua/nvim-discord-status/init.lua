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
function NvimDiscordStatus:new()
  local nvimDiscordStatus = setmetatable({
    actions = NvimDiscordStatusActions.new(),
  }, NvimDiscordStatus)

  return nvimDiscordStatus
end

local theNvimDiscordStatus = NvimDiscordStatus:new()

---@param self NvimDiscordStatus
---@param opts NvimDiscordStatusOptions?
function NvimDiscordStatus.setup(self, opts)
  if self ~= theNvimDiscordStatus then
    self = theNvimDiscordStatus
  end

  opts = opts or {}
  opts.excludedDirs = opts.excludedDirs or {}
  opts.cmdBinding = opts.cmdBinding or "<C-x>"

  if (opts.discordAppId == nil) then
    error("Missing required option 'discordAppId'")
  end

  self.opts = opts

  self.actions:connect(opts)
  self.actions:registerCommands(opts.cmdBinding)

  return self
end

function NvimDiscordStatus:excludeOrIncludeDirectory()
  self.actions:excludeOrIncludeDirectory()
end

return theNvimDiscordStatus
