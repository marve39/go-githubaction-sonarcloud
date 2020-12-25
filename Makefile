build: 
	mkdir -p bin
	go build -v -o bin/server main.go 
test:
	mkdir -p bin
	go test -v ./  -coverprofile="bin/coverage.out" -covermode=count -json > bin/report.json
get-dep:
	go get -v -t -d ./
check-lint:
	golangci-lint run
sonar-version:
	cat version.txt| awk '{print "\nsonar.projectVersion="$$1}' >> sonar-project.properties