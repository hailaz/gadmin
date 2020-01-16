package common

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
)

var CryptoKeyList = gmap.New()

type CryptoKey struct {
	Id         string `json:"kid"`
	CryptoType string `json:"cryptotype"`
	Key        string `json:"key"`
	PrivateKey string `json:"-"`
	TimeStamp  int64  `json:"timestamp"` //sec
}

// GetCryptoKey 获取加密key
//
// createTime:2019年04月26日 15:03:24
// author:hailaz
func GetCryptoKey(id string) *CryptoKey {
	var ck CryptoKey
	if k := CryptoKeyList.Get(id); k != nil {
		CryptoKeyList.Remove(id) //移除已使用密钥
		ck = k.(CryptoKey)
		return &ck
	}
	return nil
}

// RemoveTimeoutCryptoKey 移除超时加密key
//
// createTime:2019年04月26日 15:00:23
// author:hailaz
func RemoveTimeoutCryptoKey() {
	kList := make([]interface{}, 0)
	nowSec := gtime.Second()
	//遍历加密key
	CryptoKeyList.Iterator(func(k interface{}, v interface{}) bool {
		ck := v.(CryptoKey)
		if nowSec-ck.TimeStamp >= 10 {
			kList = append(kList, k)
		}
		return true
	})
	//移除超时的加密key
	for _, v := range kList {
		CryptoKeyList.Remove(v)
		glog.Debugf("remove key:%v", v)
	}
}

// GenCryptoKey 创建加密key
//
// createTime:2019年04月26日 10:27:48
// author:hailaz
func GenCryptoKey(id string) CryptoKey {
	rsakp, _ := GenRsaKey(2048)
	ck := CryptoKey{
		Id:         id,
		CryptoType: "RSA-PKCS1v1.5",
		Key:        rsakp.PublicKey,
		PrivateKey: rsakp.PrivateKey,
		TimeStamp:  gtime.Second(),
	}
	CryptoKeyList.Set(id, ck)
	return ck

}

type RASKeyPair struct {
	PrivateKey string
	PublicKey  string
}

// RAS密钥对生成可选1024 2048 4096
func GenRsaKey(bits int) (*RASKeyPair, error) {
	kp := &RASKeyPair{}
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	b := bytes.NewBuffer(make([]byte, 0))
	pem.Encode(b, block)
	kp.PrivateKey = b.String()
	//glog.Debug(b.String())
	// 生成公钥文件
	derPkix, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	b1 := bytes.NewBuffer(make([]byte, 0))
	pem.Encode(b1, block)
	//glog.Debug(b1.String())
	kp.PublicKey = b1.String()

	return kp, nil
}

// 加密
func RsaEncrypt(origData, publicKey []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext, privateKey []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
