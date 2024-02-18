TreeThisFolder/
├── cmd
│   └── main.go          # 程序入口，处理命令行参数和初始化
├── internal
│   ├── menu             # 右键菜单集成逻辑
│   │   └── register.go  # 注册和管理右键菜单项
│   ├── tree             # 文件夹结构生成逻辑
│   │   ├── generator.go # 生成文件夹结构的核心逻辑
│   │   └── formats      # 支持的文件格式
│   │       ├── txt.go
│   │       ├── md.go
│   │       ├── puml.go
│   │       ├── win_tree.go
│   │       └── emoji_tree.go
│   └── clipboard        # 剪贴板操作
│       └── clipboard.go # 复制内容到剪贴板
└── pkg
├── utils            # 通用工具函数
│   └── utils.go
└── config           # 配置相关
└── config.go
