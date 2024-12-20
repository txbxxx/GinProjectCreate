/**
 * @Author tanchang
 * @Description 模型创建实现接口
 * @Date 2024/12/20 11:40
 * @File:  createtable
 * @Software: vscode
**/

package model

// CreateTable 需要创建的表模型需实现接口
type CreateTable interface {
	IsCreate() bool
}
