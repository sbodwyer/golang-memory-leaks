# Run the server

`go run main.go`

Note: all endpoints are on port 8080

# Call an endpoint

eg
`curl http://localhost:8080/myLovelyFunction1`

# Collect a heap

`curl http://localhost:8080/debug/pprof/heap > "heap-$(date +%T).out"`

# Diff two heaps on your browser

`go tool pprof -http=:8802 -diff_base <path to base heap file> <path to the heap file you wish to compare>`