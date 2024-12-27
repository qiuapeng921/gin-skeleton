buildChat:
	gf build cmd/test/chatgpt.go  -a amd64,arm64 -s windows,darwin -p ./bin/chat

sheetNames:
	gf build cmd/test/sheetNames.go  -a amd64 -s windows -p ./bin/sheet