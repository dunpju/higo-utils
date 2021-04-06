package utils

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	RsaMap  IMap
	rsaOnce sync.Once
)

type Rsa struct {
	pubkey         []byte
	prikey         []byte
	pubFile        string
	priFile        string
	bits           int
	flag           string
	expired        time.Duration
	x509PrivateKey []byte
	privateKey     *rsa.PrivateKey
	publicKey      *rsa.PublicKey
}

func init() {
	rsaOnce.Do(func() {
		RsaMap = make(MapString)
	})
}

func (this *Rsa) PublicKey() *rsa.PublicKey {
	return this.publicKey
}

func (this *Rsa) SetPublicKey(publicKey *rsa.PublicKey) {
	this.publicKey = publicKey
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

func (this *Rsa) SetFlag(flag string) {
	this.flag = flag
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

func (this *Rsa) Prikey() []byte {
	return this.prikey
}

func (this *Rsa) SetPrivate(prikey []byte) *Rsa {
	this.prikey = prikey
	return this
}

func (this *Rsa) Pubkey() []byte {
	return this.pubkey
}

func (this *Rsa) SetPublic(pubkey []byte) *Rsa {
	this.pubkey = pubkey
	return this
}

func (this *Rsa) Build() *Rsa {
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
	privateBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: this.x509PrivateKey}
	// 生成私钥
	err = pem.Encode(bufferPrivate, &privateBlock)
	if err != nil {
		panic("Private Key build fail")
	}
	// 保存字私钥字符串
	this.prikey = bufferPrivate.Bytes()

	//生成公钥
	//获取公钥的数据
	publicKey := this.privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}

	this.SetPublicKey(&publicKey)

	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "PUBLIC KEY", Bytes: X509PublicKey}
	// 提取公钥
	bufferPublic := new(bytes.Buffer)
	err = pem.Encode(bufferPublic, &publicBlock)
	if err != nil {
		panic("Public Key extract fail")
	}
	this.pubkey = bufferPublic.Bytes()

	if "" == this.flag {
		this.flag = strconv.FormatInt(time.Now().Unix(), 10)
	}

	//放入容器
	RsaMap.Put(this.flag, this)

	return this
}

func (this *Rsa) Output() {
	//创建文件保存私钥
	privateFile, err := os.Create(this.priFile) // private.pem
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	_, err = privateFile.WriteString(string(this.prikey))
	if err != nil {
		panic(err)
	}

	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create(this.pubFile) //public.pem
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	_, err = publicFile.WriteString(string(this.pubkey))
	if err != nil {
		panic(err)
	}
}

//RSA公钥加密
func PubEncrypt(plainText []byte, r *Rsa) []byte {
	//pem解码
	block, _ := pem.Decode(r.pubkey)
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

//RSA私钥解密
func PriDecrypt(cipherText []byte, r *Rsa) []byte {
	//pem解码
	block, _ := pem.Decode(r.Prikey())
	//X509解码
	privateKey, er := x509.ParsePKCS1PrivateKey(block.Bytes)
	if er != nil {
		panic(er)
	}
	//对密文进行解密
	plainText, e := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if e != nil {
		panic(e)
	}
	//返回明文
	return plainText
}

// copy from crypt/rsa/pkcs1v5.go
var hashPrefixes = map[crypto.Hash][]byte{
	crypto.MD5:       {0x30, 0x20, 0x30, 0x0c, 0x06, 0x08, 0x2a, 0x86, 0x48, 0x86, 0xf7, 0x0d, 0x02, 0x05, 0x05, 0x00, 0x04, 0x10},
	crypto.SHA1:      {0x30, 0x21, 0x30, 0x09, 0x06, 0x05, 0x2b, 0x0e, 0x03, 0x02, 0x1a, 0x05, 0x00, 0x04, 0x14},
	crypto.SHA224:    {0x30, 0x2d, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x04, 0x05, 0x00, 0x04, 0x1c},
	crypto.SHA256:    {0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20},
	crypto.SHA384:    {0x30, 0x41, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x02, 0x05, 0x00, 0x04, 0x30},
	crypto.SHA512:    {0x30, 0x51, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x03, 0x05, 0x00, 0x04, 0x40},
	crypto.MD5SHA1:   {}, // A special TLS case which doesn't use an ASN1 prefix.
	crypto.RIPEMD160: {0x30, 0x20, 0x30, 0x08, 0x06, 0x06, 0x28, 0xcf, 0x06, 0x03, 0x00, 0x31, 0x04, 0x14},
}

// copy from crypt/rsa/pkcs1v5.go
func encrypt(c *big.Int, pub *rsa.PublicKey, m *big.Int) *big.Int {
	e := big.NewInt(int64(pub.E))
	c.Exp(m, e, pub.N)
	return c
}

// copy from crypt/rsa/pkcs1v5.go
func pkcs1v15HashInfo(hash crypto.Hash, inLen int) (hashLen int, prefix []byte, err error) {
	// Special case: crypto.Hash(0) is used to indicate that the data is
	// signed directly.
	if hash == 0 {
		return inLen, nil, nil
	}

	hashLen = hash.Size()
	if inLen != hashLen {
		return 0, nil, errors.New("crypto/rsa: input must be hashed message")
	}
	prefix, ok := hashPrefixes[hash]
	if !ok {
		return 0, nil, errors.New("crypto/rsa: unsupported hash function")
	}
	return
}

// copy from crypt/rsa/pkcs1v5.go
func leftPad(input []byte, size int) (out []byte) {
	n := len(input)
	if n > size {
		n = size
	}
	out = make([]byte, size)
	copy(out[len(out)-n:], input)
	return
}
func unLeftPad(input []byte) (out []byte) {
	n := len(input)
	t := 2
	for i := 2; i < n; i++ {
		if input[i] == 0xff {
			t = t + 1
		} else {
			if input[i] == input[0] {
				t = t + int(input[1])
			}
			break
		}
	}
	out = make([]byte, n-t)
	copy(out, input[t:])
	return
}

// copy&modified from crypt/rsa/pkcs1v5.go
func publicDecrypt(pub *rsa.PublicKey, hash crypto.Hash, hashed []byte, sig []byte) (out []byte, err error) {
	hashLen, prefix, err := pkcs1v15HashInfo(hash, len(hashed))
	if err != nil {
		return nil, err
	}

	tLen := len(prefix) + hashLen
	k := (pub.N.BitLen() + 7) / 8
	if k < tLen+11 {
		return nil, fmt.Errorf("length illegal")
	}

	c := new(big.Int).SetBytes(sig)
	m := encrypt(new(big.Int), pub, c)
	em := leftPad(m.Bytes(), k)
	out = unLeftPad(em)

	err = nil
	return
}

func PriEncrypt(privt *rsa.PrivateKey, data []byte) ([]byte, error) {
	signData, err := rsa.SignPKCS1v15(nil, privt, crypto.Hash(0), data)
	if err != nil {
		return nil, err
	}
	return signData, nil
}
func PubDecrypt(pub *rsa.PublicKey, data []byte) ([]byte, error) {
	decData, err := publicDecrypt(pub, crypto.Hash(0), nil, data)
	if err != nil {
		return nil, err
	}
	return decData, nil
}
