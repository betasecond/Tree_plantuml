package main

import (
	"Tree_plantuml/internal/clipboard"
	"Tree_plantuml/internal/tree"
	"Tree_plantuml/internal/tree/formats"
	"Tree_plantuml/internal/writer"
	"flag"
	"fmt"
	"os"
)

// OutputOption 定义输出选项类型
type OutputOption int

// 使用iota定义输出选项
const (
	None      OutputOption = iota // 默认值，不执行任何操作
	Clipboard                     // 复制到剪贴板
	File                          // 生成文件
	Both                          // 复制到剪贴板并生成文件
)

func main() {
	var format, path string
	var option int
	flag.StringVar(&format, "format", "txt", "输出格式：txt, md, puml")
	flag.StringVar(&path, "path", ".", "目录路径")
	flag.IntVar(&option, "option", 0, "选择输出选项：0=无，1=复制到剪贴板，2=保存到文件，3=两者都执行")
	flag.Parse()

	node, err := tree.GenerateFolderStructure(path)
	if err != nil {
		fmt.Println("生成文件夹结构失败:", err)
		os.Exit(1)
	}

	output := ""
	switch format {
	case "txt":
		output = formats.GenerateTXTOutput(node, 0)
	case "md":
		output = formats.GenerateMDOutput(node, 0)
	case "puml":
		output = formats.GeneratePlantUMLContent(node, 0, 20)
	default:
		fmt.Println("不支持的格式")
		os.Exit(1)
	}

	switch OutputOption(option) {
	case Clipboard:
		if err := clipboard.CopyToClipboard(output); err != nil {
			fmt.Println("复制到剪贴板失败:", err)
		} else {
			fmt.Println("内容已复制到剪贴板")
		}
	case File:
		outputPath := "output." + format
		if err := writer.WriteToFile(output, outputPath); err != nil {
			fmt.Println("写入文件失败:", err)
		} else {
			fmt.Printf("文件结构已保存到 %s\n", outputPath)
		}
	case Both:
		if err := clipboard.CopyToClipboard(output); err != nil {
			fmt.Println("复制到剪贴板失败:", err)
		} else {
			fmt.Println("内容已复制到剪贴板")
		}
		outputPath := "output." + format
		if err := writer.WriteToFile(output, outputPath); err != nil {
			fmt.Println("写入文件失败:", err)
		} else {
			fmt.Printf("文件结构已保存到 %s\n", outputPath)
		}
	default:
		fmt.Println(output)
	}
}
