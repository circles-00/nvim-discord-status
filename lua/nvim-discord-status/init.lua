local M = {}

M.setup = function(opts)
  opts = opts or {}
  opts.excludedDirs = opts.excludedDirs or {}

  if(opts.discordAppId == nil) then
    error("Missing required option 'discordAppId'")
  end

  M.opts             = opts
end

return M
