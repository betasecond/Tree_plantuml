package writer

import "os"

// WriteToFile 将内容写入指定文件
func WriteToFile(content, filepath string) error {
	return os.WriteFile(filepath, []byte(content), 0644)
}
