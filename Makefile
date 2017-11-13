.PHONY: package
package:
	go build -o run && docker build -t myapp:0.5 .
