package main

import (
	"encoding/base64"
	"fmt"
	"github.com/dengpju/higo-utils/utils"
	"time"
)

func main() {
	fmt.Println(utils.Dir("./utils").Scan().Suffix("go").Get())
	result := utils.MapOperation(make(map[string]interface{})).Put("key1", "value1").Put("key2", "value2")
	fmt.Println(result)
	result.ForEach(func(key string, value interface{}) {
		fmt.Println(key, value)
	})
	result.Replace("key1", "k")
	fmt.Println(result.Get("key1"))
	result.Remove("k")
	result.Remove("key2")
	fmt.Println(result.Len())
	fmt.Println(result)
	result.Clear()
	result.ForEach(func(key string, value interface{}) {
		fmt.Println(key, value, "for")
	})
	rsa := utils.NewRsa().SetBits(1024).Build()
	fmt.Println("当前时间戳", utils.CurrentTimestamp())
	rsa.SetExpired(utils.CurrentTimestamp() + 6)
	rsa.SetLimen(10)
	fmt.Println(rsa.Flag())

	// 公钥加密
	e := utils.PubEncrypt(rsa, []byte("123"))
	fmt.Println("公钥加密===")
	fmt.Println(base64.StdEncoding.EncodeToString(e))
	fmt.Println("私钥解密===")
	// 私钥解密
	s := utils.PriDecrypt(rsa, e).String()
	fmt.Println(s)
	fmt.Println("=====")
	time.Sleep(5 * time.Second)
	if i := utils.SecretContainer.Len(); i > 0 {
		fmt.Println("有", i, "个秘钥对")
	}
	//utils.SecretContainer.ForEach(func(key string, value interface{}) {
	//	fmt.Println(key, value.(*utils.Rsa).Expired())
	//	if utils.CurrentTimestamp() >= value.(*utils.Rsa).Expired() {
	//		fmt.Println("秘钥对过期了")
	//		utils.SecretContainer.Remove(key)
	//		fmt.Println("删除过期秘钥对")
	//	}
	//})
	utils.SecretExpiredClear()
	if i := utils.SecretContainer.Len(); i <= 0 {
		fmt.Println("有", i, "个秘钥对")
	}

	// 私钥加密
	s1 := utils.PriEncrypt(rsa, []byte("789")).Base64Encode()
	fmt.Println(s1)
	fmt.Println(s1.Base64Decode())
	fmt.Println(s1.Base64Decode().Base64Encode().Base64Decode())
	// 公钥解密
	ss := utils.PubDecrypt(rsa, s1.Base64Decode()).String()
	fmt.Println(ss)
	fmt.Println(utils.SecretContainer.Exist(rsa.Flag()))
	if utils.SecretContainer.Exist(rsa.Flag()) {
		fmt.Println(utils.SecretContainer.Get(rsa.Flag()).(*utils.Rsa).Flag())
	}
}
