说明项目的运行环境、依赖安装步骤和启动方式。
项目使用go语言,mysql数据库,可部署到windows,linux,mac具备对应golang开发环境和mysql的机器上
依赖准备包括
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/gin-gonic/gin
go get github.com/spf13/viper
go get -u go.uber.org/zap
go get github.com/golang-jwt/jwt/v5

启动方式
go run main.go