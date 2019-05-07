package config

import (
	"github.com/unionpay/common"
)

func InitConfig() error {
	var err error
	//初始化证书
	//pfxpath和pfxpwd同时存在，privatepath和certpath同时存在，两组传任意一组皆可
	err = common.LoadCert(&common.Config{
		//银联提供的pfx证书存放路径,商户私钥
		PfxPath: "static/CA/700000000000001_acp.p12",

		//pfx证书密码
		PfxPwd: "000000",

		//数据加密证书路径，银联公钥
		EncryptCertPath: "static/CA/verify_sign_acp.cer",

		//用户私钥地址，私钥通过pfx解析得到
		// openssl pkcs12 -in xxxx.pfx -nodes -out server.pem 生成为原生格式pem 私钥
		// openssl rsa -in server.pem -out server.key  生成为rsa格式私钥文件
		PrivatePath: "static/CA/server.key",

		//用户证书，通过pfx解析得到
		// openssl pkcs12 -in xxxx.pfx -clcerts -nokeys -out key.cert
		CertPath: "static/CA/key.cert",
	})
	//配置环境信息
	common.SetConfig(&common.Config{
		// 商户号
		MerId: "777290058169193",
		//域名信息
		Url: "https://gateway.test.95516.com",
	})
	return err
}

