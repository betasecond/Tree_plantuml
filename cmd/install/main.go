package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"Tree_plantuml/pkg/utils" // 请替换为您的utils包的实际导入路径
)

func main() {
	// 1. 复制tree.exe到指定目录
	sourcePath := "binary/tree/tree.exe"
	destinationPath := filepath.Join(os.Getenv("APPDATA"), "TreeThisFolder", "tree.exe")

	// 创建目标目录
	if err := os.MkdirAll(filepath.Dir(destinationPath), os.ModePerm); err != nil {
		fmt.Printf("创建目录失败: %v\n", err)
		return
	}

	// 复制文件
	if err := copyFile(sourcePath, destinationPath); err != nil {
		fmt.Printf("复制文件失败: %v\n", err)
		return
	}

	// 2. 修改注册表添加右键菜单选项
	regPath := `Software\Classes\*\shell\TreeThisFolder`
	command := fmt.Sprintf(`"%s" "%%1"`, destinationPath) // 使用%%1传递文件路径

	// 添加命令
	if err := utils.ModifyRegistry(regPath, utils.AddAction, "command", command); err != nil {
		fmt.Printf("注册表修改失败: %v\n", err)
		return
	}

	fmt.Println("安装成功！")
}

// copyFile 复制文件从源路径到目标路径
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}
