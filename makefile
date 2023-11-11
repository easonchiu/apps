dev:
	air

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./run -tags=jsoniter -trimpath -ldflags '-s -w'

push: build
	ssh ysgame 'sudo systemctl stop ysgamestudio'
	scp ./run ysgame:/usr/go/ysgamestudio
	ssh ysgame 'sudo systemctl restart ysgamestudio && sudo systemctl status ysgamestudio'
	rm -f ./run
