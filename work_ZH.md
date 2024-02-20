[EN](work_EN.md) | [中文](work_ZH.md)

# TreeThisFolder 项目结构

TreeThisFolder 是一个增强 Windows 上下文菜单（右键菜单）的工具，提供选项以各种格式直接生成和显示目录结构。本文档概述了 TreeThisFolder 项目的结构和组件。

## 项目布局

Tree_plantuml/</br>
├── README.md</br>
├── assets</br>
├── cmd</br>
│ ├── build_and_run.ps1</br>
│ ├── install</br>
│ │ └── main.go</br>
│ ├── tree</br>
│ │ └── main.go</br>
│ ├── uninstall</br>
│ │ └── main.go</br>
│ └── config.json</br>
├── go.mod</br>
├── go.sum</br>
├── grep</br>
├── internal</br>
│ ├── clipboard</br>
│ │ └── clipboard.go</br>
│ ├── menu</br>
│ │ └── register.go</br>
│ ├── tree</br>
│ │ ├── formats</br>
│ │ │ ├── md.go</br>
│ │ │ ├── puml.go</br>
│ │ │ └── txt.go</br>
│ │ └── generator.go</br>
│ └── writer</br>
│ └── write.go</br>
├── main.go</br>
├── myapp.exe</br>
├── output</br>
│ ├── go_build_Tree_plantuml_cmd_tree.exe</br>
│ └── output.md</br>
├── pkg</br>
│ ├── config</br>
│ │ └── config.go</br>
│ ├── utils</br>
│ │ └── utils.go</br>
└── work.md</br>

## 关键组件

### cmd/main.go

- 解析命令行参数。
- 基于配置初始化应用程序。

### internal/menu/register.go

- 包含基于配置添加、移除和更新 Windows 上下文菜单项的逻辑。

### internal/tree
- 包含生成文件夹结构表示的核心逻辑（generator.go）和特定格式实现（txt.go、md.go、puml.go 等）。

### internal/clipboard/clipboard.go

- 实现将生成的文件夹结构复制到剪贴板的功能。

### pkg/utils/utils.go

- 提供跨应用程序使用的常见工具函数。

### pkg/config/config.go

- 定义配置文件（config.json）的结构。
- 实现加载和解析配置以控制功能，如源文件路径和启用的格式。

## 安装逻辑（install/main.go）

- 读取 config.json 以确定源文件路径和启用哪些功能。
- 根据配置使用 internal/menu/register.go 中定义的逻辑设置 Windows 上下文菜单项。

## 卸载逻辑（install/uninstall.go）

- 移除先前添加的上下文菜单项。
- 清理任何特定于应用程序的文件或设置。

## 配置（pkg/config/config.go）

- 应用程序的行为，如源文件路径和可用功能（例如，支持的输出格式），由 config.json 驱动。
- 这种方法允许轻松自定义和扩展应用程序的功能。
- 这种结构和组件划分确保 TreeThisFolder 是模块化的、可配置的，并且易于维护或扩展。


