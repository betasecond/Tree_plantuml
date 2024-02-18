package clipboard

import (
	"github.com/atotto/clipboard"
)

// CopyToClipboard 接受一个字符串参数并将其复制到系统剪贴板
func CopyToClipboard(content string) error {
	return clipboard.WriteAll(content)
}
