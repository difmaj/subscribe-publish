test: 
	- echo "Running tests..."
	- go vet $(go list ./...)
	- go test -race $(go list ./...) 

test-coverage:
	- echo "Running tests with coverage..."
	- go test -coverprofile /tmp/cover.out $(go list ./...) 
	- go tool cover -html=/tmp/cover.out -o /tmp/cover.html
	- google-chrome /tmp/cover.html
