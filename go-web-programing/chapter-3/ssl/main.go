//生成个人使用的 ssl证书以及服务器私钥
// 本文件尚未通过测试，暂存
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main()  {
	max := new(big.Int).Lsh(big.NewInt(1), 128) // 创建一个大整数，并且末尾 128 全部置为 0，首位为 1
	serialNumber, _ := rand.Int(rand.Reader, max) // 创建一个最大值为 max 的随机数

	// Name表示X.509专有名称。这仅包括DN的公共元素。解析时，所有元素都存储在Names中，并且可以从那里提取非标准元素。编组时，会追加ExtraNames中的元素并覆盖具有相同OID的其他值。
	subject := pkix.Name{
		Organization: []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName: "Go Web Programming",
	}

	template := x509.Certificate{ // Certificate 证书代表X.509证书。
		SerialNumber: serialNumber,
		Subject: subject,
		NotBefore: time.Now(),
		NotAfter: time.Now(),
		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey{pk}})
	keyOut.Close()