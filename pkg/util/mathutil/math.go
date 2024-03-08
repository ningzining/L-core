package mathutil

import (
	"fmt"
	"strconv"
	"strings"
)

// BinaryToHex
// 将二进制字符串转为十六进制字符串，并且根据二进制的长度生成十六进制位数,空位补0
func BinaryToHex(binaryStr string) (string, error) {
	hexStrLen := len(binaryStr) / 4
	if len(binaryStr)%4 > 0 {
		hexStrLen++
	}
	// 将二进制字符串转换为整数
	i, err := strconv.ParseInt(binaryStr, 2, 64)
	if err != nil {
		return "", err
	}

	// 将整数转换为十六进制字符串
	return fmt.Sprintf("%0*s", hexStrLen, strings.ToUpper(strconv.FormatInt(i, 16))), nil
}

// HexToBinary
// 将十六进制字符串转为二进制字符串，并且根据十六进制的长度生成二进制位数,空位补0
func HexToBinary(hexStr string) (string, error) {
	binaryStrLen := len(hexStr) * 4

	decimal, err := strconv.ParseInt(hexStr, 16, 0)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%0*b", binaryStrLen, decimal), nil
}
