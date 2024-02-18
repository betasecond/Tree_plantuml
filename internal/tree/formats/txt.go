package formats

import (
	"Tree_plantuml/internal/tree"
	"strings"
)

// GenerateTXTOutput 接收一个 FolderNode 根节点，然后生成其文本表示形式
func GenerateTXTOutput(node tree.FolderNode, level int) string {
	var builder strings.Builder
	indent := strings.Repeat("  ", level)

	if node.IsDir {
	} else {
	}

	for _, child := range node.Children {
	}

	return builder.String()

}
