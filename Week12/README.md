# External API and save result to MongoDB

```bash
git clone https://github.com/sojoudian/WINP2000_M06.git
curl https://api.ipify.org\?format\=json
docker run -d --name mongodb -p 27017:27017 -v mongodb:/data/db mongo:latest

go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options

"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
```
## How to test the project
 ```bash
go test -v
ls
go test
go test -v
go test -v -run TestFileServer
go test ./...
```
