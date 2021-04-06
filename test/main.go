package main

import (
	"encoding/base64"
	"fmt"
	"github.com/dengpju/higo-utils/utils"
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
	fmt.Println(rsa.Flag())
	e := utils.PubEncrypt(rsa, []byte("123"))
	fmt.Println("公钥加密===")
	fmt.Println(base64.StdEncoding.EncodeToString(e))
	fmt.Println("私钥解密===")
	s := utils.PriDecrypt(rsa, e).String()
	fmt.Println(s)
	fmt.Println("=====")
	s1 := utils.PriEncrypt(rsa.PrivateKey(), []byte("789")).Base64Encode()
	fmt.Println(s1)
	fmt.Println(s1.Base64Decode())
	fmt.Println(s1.Base64Decode().Base64Encode().Base64Decode())
	ss := utils.PubDecrypt(rsa.PublicKey(), s1.Base64Decode()).String()
	fmt.Println(ss)
	fmt.Println((utils.RsaMap.Get(rsa.Flag()).(*utils.Rsa).Flag()))
}
