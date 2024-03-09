local M = {}

M.removeLastThreeParts = function(path)
  local parts = {}
  for part in path:gmatch("[^/\\]+") do
    table.insert(parts, part)
  end

  -- Ensure we have more than 3 parts
  if #parts > 3 then
    -- Remove the last 3 parts
    for i = 1, 3 do
      table.remove(parts, #parts)
    end
  else
    -- If there are 3 or fewer parts, return the original path
    return path
  end

  -- Concatenate the remaining parts with the path separator
  return table.concat(parts, "/")
end


M.get_script_path = function()
  local src = debug.getinfo(1).source
  if src:sub(1, 1) == "@" then
    return src:sub(2)
  else
    return nil
  end
end

M.asyncSleep = function(seconds, callback)
  local command = "sleep " .. seconds
  local job_id = vim.fn.jobstart(command, {
    on_exit = function(_, code, _)
      if code == 0 then
        callback()
      else
        print("Error: Sleep command failed.")
      end
    end
  })
  return job_id
end

return M
