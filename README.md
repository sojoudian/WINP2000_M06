# WINP2000_M06

## Compile for Linux:
Run the following command to compile for Linux:\
'''bash
GOOS=linux GOARCH=amd64 go build -o helloserver-linux main.go
'''
## Compile for Windows:
Run the following command to compile for Windows:\
'''bash
GOOS=windows GOARCH=amd64 go build -o helloserver-windows.exe main.go
'''

## Compile for macOS:
Run the following command to compile for macOS (Darwin):\
'''bash
GOOS=darwin GOARCH=arm64 go build -o helloserver-mac main.go

'''
