/**
 * @Author tanchang
 * @Description DB接口
 * @Date   2024/12/18 15:02
 * @File:  DbInterface
 * @Software: vscode
 **/

package DB

type DBInterface[D any] interface {
	Connect() D
}

