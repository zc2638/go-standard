package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

const (
	CertFile = "testdata/x509_cert.pem"
	KeyFile  = "testdata/x509_key.pem"
)

// 解析X.509编码的证书和密钥
func main() {

	// x509.MarshalPKCS1PrivateKey 参考rsa示例中generateRSAKey方法
	// x509.MarshalPKCS8PrivateKey 参考rsa示例中generateRSAKey方法

	// x509.MarshalPKIXPublicKey   参考rsa示例中generateRSAKey方法
	// x509.MarshalPKCS1PublicKey  参考rsa示例中generateRSAKey方法

	// x509.ParsePKCS1PrivateKey   参考rsa示例中BuildRSAPKCS1PrivateKey方法
	// x509.ParsePKCS8PrivateKey   参考rsa示例中BuildRSAPrivateKey方法

	// x509.ParsePKCS1PublicKey    参考rsa示例中BuildRSAPKCS1PublicKey方法
	// x509.ParsePKIXPublicKey     参考rsa示例中BuildRSAPublicKey方法

	const RootPem = ``
	// 创建一个新的 证书集合/证书池
	roots := x509.NewCertPool()
	// 试图解析一系列PEM编码的证书。它将找到的任何证书都添加到证书池中，如果所有证书都成功被解析，会返回true
	roots.AppendCertsFromPEM([]byte(RootPem))
	// 添加一个证书
	roots.AddCert(new(x509.Certificate))
	// 返回池中所有证书的DER编码的持有者的列表
	roots.Subjects()

	// 生成证书及密钥
	createCertificate()

	// 通过root证书认证客户端证书
	Verify()
}

func createCertificate() {

	//把 1 左移 128 位，返回给 big.Int
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	//返回在 [0, max) 区间均匀随机分布的一个随机值
	serialNumber, _ := rand.Int(rand.Reader, max)

	// pkix.Name代表一个X.509识别名。只包含识别名的公共属性，额外的属性被忽略
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	// 初始化一个X.509证书
	template := x509.Certificate{
		SerialNumber: serialNumber, // SerialNumber 是 CA 颁布的唯一序列号，在此使用一个大随机数来代表它
		Subject:      subject,
		NotBefore:    time.Now(),                                                   // 有效期起始时间
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),                         // 有效期截止时间
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, // KeyUsage 与 ExtKeyUsage 用来表明该证书是用来做服务器认证的。代表给定密钥的合法操作集
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},               // 密钥扩展用途的序列
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
	}

	//生成一对具有指定字位数的RSA密钥
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	// 基于模板创建一个新的证书。会用到模板的如下字段：SerialNumber、Subject、NotBefore、NotAfter、KeyUsage、ExtKeyUsage、UnknownExtKeyUsage、BasicConstraintsValid、IsCA、MaxPathLen、SubjectKeyId、DNSNames、PermittedDNSDomainsCritical、PermittedDNSDomains、SignatureAlgorithm。
	//该证书会使用parent签名。如果parent和template相同，则证书是自签名的。Pub参数是被签名者的公钥，而priv是签名者的私钥。
	//只支持RSA和ECDSA类型的密钥
	//返回的切片是DER编码的证书。
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)

	// 创建cert.pem文件，如果文件存在内容重置为空
	certOut, err := os.Create(CertFile)
	defer certOut.Close()
	if err != nil {
		log.Fatal(err)
	}
	// 初始化一个PEM编码的结构
	certBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: derBytes,
	}
	// 将Block的pem编码写入文件
	if err := pem.Encode(certOut, certBlock); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Certificate生成成功")

	// 创建key.pem文件，如果文件存在内容重置为空
	keyOut, _ := os.Create(KeyFile)
	defer keyOut.Close()
	// 初始化一个PEM编码的结构
	keyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(pk),
	}
	// 将Block的pem编码写入文件
	if err := pem.Encode(keyOut, keyBlock); err != nil {
		log.Fatal()
	}
	fmt.Println("Key生成成功")
}

