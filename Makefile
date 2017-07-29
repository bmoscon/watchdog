all: demo

demo:
	go build watchdog.go
	go build -o example demo/demo.go


clean:
	rm -f watchdog
	rm -f watchdog.log
	rm -f example

.PHONY: demo