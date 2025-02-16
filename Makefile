install:
	go build -o treedo main.go
	sudo mv treedo /usr/local/bin/

uninstall:
	sudo rm /usr/local/bin/treedo 