package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net"
	"time"
)

// 实现了TLS 1.2安全协议
func main() {

	var conn net.Conn

	// 创建一个新的 证书集合/证书池
	roots := x509.NewCertPool()
	// 试图解析一系列PEM编码的证书。它将找到的任何证书都添加到证书池中，如果所有证书都成功被解析，会返回true
	ok := roots.AppendCertsFromPEM([]byte(RootPEM))
	if !ok {
		panic("failed to parse root certificate")
	}

	// 读取并解析一对文件获取公钥和私钥，返回证书链。这些文件必须是PEM编码
	cert, err := tls.LoadX509KeyPair("testdata/tls-example-cert.pem", "testdata/tls-example-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	// 解析一对PEM编码的数据获取公钥和私钥
	cert2, err := tls.X509KeyPair([]byte(CertPem), []byte(KeyPem))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cert2)

	// TLS客户端或服务端配置
	// 在本类型的值提供给TLS函数后，就不应再修改该值。Config类型值可能被重用；tls包也不会修改它
	var cfg = &tls.Config{
		Certificates:             []tls.Certificate{cert}, // 不少于一个证书的链，用于提供给连接的另一端，服务端必须保证至少有一个证书
		Rand:                     nil,                     // Rand提供用于生成随机数和RSA盲签名的熵源，该接口必须能安全的用于并发。如果Rand是nil，会使用crypto/rand包的密码用随机数读取器。
		Time:                     nil,                     // Time返回当前时间，如果是nil会使用time.Now
		NameToCertificate:        nil,                     // // 映射证书名到证书。注意证书名可以是"*.example.com "的格式，因此证书名不是必须为域名。如果本字段为nil，Certificates字段的第一个成员会被用于所有连接
		RootCAs:                  roots,                   // 定义权威根证书，客户端会在验证服务端证书时用到本字段。如果RootCAs是nil，TLS会使用主机的根CA池
		NextProtos:               nil,                     // 可以支持的应用层协议的列表
		ServerName:               "",                     // 用于认证返回证书的主机名（除非设置了InsecureSkipVerify）。也被用在客户端的握手里，以支持虚拟主机。
		ClientAuth:               tls.NoClientCert,        // 决定服务端的认证策略，默认是NoClientCert
		ClientCAs:                nil,                     // 定义权威根证书，服务端会在采用ClientAuth策略时使用它来认证客户端证书
		InsecureSkipVerify:       true,                    // 控制客户端是否跳过认证服务端的证书链和主机名。如果InsecureSkipVerify为true，TLS连接会接受服务端提供的任何证书和该证书中的任何主机名。此时，TLS连接容易遭受中间人攻击，这种设置只应用于测试
		CipherSuites:             nil,                     // 支持的加密组合列表。如果CipherSuites为nil，TLS连接会使用本包的实现支持的密码组合列表
		PreferServerCipherSuites: false,                   // 本字段控制服务端是选择客户端最期望的密码组合还是服务端最期望的密码组合。如果本字段为true，服务端会优先选择CipherSuites字段中靠前的密码组合使用
		SessionTicketsDisabled:   false,                   // 可以设为false以关闭会话恢复支持
		SessionTicketKey:         [32]byte{},                     // 被TLS服务端用于提供会话恢复服务。如果本字段为零值，它会在第一次服务端握手之前填写上随机数据。如果多个服务端都在终止和同一主机的连接，它们应拥有相同的SessionTicketKey。如果SessionTicketKey泄露了，使用该键的之前的记录和未来的TLS连接可能会被盗用
		ClientSessionCache:       nil,                     // 是ClientSessionState的缓存，用于恢复TLS会话
		MinVersion:               0,                       // 可接受的最低SSL/TLS版本。如果为0，会将SSLv3作为最低版本
		MaxVersion:               0,                       // 可接受的最高SSL/TLS版本。如果为0，会将本包使用的版本作为最高版本，目前是TLS 1.2
		CurvePreferences:         nil,                     // 用于ECDHE握手的椭圆曲线的ID，按优先度排序。如为空，会使用默认值
	}

	// 客户端。使用证书访问tls网络
	client(roots)
	clientWithDialer(roots)
	// 服务端。创建一个认证证书的监听
	server(cert)

	// 使用conn作为下层传输接口返回一个客户端TLS连接。配置参数config必须是非nil的且必须设置了ServerName或者InsecureSkipVerify字段
	tls.Client(conn, cfg)
	// 使用conn作为下层传输接口返回一个服务端TLS连接。配置参数config必须是非nil的且必须含有至少一个证书
	tls.Server(conn, cfg)
}

