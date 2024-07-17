dev:
	air

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./run -tags=jsoniter -trimpath -ldflags '-s -w'

docker: build
	docker build -t easonchiu19860129/www:latest .
	docker push easonchiu19860129/www:latest
	rm -rf ./run
