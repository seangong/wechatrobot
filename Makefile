
run:
	go run cmd/main.go --RobotKey="899220cd-5ed6-44ad-b053-f3785033da7f"


.PHONY: build
build:
	go build -o wechatrobot cmd/main.go
