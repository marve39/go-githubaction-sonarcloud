test:
	mkdir bin
	go test ./  -coverprofile="bin/coverage.out" -covermode=count -json > bin/report.json