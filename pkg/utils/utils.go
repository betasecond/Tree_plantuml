package utils

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
)

type RegAction int

const (
	AddAction RegAction = iota
	DeleteAction
)

// ModifyRegistry 修改注册表项
// ModifyRegistry 修改指定路径下的注册表项。
// 参数:
//
//	path: 注册表项的路径，如"HKEY_CURRENT_USER\Software\MyApp"。
//	action: 操作类型，AddAction 表示添加一个新的键值对，DeleteAction 表示删除键值对。
//	key: 要添加或删除的键的名称。
//	value: 当 action 为 AddAction 时，此参数表示键对应的值。如果 action 是 DeleteAction，此参数将被忽略。
//
// 返回值:
//
//	error: 如果操作成功完成，返回 nil。否则返回一个包含错误详情的 error 对象。
//
// 示例：
//
//	err := ModifyRegistry("Software\\MyApp", AddAction, "NewValueName", "NewValueData")
//	if err != nil {
//	    log.Fatalf("注册表修改失败: %v", err)
//	}
func ModifyRegistry(path string, action RegAction, key string, value string) error {
	// 打开注册表项
	k, err := registry.OpenKey(registry.CURRENT_USER, path, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer k.Close()

	switch action {
	case AddAction:
		// 设置键值
		if err := k.SetStringValue(key, value); err != nil {
			return err
		}
	case DeleteAction:
		// 删除键值
		if err := k.DeleteValue(key); err != nil {
			return err
		}
	default:
		return fmt.Errorf("未知的操作")
	}

	return nil
}