func Verify() {

	const rootPEM = `
-----BEGIN CERTIFICATE-----
MIIEBDCCAuygAwIBAgIDAjppMA0GCSqGSIb3DQEBBQUAMEIxCzAJBgNVBAYTAlVT
MRYwFAYDVQQKEw1HZW9UcnVzdCBJbmMuMRswGQYDVQQDExJHZW9UcnVzdCBHbG9i
YWwgQ0EwHhcNMTMwNDA1MTUxNTU1WhcNMTUwNDA0MTUxNTU1WjBJMQswCQYDVQQG
EwJVUzETMBEGA1UEChMKR29vZ2xlIEluYzElMCMGA1UEAxMcR29vZ2xlIEludGVy
bmV0IEF1dGhvcml0eSBHMjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AJwqBHdc2FCROgajguDYUEi8iT/xGXAaiEZ+4I/F8YnOIe5a/mENtzJEiaB0C1NP
VaTOgmKV7utZX8bhBYASxF6UP7xbSDj0U/ck5vuR6RXEz/RTDfRK/J9U3n2+oGtv
h8DQUB8oMANA2ghzUWx//zo8pzcGjr1LEQTrfSTe5vn8MXH7lNVg8y5Kr0LSy+rE
ahqyzFPdFUuLH8gZYR/Nnag+YyuENWllhMgZxUYi+FOVvuOAShDGKuy6lyARxzmZ
EASg8GF6lSWMTlJ14rbtCMoU/M4iarNOz0YDl5cDfsCx3nuvRTPPuj5xt970JSXC
DTWJnZ37DhF5iR43xa+OcmkCAwEAAaOB+zCB+DAfBgNVHSMEGDAWgBTAephojYn7
qwVkDBF9qn1luMrMTjAdBgNVHQ4EFgQUSt0GFhu89mi1dvWBtrtiGrpagS8wEgYD
VR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAQYwOgYDVR0fBDMwMTAvoC2g
K4YpaHR0cDovL2NybC5nZW90cnVzdC5jb20vY3Jscy9ndGdsb2JhbC5jcmwwPQYI
KwYBBQUHAQEEMTAvMC0GCCsGAQUFBzABhiFodHRwOi8vZ3RnbG9iYWwtb2NzcC5n
ZW90cnVzdC5jb20wFwYDVR0gBBAwDjAMBgorBgEEAdZ5AgUBMA0GCSqGSIb3DQEB
BQUAA4IBAQA21waAESetKhSbOHezI6B1WLuxfoNCunLaHtiONgaX4PCVOzf9G0JY
/iLIa704XtE7JW4S615ndkZAkNoUyHgN7ZVm2o6Gb4ChulYylYbc3GrKBIxbf/a/
zG+FA1jDaFETzf3I93k9mTXwVqO94FntT0QJo544evZG0R0SnU++0ED8Vf4GXjza
HFa9llF7b1cq26KqltyMdMKVvvBulRP/F/A8rLIQjcxz++iPAsbw+zOzlTvjwsto
WHPbqCRiOwY1nQ2pM714A5AuTHhdUDqB1O6gyHA43LL5Z/qHQF1hwFGPa4NrzQU6
yuGnBXj8ytqU0CwIPX4WecigUCAkVDNx
-----END CERTIFICATE-----`

	const certPEM = `
-----BEGIN CERTIFICATE-----
MIIDujCCAqKgAwIBAgIIE31FZVaPXTUwDQYJKoZIhvcNAQEFBQAwSTELMAkGA1UE
BhMCVVMxEzARBgNVBAoTCkdvb2dsZSBJbmMxJTAjBgNVBAMTHEdvb2dsZSBJbnRl
cm5ldCBBdXRob3JpdHkgRzIwHhcNMTQwMTI5MTMyNzQzWhcNMTQwNTI5MDAwMDAw
WjBpMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwN
TW91bnRhaW4gVmlldzETMBEGA1UECgwKR29vZ2xlIEluYzEYMBYGA1UEAwwPbWFp
bC5nb29nbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfRrObuSW5T7q
5CnSEqefEmtH4CCv6+5EckuriNr1CjfVvqzwfAhopXkLrq45EQm8vkmf7W96XJhC
7ZM0dYi1/qOCAU8wggFLMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAa
BgNVHREEEzARgg9tYWlsLmdvb2dsZS5jb20wCwYDVR0PBAQDAgeAMGgGCCsGAQUF
BwEBBFwwWjArBggrBgEFBQcwAoYfaHR0cDovL3BraS5nb29nbGUuY29tL0dJQUcy
LmNydDArBggrBgEFBQcwAYYfaHR0cDovL2NsaWVudHMxLmdvb2dsZS5jb20vb2Nz
cDAdBgNVHQ4EFgQUiJxtimAuTfwb+aUtBn5UYKreKvMwDAYDVR0TAQH/BAIwADAf
BgNVHSMEGDAWgBRK3QYWG7z2aLV29YG2u2IaulqBLzAXBgNVHSAEEDAOMAwGCisG
AQQB1nkCBQEwMAYDVR0fBCkwJzAloCOgIYYfaHR0cDovL3BraS5nb29nbGUuY29t
L0dJQUcyLmNybDANBgkqhkiG9w0BAQUFAAOCAQEAH6RYHxHdcGpMpFE3oxDoFnP+
gtuBCHan2yE2GRbJ2Cw8Lw0MmuKqHlf9RSeYfd3BXeKkj1qO6TVKwCh+0HdZk283
TZZyzmEOyclm3UGFYe82P/iDFt+CeQ3NpmBg+GoaVCuWAARJN/KfglbLyyYygcQq
0SgeDh8dRKUiaW3HQSoYvTvdTuqzwK4CXsr3b5/dAOY8uMuG/IAR3FgwTbZ1dtoW
RvOTa8hYiU6A475WuZKyEHcwnGYe57u2I2KbMgcKjPniocj4QzgYsVAVKW3IwaOh
yE+vPxsiUkvQHdO2fojCkY8jg70jxM+gu59tPDNbw3Uh/2Ij310FgTHsnGQMyA==
-----END CERTIFICATE-----`

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}

	block, _ := pem.Decode([]byte(certPEM))
	if block == nil {
		panic("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("failed to parse certificate: " + err.Error())
	}

	opts := x509.VerifyOptions{
		DNSName: "mail.google.com",
		Roots:   roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		panic("failed to verify certificate: " + err.Error())
	}
}