// 声明根权威证书内容
const RootPEM = `
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

const CertPem = `
-----BEGIN CERTIFICATE-----
MIIBhTCCASugAwIBAgIQIRi6zePL6mKjOipn+dNuaTAKBggqhkjOPQQDAjASMRAw
DgYDVQQKEwdBY21lIENvMB4XDTE3MTAyMDE5NDMwNloXDTE4MTAyMDE5NDMwNlow
EjEQMA4GA1UEChMHQWNtZSBDbzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABD0d
7VNhbWvZLWPuj/RtHFjvtJBEwOkhbN/BnnE8rnZR8+sbwnc/KhCk3FhnpHZnQz7B
5aETbbIgmuvewdjvSBSjYzBhMA4GA1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggr
BgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1UdEQQiMCCCDmxvY2FsaG9zdDo1
NDUzgg4xMjcuMC4wLjE6NTQ1MzAKBggqhkjOPQQDAgNIADBFAiEA2zpJEPQyz6/l
Wf86aX6PepsntZv2GYlA5UpabfT2EZICICpJ5h/iI+i341gBmLiAFQOyTDT+/wQc
6MF9+Yw1Yy0t
-----END CERTIFICATE-----
`
const KeyPem = `
-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIIrYSSNQFaA2Hwf1duRSxKtLYX5CB04fSeQ6tF1aY/PuoAoGCCqGSM49
AwEHoUQDQgAEPR3tU2Fta9ktY+6P9G0cWO+0kETA6SFs38GecTyudlHz6xvCdz8q
EKTcWGekdmdDPsHloRNtsiCa697B2O9IFA==
-----END EC PRIVATE KEY-----
`

func client(roots *x509.CertPool) {

	// 使用net.Dial连接指定的网络和地址，然后发起TLS握手，返回生成的TLS连接。Dial会将nil的配置视为零值的配置
	conn, err := tls.Dial("tcp", "www.baidu.com:443", &tls.Config{RootCAs: roots})
	if err != nil {
		panic("failed to connect: " + err.Error())
	}

	// 关闭连接
	if err := conn.Close(); err != nil {
		panic(err)
	}
}

func clientWithDialer(roots *x509.CertPool) {

	// 初始化一个net.Dialer,地址建立连接时的参数
	dialer := new(net.Dialer)
	dialer.Timeout = time.Minute * 2 // dial操作等待连接建立的最大时长，默认值代表没有超时
	dialer.Deadline = time.Now().Add(time.Minute * 2) // 一个具体的时间点期限，超过该期限后，dial操作就会失败

	// 使用dialer.Dial连接指定的网络和地址，然后发起TLS握手，返回生成的TLS连接。dialer中的超时和期限设置会将连接和TLS握手作为一个整体来应用
	conns, err := tls.DialWithDialer(dialer, "tcp", "www.baidu.com:443", &tls.Config{RootCAs: roots})
	if err != nil {
		log.Fatal(err)
	}
	if err := conns.Close(); err != nil {
		log.Fatal(err)
	}
}

func server(cert tls.Certificate) {

	// TLS客户端或服务端配置
	// 在本类型的值提供给TLS函数后，就不应再修改该值。Config类型值可能被重用；tls包也不会修改它
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}

	// 创建一个TLS监听器，使用net.Listen函数接收给定地址上的连接。配置参数config必须是非nil的且必须含有至少一个证书
	listener, err := tls.Listen("tcp", ":2000", cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 关闭监听，并使任何阻塞的Accept操作都会不再阻塞并返回错误
	defer listener.Close()

	// 返回该接口的网络地址
	listener.Addr()

	// 创建一个TLS监听器，该监听器接受inner接收到的每一个连接，并调用Server函数包装这些连接。配置参数config必须是非nil的且必须含有至少一个证书
	netListener := tls.NewListener(listener, cfg)
	if err := netListener.Close(); err != nil {
		log.Fatal(err)
	}

	for {
		// 等待并返回下一个连接到该接口的连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go func(c net.Conn) {

			// 在这些完成需要做的事

			// 关闭连接
			c.Close()
		}(conn)
	}
}
