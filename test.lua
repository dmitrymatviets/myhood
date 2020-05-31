-- Load URL paths from the file
function load_params_from_file(file)
  lines = {}

  -- Check if the file exists
  -- Resource: http://stackoverflow.com/a/4991602/325852
  local f=io.open(file,"r")
  if f~=nil then
    io.close(f)
  else
    -- Return the empty array
    return lines
  end

  -- If the file exists loop through all its lines
  -- and add them into the lines array
  for line in io.lines(file) do
    if not (line == '') then
      lines[#lines + 1] = line
    end
  end

  return lines
end

-- Load URL paths from file
params = load_params_from_file("./searchphrases.txt")

print("params: Found " .. #params .. " params")

request = function()
  math.randomseed(os.clock()*100000000000)
  -- Get the next paths array element
  param1 = params[math.random(1, #params)]
  param2 = params[math.random(1, #params)]

  wrk.method = "POST"
  wrk.body = '{"meta": {},"data": {"namePrefix": "' .. param1 ..'","surnamePrefix": "' .. param2 ..'","session": "972ce15e-e3ae-496f-9ba5-b7a250b8dfc5"}}'
  wrk.headers["Content-Type"] = "application/json"

  return wrk.format(nil, "/v1/user/search")

end