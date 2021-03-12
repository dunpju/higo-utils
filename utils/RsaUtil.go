package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"
)

type Rsa struct {
	public         []byte
	private        []byte
	pubFile        string
	priFile        string
	bits           int
	flag           string
	expired        time.Duration
	x509PrivateKey []byte
	privateKey     *rsa.PrivateKey
}

func (this *Rsa) PrivateKey() *rsa.PrivateKey {
	return this.privateKey
}

func (this *Rsa) SetPrivateKey(privateKey *rsa.PrivateKey) *Rsa {
	this.privateKey = privateKey
	return this
}

func (this *Rsa) X509PrivateKey() []byte {
	return this.x509PrivateKey
}

func (this *Rsa) SetX509PrivateKey(x509PrivateKey []byte) *Rsa {
	this.x509PrivateKey = x509PrivateKey
	return this
}

func (this *Rsa) Expired() time.Duration {
	return this.expired
}

func (this *Rsa) SetExpired(expired time.Duration) *Rsa {
	this.expired = expired
	return this
}

func (this *Rsa) Bits() int {
	return this.bits
}

func (this *Rsa) SetBits(bits int) *Rsa {
	this.bits = bits
	return this
}

func (this *Rsa) Flag() string {
	return this.flag
}

func NewRsa() *Rsa {
	return &Rsa{}
}

func (this *Rsa) PriFile() string {
	return this.priFile
}

func (this *Rsa) SetPriFile(priFile string) *Rsa {
	this.priFile = priFile
	return this
}

func (this *Rsa) PubFile() string {
	return this.pubFile
}

func (this *Rsa) SetPubFile(pubFile string) *Rsa {
	this.pubFile = pubFile
	return this
}

func (this *Rsa) Private() []byte {
	return this.private
}

func (this *Rsa) SetPrivate(private []byte) *Rsa {
	this.private = private
	return this
}

func (this *Rsa) Public() []byte {
	return this.public
}

func (this *Rsa) SetPublic(public []byte) *Rsa {
	this.public = public
	return this
}

func (this *Rsa) Generate() *Rsa {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, this.bits)
	if err != nil {
		panic(err)
	}
	this.privateKey = privateKey
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	this.x509PrivateKey = x509.MarshalPKCS1PrivateKey(this.privateKey)

	//使用pem格式对x509输出的内容进行编码
	bufferPrivate := new(bytes.Buffer)
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: this.x509PrivateKey}
	// 生成私钥
	err = pem.Encode(bufferPrivate, &privateBlock)
	if err != nil {
		panic("Private Key build fail")
	}
	// 保存字私钥字符串
	this.private = bufferPrivate.Bytes()

	//生成公钥
	//获取公钥的数据
	publicKey := this.privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}

	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	// 提取公钥
	bufferPublic := new(bytes.Buffer)
	err = pem.Encode(bufferPublic, &publicBlock)
	if err != nil {
		panic("Public Key extract fail")
	}
	this.public = bufferPublic.Bytes()

	return this
}

func (this *Rsa) Output() {
	//创建文件保存私钥
	privateFile, err := os.Create(this.priFile) // private.pem
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	privateFile.WriteString(string(this.private))

	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create(this.pubFile) //public.pem
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	publicFile.WriteString(string(this.public))
}

//RSA私钥加密
func RsaPriEncrypt(plainText []byte, r *Rsa) []byte {
	//pem解码
	block, _ := pem.Decode(r.private)
	fmt.Println(block)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}
	//返回密文
	return cipherText
}

//RSA公钥解密
func RsaPubDecrypt(cipherText []byte, r *Rsa) []byte {
	//pem解码
	block, _ := pem.Decode(r.public)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	//返回明文
	return plainText
}
