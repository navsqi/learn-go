1. go mod init go-api
2. go get github.com/githubnemo/CompileDaemon
3. go install github.com/githubnemo/CompileDaemon
4. go get -u github.com/gin-gonic/gin
5. go get -u gorm.io/gorm
6. go get -u gorm.io/driver/postgres
7. go get github.com/joho/godotenv


go run migrate/migrate.go
CompileDaemon -command="./go-api" 