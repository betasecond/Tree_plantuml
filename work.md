# TreeThisFolder Project Structure

TreeThisFolder is a utility to enhance the Windows context menu, providing options to generate and display the directory structure in various formats directly from the right-click menu. This document outlines the structure and components of the TreeThisFolder project.

## Project Layout

Tree_plantuml/</br>
├── README.md</br>
├── assets</br>
├── cmd</br>
│   ├── build_and_run.ps1</br>
│   ├── install</br>
│   │   └── main.go</br>
│   ├── tree</br>
│   │   └── main.go</br>
│   ├── uninstall</br>
│   │   └── main.go</br>
│   └── config.json</br>
├── go.mod</br>
├── go.sum</br>
├── grep</br>
├── internal</br>
│   ├── clipboard</br>
│   │   └── clipboard.go</br>
│   ├── menu</br>
│   │   └── register.go</br>
│   ├── tree</br>
│   │   ├── formats</br>
│   │   │   ├── md.go</br>
│   │   │   ├── puml.go</br>
│   │   │   └── txt.go</br>
│   │   └── generator.go</br>
│   └── writer</br>
│       └── write.go</br>
├── main.go</br>
├── myapp.exe</br>
├── output</br>
│   ├── go_build_Tree_plantuml_cmd_tree.exe</br>
│   └── output.md</br>
├── pkg</br>
│   ├── config</br>
│   │   └── config.go</br>
│   ├── utils</br>
│   │   └── utils.go</br>
└── work.md</br>


## Key Components

### cmd/main.go

- Entry point for the application.
- Parses command-line arguments.
- Initializes application based on configuration.

### internal/menu/register.go

- Contains logic to add, remove, and update Windows context menu items based on the configuration.

### internal/tree

- Contains the core logic (`generator.go`) and format-specific implementations (`txt.go`, `md.go`, `puml.go`, etc.) for generating folder structure representations.

### internal/clipboard/clipboard.go

- Implements functionality to copy the generated folder structure to the clipboard.

### pkg/utils/utils.go

- Provides common utility functions used across the application.

### pkg/config/config.go

- Defines the structure for the configuration file (`config.json`).
- Implements loading and parsing of configuration to control features like source file path and enabled formats.

## Installation Logic (install/main.go)

- Reads `config.json` to determine the source file path and which features to enable.
- Sets up the Windows context menu items according to the configuration, using logic defined in `internal/menu/register.go`.

## Uninstallation Logic (install/uninstall.go)

- Removes previously added context menu items.
- Cleans up any application-specific files or settings.

## Configuration (pkg/config/config.go)

- The application's behavior, such as the source file path and available features (e.g., supported output formats), is driven by `config.json`.
- This approach allows for easy customization and extension of the application's capabilities.

This structure and component breakdown ensures that TreeThisFolder is modular, configurable, and easy to maintain or extend.
