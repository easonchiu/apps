build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./ysgame -tags=jsoniter -trimpath -ldflags '-s -w'

push: build
	ssh ys 'sudo systemctl stop ysgame'
	scp ./ysgame ys:/usr/go
	ssh ys 'sudo systemctl restart ysgame && sudo systemctl status ysgame'
	rm -f ./ysgame
