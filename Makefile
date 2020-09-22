all: pull push

.PHONY: pull push build

pull:
	git pull

push:
	git pull
	git add .
	git commit -m "Update."
	git push origin master

build:
	cd cmd && \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../webhook && \
	upx ../webhook

copy:
	cp -f /c/Users/norton/go/src/webhook/webhook /d/projects/devops/golang/kevin/webhook/webhook
	cp -f /c/Users/norton/go/src/webhook/templates/config.yml /d/projects/devops/golang/kevin/webhook/templates/config.yml