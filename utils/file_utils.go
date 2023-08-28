package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// FileNameEncode 文件名处理
func FileNameEncode(fileName string) string {
	// 获取当前时间戳
	timeUnix := time.Now().Unix()
	// 脱敏
	// 随机数种子
	rand.Seed(time.Now().UnixNano())
	fileID := strconv.FormatInt(timeUnix, 10) + strconv.Itoa(rand.Intn(1000))
	// 拆分文件名和后缀
	fileSuffix := fileName[strings.LastIndex(fileName, "."):]
	// 拼接新的文件名
	newFileName := fileID + fileSuffix
	return newFileName
}
