package tlsutils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/dengpju/higo-utils/utils"
	"github.com/dengpju/higo-utils/utils/fileutils"
	"math/big"
	"net"
	"os"
	"time"
)

var (
	CrtFileName        = "higo.crt.pem"
	KeyFileName        = "higo.key.pem"
	Organization       = "WWW.YUMI.COM"
	OrganizationalUnit = "ITs"
	CommonName         = "YUMI.COM Web"
	IP                 = "127.0.0.1"
)

type TLS struct {
	bits   int
	outDir string
	crt    string
	key    string
}

func (this *TLS) Bits() int {
	return this.bits
}

func (this *TLS) SetBits(bits int) *TLS {
	this.bits = bits
	return this
}

// 构造函数
func NewTLS(outDir string, crt string, key string) *TLS {
	return &TLS{bits: 1024, outDir: outDir, crt: crt, key: key}
}

// 生成ssl证书
func (this *TLS) Build() {

	// 创建输出目录
	this.createOutDir()

	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)

	// 定义：引用IETF的安全领域的公钥基础实施（PKIX）工作组的标准实例化内容
	subject := pkix.Name{
		Organization:       []string{Organization},
		OrganizationalUnit: []string{OrganizationalUnit},
		CommonName:         CommonName,
	}

	// 设置 SSL证书的属性用途
	certificate509 := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(100 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:  []net.IP{net.ParseIP(IP)},
	}

	// 生成指定位数密匙
	pk, err := rsa.GenerateKey(rand.Reader, this.bits)
	if err != nil {
		panic(err)
	}

	// 生成 SSL公匙
	derBytes, err := x509.CreateCertificate(rand.Reader, &certificate509, &certificate509, &pk.PublicKey, pk)
	if err != nil {
		panic(err)
	}
	certOut, err := os.Create(this.crt)
	if err != nil {
		panic(err)
	}
	err = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	if err != nil {
		panic(err)
	}
	err = certOut.Close()
	if err != nil {
		panic(err)
	}

	// 生成 SSL私匙
	keyOut, err := os.Create(this.key)
	if err != nil {
		panic(err)
	}
	err = pem.Encode(keyOut, &pem.Block{Type: "RAS PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	if err != nil {
		panic(err)
	}
	err = keyOut.Close()
	if err != nil {
		panic(err)
	}
}

// 更新ssl证书
func (this *TLS) Update() {
	// 创建输出目录
	this.createOutDir()
	// 判断创建时间
	crtFile := fileutils.NewFile(this.crt, os.O_APPEND, os.ModePerm)
	if crtFile.Exist() {
		createTimestamp := crtFile.CreateTimestamp()
		if (utils.Time() - createTimestamp) > utils.Random().IntHour24ToSecond() {
			this.Build() // 重新生成证书
		}
	}
	fmt.Println("更新ssl证书")
}

// 创建输出目录
func (this *TLS) createOutDir() {
	// 目录不存在，并创建
	if _, err := os.Stat(this.outDir); os.IsNotExist(err) {
		if os.Mkdir(this.outDir, os.ModePerm) != nil {
		}
	}

	this.crt = this.outDir + CrtFileName
	this.key = this.outDir + KeyFileName
}
