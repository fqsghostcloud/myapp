.PHONY: package
package:
	go build -o myapp && docker build -t myapp:0.1 .
