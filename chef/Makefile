HANDLER ?= main
PACKAGE ?= $(HANDLER)
GOPATH  ?= $(HOME)/go
GOOS    ?= linux
GOOSDEV	?= $(shell uname -s)
GOARCH  ?= amd64
S3TMPBUCKET	?= pahud-tmp

STACKNAME	?= telegramApiGW

LAMBDAARN ?= "arn:aws:lambda:ap-northeast-1:147196629674:function:chef_node_checkin"
WORKDIR = $(CURDIR:$(GOPATH)%=/go%)
ifeq ($(WORKDIR),$(CURDIR))
	WORKDIR = /tmp
endif

run:
	go run chefnode.go slack.go kms.go main.go

build:
	@echo "Building..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags='-w -s' -o $(HANDLER)

init:
	go get github.com/dustin/go-humanize
	go get github.com/go-chef/chef
	go get github.com/jedib0t/go-pretty/table
	go get github.com/aws/aws-sdk-go/aws
	go get github.com/joho/godotenv
	go get github.com/aws/aws-lambda-go/lambda
zip:
	zip -r main.zip encrypted_pem.txt main
lambda: clean build zip 
	aws lambda update-function-code --region ap-northeast-1 --profile prd --function-name arn:aws:lambda:ap-northeast-1:144:function:chef_node_checkin --zip-file fileb://./main.zip
clean:
	rm -f main main.zip
