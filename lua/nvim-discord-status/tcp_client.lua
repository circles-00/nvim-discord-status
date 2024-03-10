-- tcp_client.lua

local socket = require("socket")

-- Create a table to hold the client object
---@class TCPClient
local TCPClient = {}

-- Constructor
function TCPClient:new()
    local obj = {
        socket = nil,
        host = nil,
        port = nil
    }
    setmetatable(obj, self)
    self.__index = self
    return obj
end

-- Connect to a TCP server
function TCPClient:connect(host, port)
    self.host = host
    self.port = port
    self.socket = socket.tcp()
    self.socket:settimeout(5)  -- Set a timeout, adjust as needed
    local success, err = self.socket:connect(host, port)
    if not success then
        return nil, "Failed to connect: " .. err
    end
    return true
end

-- Send data to the server
function TCPClient:send(data)
    if not self.socket then
        return nil, "Socket not initialized"
    end
    local bytes, err, partial = self.socket:send(data)
    if not bytes then
        return nil, "Error sending data: " .. err
    end
    return bytes
end

-- Receive data from the server
function TCPClient:receive()
    if not self.socket then
        return nil, "Socket not initialized"
    end
    local response, err, partial = self.socket:receive(100)
    if not response then
        return nil, "Error receiving data: " .. err
    end
    return response
end

-- Close the connection
function TCPClient:close()
    if self.socket then
        self.socket:close()
        self.socket = nil
    end
end

return TCPClient
