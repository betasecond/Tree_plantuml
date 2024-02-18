package cmd

import (
	"Tree_plantuml/internal/clipboard"
	"Tree_plantuml/internal/tree"
	"Tree_plantuml/internal/tree/formats"
	"flag"
	"fmt"
	"os"
)

func main() {
	var format string
	flag.StringVar(&format, "format", "txt", "输出格式：txt, md, puml")
	var path string
	flag.StringVar(&path, "path", ".", "目录路径")
	flag.Parse()

	node, err := tree.GenerateFolderStructure(path)
	if err != nil {
		fmt.Println("生成文件夹结构失败:", err)
		os.Exit(1)
	}

	var output string
	switch format {
	case "txt":
		output = formats.GenerateTXTOutput(node, 0)
	case "md":
		output = formats.GenerateMDOutput(node, 0)
	case "puml":
		output = formats.GeneratePlantUMLContent(node, 0, 20) // 假设levelLimit为20
	default:
		fmt.Println("不支持的格式")
		os.Exit(1)
	}

	fmt.Println(output)
	// 可选：将输出复制到剪贴板
	if err := clipboard.CopyToClipboard(output); err != nil {
		fmt.Println("复制到剪贴板失败:", err)
	} else {
		fmt.Println("内容已复制到剪贴板")
	}
}
