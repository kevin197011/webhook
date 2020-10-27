all: pull push

.PHONY: pull push build

pull:
	git pull

push: fmt
	git pull
	git add .
	git commit -m "Update."
	git push origin master

build: fmt
	go mod tidy
	cd cmd && \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../webhook && \
	upx ../webhook

fmt:
	sudo chown -R norton:norton ./
	gofmt -l -w ./