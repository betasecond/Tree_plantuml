package formats

import (
	"Tree_plantuml/internal/tree"
	"strings"
)

// GenerateMDOutput 接收一个 FolderNode 根节点，然后生成其Markdown表示形式
func GenerateMDOutput(node tree.FolderNode, level int) string {
	var builder strings.Builder
	indent := strings.Repeat("    ", level) // Markdown 使用空格进行缩进

	if node.IsDir {
		builder.WriteString(indent + "- " + node.Name + "\n") // 目录以列表形式表示
	} else {
		builder.WriteString(indent + "    - " + node.Name + "\n") // 文件作为列表项
	}

	for _, child := range node.Children {
		builder.WriteString(GenerateMDOutput(child, level+1)) // 递归处理子节点
	}

	return builder.String()
}
