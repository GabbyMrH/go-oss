package logger

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// LogError 当存在错误时记录日志,抛出示例:logger.LogError(errors.New("这是一个错误"))
func LogError(error error) {
	if error != nil {
		filePath := "storage/logs/" + time.Now().Format("2006-01-02") + ".log"
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
		if err != nil {
			fmt.Println("文件打开失败")
		}
		// 延迟关闭文件
		defer file.Close()
		// 写入文件(使用带缓存的*Writer)
		write := bufio.NewWriter(file)
		dateTime := "[" + time.Now().Format("2006-01-02 15:04:05") + "] "
		write.WriteString(dateTime + error.Error() + "\r\n")
		write.Flush()
	}

}
