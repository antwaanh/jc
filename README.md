# JumpCloud Hash Server

## Requirements
### Hash and Encode a Password String
Write an HTTP server that listens on a given port. Your server should be able to process multiple connections simultaneously. And provide the following endpoints:

| Method      | Endpoint    | Description
| ----------- | ----------- | -------------------------------
| POST        | /hash       | Accepts password string to hash
| GET         | /hash/:id   | Retrieves hashed password
| GET         | /stats      | Returns server statistics
| POST        | /shutdown   | Gracefully shutdown server

## Run
1. Clone repo
2. `cd jc`
3. `go run jc` 


