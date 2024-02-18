package tree

import (
	"fmt"
	_ "io/ioutil"
	"os"
	_ "os"
	"path/filepath"
	_ "path/filepath"
)

// FolderNode 表示一个文件夹结构的节点。
// 它可以是一个目录或文件。对于目录节点，Children字段包含其子节点
type FolderNode struct {
	Name     string       // 文件或目录的名称
	IsDir    bool         // 指示该节点是否为目录
	Children []FolderNode // 子节点列表
}

// GenerateFolderStructure 生成给定路径的文件夹结构

func GenerateFolderStructure(rootPath string) (FolderNode, error) {
	return generateStructure(rootPath)
}

// generateStructure 是GenerateFolderStructure的递归帮助函数
func generateStructure(path string) (FolderNode, error) {
	node := FolderNode{
		Name:     filepath.Base(path),
		IsDir:    true,
		Children: nil,
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return FolderNode{}, err
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return FolderNode{}, err // 在递归中传播错误
		}
		if entry.IsDir() {
			childNode, err := generateStructure(filepath.Join(path, entry.Name()))
			if err != nil {
				return FolderNode{}, err // 在递归中传播错误
			}
			node.Children = append(node.Children, childNode)
		} else {
			node.Children = append(node.Children, FolderNode{Name: entry.Name(), IsDir: false})
		}
		// 未来可能会使用 info 变量的示例
		fmt.Println(info.Size()) // 打印文件大小，实际开发时可能会根据文件属性做更复杂的逻辑处理
	}

	return node, nil
}
