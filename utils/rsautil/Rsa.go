package rsautil

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/dengpju/higo-utils/utils/encodeutil"
	"github.com/dengpju/higo-utils/utils/maputil"
	"github.com/dengpju/higo-utils/utils/randomutil"
	"github.com/dengpju/higo-utils/utils/timeutil"
	"math/big"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	SecretContainer maputil.IMap
	rsaOnce         sync.Once
	FlagDES         func(rsa *Rsa) string
	Perturb         func(rsa *Rsa) string
)

func init() {
	rsaOnce.Do(func() {
		SecretContainer = maputil.Array()
		FlagDES = func(rsa *Rsa) string {
			return rsa.Flag
		}
		Perturb = func(rsa *Rsa) string {
			str := strings.Replace(string(rsa.GetPubkey()), "-----BEGIN PUBLIC KEY-----\n", "", 1)
			str = strings.Replace(str, "-----END PUBLIC KEY-----\n", "", 1)
			pubs := strings.Split(str, "\n")
			rowLen := len(pubs[0])
			newPubkey := make([]string, 0)
			for _, v := range pubs {
				newPubkey = append(newPubkey, v)
				if len(v) == rowLen {
					newPubkey = append(newPubkey, randomutil.Random().String(rowLen))
				}
			}
			return strings.Join(newPubkey, "\n")
		}
	})
}

type Bytes []byte

func (this Bytes) String() BytesString {
	return BytesString(this)
}

func (this Bytes) Base64Encode() BytesString {
	return BytesString(encodeutil.Base64Encode(this))
}

type BytesString string

func (this BytesString) Base64Decode() Bytes {
	return encodeutil.Base64Decode(string(this))
}

type Rsa struct {
	Pubkey         []byte          `json:"pubkey"`
	Prikey         []byte          `json:"prikey"`
	PubFile        string          `json:"pubfile"`
	PriFile        string          `json:"prifile"`
	Bits           int             `json:"bits"`
	Flag           string          `json:"flag"`
	Expired        int64           `json:"expired"`
	X509PrivateKey []byte          `json:"x_509_private_key"`
	PrivateKey     *rsa.PrivateKey `json:"private_key"`
	PublicKey      *rsa.PublicKey  `json:"public_key"`
	Limen          int             `json:"limen"`
	Counter        int             `json:"counter"`
}

func NewRsa() *Rsa {
	return &Rsa{}
}

func (this *Rsa) GetCounter() int {
	return this.Counter
}

func (this *Rsa) SetCounter(counter int) *Rsa {
	this.Counter = counter
	return this
}

func (this *Rsa) IncCounter(counter int) *Rsa {
	this.Counter += counter
	return this
}

func (this *Rsa) GetLimen() int {
	return this.Limen
}

func (this *Rsa) SetLimen(limen int) *Rsa {
	this.Limen = limen
	return this
}

func (this *Rsa) GetPublicKey() *rsa.PublicKey {
	return this.PublicKey
}

func (this *Rsa) SetPublicKey(publicKey *rsa.PublicKey) *Rsa {
	this.PublicKey = publicKey
	return this
}

func (this *Rsa) GetPrivateKey() *rsa.PrivateKey {
	return this.PrivateKey
}

func (this *Rsa) SetPrivateKey(privateKey *rsa.PrivateKey) *Rsa {
	this.PrivateKey = privateKey
	return this
}

func (this *Rsa) GetX509PrivateKey() []byte {
	return this.X509PrivateKey
}

func (this *Rsa) SetX509PrivateKey(x509PrivateKey []byte) *Rsa {
	this.X509PrivateKey = x509PrivateKey
	return this
}

func (this *Rsa) GetExpired() int64 {
	return this.Expired
}

func (this *Rsa) SetExpired(expired int64) *Rsa {
	this.Expired = expired
	return this
}

func (this *Rsa) GetBits() int {
	return this.Bits
}

func (this *Rsa) SetBits(bits int) *Rsa {
	this.Bits = bits
	return this
}

