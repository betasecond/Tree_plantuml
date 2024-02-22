package main

import (
	"Tree_plantuml/pkg/config"
	"Tree_plantuml/pkg/utils"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("cmd/config.json") // 指定配置文件的路径
	if err != nil {
		fmt.Printf("加载配置失败: %v\n", err)
		return
	}

	// 根据配置文件设置源路径和目标路径
	sourcePath := cfg.SourcePath
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

	// 示例：根据配置决定是否添加注册表项
	if cfg.Features.ContextMenu {
		// 修改注册表添加右键菜单选项
		regPath := `Software\Classes\*\shell\TreeThisFolder`
		command := fmt.Sprintf(`"%s" "%%1"`, destinationPath) // 使用%%1传递文件路径

		// 添加命令
		if err := utils.ModifyRegistry(regPath, utils.AddAction, "command", command); err != nil {
			fmt.Printf("注册表修改失败: %v\n", err)
			return
		}
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
