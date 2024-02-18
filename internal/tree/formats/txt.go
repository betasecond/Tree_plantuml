package formats

import (
	"Tree_plantuml/internal/tree"
	"fmt"
	"strings"
)

// GenerateTXTOutput 接收一个 FolderNode 根节点，然后生成其文本表示形式
func GenerateTXTOutput(node tree.FolderNode, level int) string {
	var builder strings.Builder
	indent := strings.Repeat("  ", level)

	if node.IsDir {
		builder.WriteString(fmt.Sprint("%s[%s]\n", indent, node.Name))
	} else {
		builder.WriteString(fmt.Sprintf("%s%s\n", indent, node.Name))
	}

	for _, child := range node.Children {
		builder.WriteString(GenerateTXTOutput(child, level+1)) // 递归处理子节点
	}

	return builder.String()

}
