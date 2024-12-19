/**
 * @Author tanchang
 * @Description DB接口
 * @Date   2024/12/18 15:02
 * @File:  DbFacotryInterface
 * @Software: vscode
 **/

package DB

// TODO 待用
type DbFacotryInterface[D any] interface {
	CreateDB() DBInterface[D]
}
