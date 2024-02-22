package main

import (
	"Tree_plantuml/pkg/config" // 请替换为您的config包的实际导入路径
	"Tree_plantuml/pkg/utils"  // 请替换为您的utils包的实际导入路径
	"fmt"
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

	// 根据配置文件设置目标路径
	destinationPath := filepath.Join(os.Getenv("APPDATA"), "TreeThisFolder", "tree.exe")

	// 删除目标文件
	if err := os.Remove(destinationPath); err != nil {
		fmt.Printf("删除文件失败: %v\n", err)
		// 如果文件不存在，则继续执行后续的卸载步骤
		if !os.IsNotExist(err) {
			return
		}
	}

	// 示例：根据配置决定是否移除注册表项
	if cfg.Features.ContextMenu {
		// 移除注册表项
		regPath := `Software\Classes\*\shell\TreeThisFolder`
		if err := utils.ModifyRegistry(regPath, utils.DeleteAction, "", ""); err != nil {
			fmt.Printf("移除注册表项失败: %v\n", err)
			return
		}
	}

	fmt.Println("卸载成功！")
}
