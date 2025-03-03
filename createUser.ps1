$headers = @{
    "Content-Type" = "application/json"
}

$body = @{
    name = "Amaru"
} | ConverTo-Json 

Invoke-WebRequest -Uri "http://localhost:3001/api/createUser" -Method POST -Headers $headers -Body $body