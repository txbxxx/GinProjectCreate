/**
 * @Author tanchang
 * @Description 数据库连接工厂
 * @Date 2024/7/11 15:57
 * @File:  Sqlfacotry.go
 * @Software: vscode
 **/

package dbfacotry

import (
	"Go-WebCreate/utils/DB/interface"
	"github.com/go-redis/redis/v8"
	redisConn "Go-WebCreate/utils/DB/redis"
)

type SqlFacotry struct {
}

func (s *SqlFacotry) CreateDB() DB.DBInterface[*redis.Client] {
	// TODO 创建数据库连接
}
