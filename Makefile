build-HelloGopher:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o ./bootstrap ./main.go
	cp ./bootstrap $(ARTIFACTS_DIR)/.

sam-build:
	sam build -t ./local/template.yml

sam-run:
	sam local invoke HelloGopher -e ./local/event.json