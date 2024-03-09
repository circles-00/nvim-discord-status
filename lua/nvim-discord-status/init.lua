local M = {}

M.setup = function(opts)
  opts = opts or {}

  local discordAppId = opts.discordAppId or error("Missing required option 'discordAppId'")

  opts.discordAppId  = discordAppId

  M.opts             = opts
end

return M
