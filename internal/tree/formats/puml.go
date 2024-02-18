package formats

import (
	"Tree_plantuml/internal/tree"
	"strings"
)

// cleanIdentifier 用于替换标识符中的特殊字符
func cleanIdentifier(identifier string) string {
	// 根据需要替换更多特殊字符
	return strings.NewReplacer(".", "_", "/", "_", "-", "_", " ", "_").Replace(identifier)
}

// GeneratePlantUMLContent 生成PlantUML内容
func GeneratePlantUMLContent(node tree.FolderNode, level int, levelLimit int) string {
	if level > levelLimit {
		return ""
	}

	var result strings.Builder
	identifier := cleanIdentifier(node.Name)

	if node.IsDir {
		result.WriteString("package \"" + node.Name + "\" as " + identifier + " {\n")
		for _, child := range node.Children {
			result.WriteString(GeneratePlantUMLContent(child, level+1, levelLimit))
		}
		result.WriteString("}\n")
	} else {
		result.WriteString("class " + identifier + "\n")
	}

	if level == 1 {
		// 如果是根节点，可能需要添加一些PlantUML特定的设置或者语法
		return "@startuml\n" + result.String() + "@enduml\n"
	}

	return result.String()
}