func (this *Rsa) GetFlag() string {
	return this.Flag
}

func (this *Rsa) SetFlag(flag string) *Rsa {
	this.Flag = flag
	return this
}

func (this *Rsa) GetPriFile() string {
	return this.PriFile
}

func (this *Rsa) SetPriFile(priFile string) *Rsa {
	this.PriFile = priFile
	return this
}

func (this *Rsa) GetPubFile() string {
	return this.PubFile
}

func (this *Rsa) SetPubFile(pubFile string) *Rsa {
	this.PubFile = pubFile
	return this
}

func (this *Rsa) GetPrikey() []byte {
	return this.Prikey
}

func (this *Rsa) SetPrikey(prikey []byte) *Rsa {
	this.Prikey = prikey
	return this
}

func (this *Rsa) GetPubkey() []byte {
	return this.Pubkey
}

func (this *Rsa) SetPubkey(pubkey []byte) *Rsa {
	this.Pubkey = pubkey
	return this
}

func (this *Rsa) Build() *Rsa {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, this.Bits)
	if err != nil {
		panic(err)
	}
	this.PrivateKey = privateKey
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	this.X509PrivateKey = x509.MarshalPKCS1PrivateKey(this.PrivateKey)

	//使用pem格式对x509输出的内容进行编码
	bufferPrivate := new(bytes.Buffer)
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA PRIVATE KEY", Bytes: this.X509PrivateKey}
	// 生成私钥
	err = pem.Encode(bufferPrivate, &privateBlock)
	if err != nil {
		panic("Private Key build fail")
	}
	// 保存字私钥字符串
	this.Prikey = bufferPrivate.Bytes()

	//生成公钥
	//获取公钥的数据
	publicKey := this.PrivateKey.PublicKey
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
	this.Pubkey = bufferPublic.Bytes()

	if "" == this.Flag {
		this.Flag = strconv.FormatInt(time.Now().Unix(), 10)
	}
	this.Flag = FlagDES(this)
	//放入容器
	SecretContainer.Put(this.Flag, this)

	return this
}

func (this *Rsa) Output() {
	//创建文件保存私钥
	privateFile, err := os.Create(this.PriFile) // private.pem
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	_, err = privateFile.WriteString(string(this.Prikey))
	if err != nil {
		panic(err)
	}

	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create(this.PubFile) //public.pem
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	_, err = publicFile.WriteString(string(this.Pubkey))
	if err != nil {
		panic(err)
	}
}

func (this *Rsa) Perturb() string {
	return Perturb(this)
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

// 私钥加密
func PriEncrypt(r *Rsa, data []byte) Bytes {
	//计数
	r.IncCounter(1)
	signData, err := rsa.SignPKCS1v15(nil, r.GetPrivateKey(), crypto.Hash(0), data)
	if err != nil {
		panic(err)
	}
	return signData
}

// 公钥解密
func PubDecrypt(r *Rsa, data []byte) Bytes {
	//计数
	r.IncCounter(1)
	decData, err := publicDecrypt(r.GetPublicKey(), crypto.Hash(0), nil, data)
	if err != nil {
		print(err)
	}
	return decData
}

//公钥加密
func PubEncrypt(r *Rsa, plainText []byte) Bytes {
	//计数
	r.IncCounter(1)
	//pem解码
	block, _ := pem.Decode(r.Pubkey)
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

//私钥解密
func PriDecrypt(r *Rsa, cipherText []byte) Bytes {
	//计数
	r.IncCounter(1)
	//pem解码
	block, _ := pem.Decode(r.GetPrikey())
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

// 密钥过期清除
func SecretExpiredClear() {
	SecretContainer.ForEach(func(key string, value interface{}) {
		r := value.(*Rsa)
		if r.GetExpired() > 0 && timeutil.Time() >= r.GetExpired() {
			SecretContainer.Remove(key)
		} else if r.GetLimen() > 0 && r.GetCounter() >= r.GetLimen() {
			SecretContainer.Remove(key)
		}
	})
}
