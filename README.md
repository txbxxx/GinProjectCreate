# :trophy: 一个Gin项目的脚手架

## 脚手架内容
1.  `Gin` 前端Web框架 https://github.com/gin-gonic/gin
2.  `Gorm` ORM框架 https://github.com/jinzhu/gorm
3.  `MariaDB/Mysql` 数据库
4.  `Redis` 缓存 https://github.com/redis/redis
5.  `Jwt` 鉴权 https://github.com/golang-jwt/jwt/v5
6.  `Logrus` 日志 https://github.com/sirupsen/logrus
7.  `Gin-cors` 跨域 https://github.com/gin-contrib/cors
8.  `godotenv` 读取.env文件,方便写配置文件 https://github.com/joho/godotenv 
9.   `go-redis` redis操作工具 https://github.com/go-redis/redis/v8

## 目录内容设计
1.  `utils` 工具包，存放连接数据库，`Token`工具等等
2.  `router` 路由，设置对外访问的接口，一般就一个`route.go`文件
3.  `control` 控制层，处理业务逻辑，一般是整合`service`内容
4.  `service` 服务层，对数据和数据库进行操作
5.  `model` 模型层，定义数据模型
6.  `serialize` 序列化，将数据序列化成json格式
7.  `middleware` 中间件，对请求进行拦截，比如鉴权，日志记录等等
8.  `test` 测试用例 
