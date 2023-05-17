package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strings"
)

// 获取本地网卡、IP信息
func GetLocalMac() (IPAddrs string, MacAddrs string) {

	IPAddrs = ""
	MacAddrs = ""
	//全部网卡接口
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Error:" + err.Error())
	}
	//接口一个个处理
	for _, inter := range interfaces {
		//IP地址
		addrs, _ := inter.Addrs()
		//只取有效的IP地址
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					IPAddrs += ipnet.IP.String()
					IPAddrs += ","
				}
			}
		}

		//网卡MAC地址
		if inter.HardwareAddr != nil {
			MacAddrs += inter.HardwareAddr.String()
			MacAddrs += ","
		}
	}

	IPAddrs = strings.TrimRight(IPAddrs, ",")
	MacAddrs = strings.TrimRight(MacAddrs, ",")
	return IPAddrs, MacAddrs
}

// 使用PKCS7进行填充，IOS也是7
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// aes加密，填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
func AesCBCEncrypt(rawData []byte) ([]byte, error) {
	var key []byte = []byte{'9', '5', '2', '7', 'l', 'R', '*', ';', '@', 'R', ':', '2', 'G', 'F', 'F', '1'}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	//填充原文
	blockSize := block.BlockSize()
	rawData = PKCS7Padding(rawData, blockSize)
	//初始向量IV必须是唯一，但不需要保密
	cipherText := make([]byte, len(rawData))
	//block大小 16
	var iv []byte = []byte{'g', 'f', 'd', 'e', 'r', 't', 'f', 'g', 'h', 'j', 'k', 'u', 'y', 'r', 't', 'g'}

	//block大小和初始向量大小一定要一致
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText, rawData)

	return cipherText, nil
}

// int 转小端 4字节 []byte
func IntToBytesLittleEndian(n int) ([]byte, error) {

	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, &tmp)
	return bytesBuffer.Bytes(), nil

}

func GetMachineInfo() string {
	//取主机名
	host, err := os.Hostname()
	if err != nil {
		host = "NULL"
	}

	IPAddrs, MacAddrs := GetLocalMac()
	msgInfo := fmt.Sprintf("%v|%v|%v", host, IPAddrs, MacAddrs)

	cipherText, _ := AesCBCEncrypt([]byte(msgInfo))
	msgInfo = base64.StdEncoding.EncodeToString(cipherText)
	outMsg := fmt.Sprintf("{\"info\":\"%v\"}", msgInfo)

	return outMsg
}

func main() {
	outMsg := GetMachineInfo()
	msgLen := len(outMsg)
	outlen, _ := IntToBytesLittleEndian(msgLen)
	os.Stdout.Write(outlen)
	os.Stdout.WriteString(outMsg + "\n")
}
