package initialize

import (
	"github.com/smartwalle/alipay/v3"
	"gomall/services/payment/config"
	"log"
)

// TODO : 考虑和证书一起写入配置文件
const (
	kAppId        = "9021000142614732" // 支付宝应用id
	kPrivateKey   = "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCOcYipGJjqV2S7QnJuUoRdM/rhh8AJPJhB+PiixvqOwG9ai5O70bR6wPAoxtahhkANxylRon0kAy31f4NdNaA9LdnnhLKd7dUjkPfQ4wwYdryewsVU9Qh1liJn+twFRJyrJuUr/JSkwmmY/weSL04Zepf24g9fwzQkWMrsimH3+kEwc8ZelqjDsf6Febm/xMMTI+UvAvtGsGQVUGthRfLOsOSWYu7cG/uS6KRBpyswehem1I73KdnJhDEPqvPqdiEDalCW1focN+o+YFVfiC0TTmIQlcO7Gvo/35FkB+JfNuDo4KYfLi1woGGLz4KjQlz1puLWRZeSnNE1i/75hhqbAgMBAAECggEABlcNUyFyJPQQ2rjhaQGpPJDEuOcW6BJXYJBZWL1sh5APJMOTpsd1tgUCa342LWhRT4uuziBiW4j8sbGkQjDR2gdBbKmeXaMpWnToYtbIQgp+6L4YlGh3oOw2ydQDmQRtyLxpdOTAJPk2RAYN6zbJzh6DxwxWzNdeacO9/HtIUwYyaPqeKie8+M+Rmsn9DRubQSXII4UFRRhx5pbqGO/eNXClrnl1NMDdKr7cNpsZTdSdhSYNriLiOeqvnDvwdCA/tMBeZfZseJx22JwLNvSwqdO/5UQ8IaJgOVhzh1b5mgHhoYu0nAS8rGDFlMpiiOUePoFLLgqSLnRudfGda/n6yQKBgQDU7sqSHRfaaWZt1JFUnjMcdpEzhGj6U0vC/ZAoQP9GBcK6ziiyoyjR5barU8DWH0nATEq+Sbd3TOZm98rGHOVd78Kb+podRr0LdXwFLNxLeWim6SP7vvjj4kv9bVyQ0F2Wcrhd3qp9s+KPsevDUgicnjJTkvnuI4WVad9H6OEu9wKBgQCrQPTXgw/G2EXpeWP6JDtZfa2GHD0VNqSyBxezNd5QWFhYBY96xx3/UnKRgWY1ZDfboXyqfBNOq5OQFwmhNBjOktCW1IEn9GCtXlvyjnpq0IlXuBBUxAJUhmn9eJmCF3NufncUDuYYE1KqqQJvbgqy8nuFOoqRNaQZGM9vwDg0fQKBgQCg0+u1CLxnf4yaEB/k5ch9CyEI5E3WJOvoT1R+0vj8joVSSzx6ELpYL0UViqDwGZm+4ODjcRJdzXuI8kf58wFbPiijX1jgG/nVmdsenY+WghEFYLqI/ulGVjpHJD7yMi8931BZtkDXyPKqzhvg3ykaAnLIpQ6ZS8Mt41V+HutWaQKBgARLFZhNjdizVVVcGLiNrfs5Xl9NV+6vNwPLj7mLcS5ceKKESSuP0F21SHADaXePMqNL8h8oCyfev01OdoxXDQQoxBfz7eT9iGrwQafcEI+a+MZ9M9OcMl7CG+gh3N9ZDSjI/N1A3l3eJiVnJUt728LOt3AInq6zRJDogVLQ49fVAoGBAKgGX9UUaoI7H6fgWBIHzaGfe0cauLrj+A9F9qnwmeTJZzb4TtCG23D57oQSGNm1TOYl857CE9H4tYP6kdpp9zYyEIIlP2ScxeGI7+uQqAwmwsvWMcvd+mZxKMbXrFodSbhVrbdI8aurCG0Oib11RTkCNm8Bo1uC9xKXroRfRPpl"
	certPath      = "./data/cert/"
	GatewayDomain = " http://g47efi.natappfree.cc/api/v1" // TODO 部署到服务上就不用内网穿透了
)

func InitAlipay() {
	config.GetConf()

	var err error
	svcContext.AlipayCli, err = alipay.New(kAppId, kPrivateKey, false)
	if err != nil {
		panic(err)
	}

	// 加载证书
	if err = svcContext.AlipayCli.LoadAppCertPublicKeyFromFile(certPath + "appPublicCert.crt"); err != nil {
		log.Panicln("加载证书发生错误", err)
		return
	}
	if err = svcContext.AlipayCli.LoadAliPayRootCertFromFile(certPath + "alipayRootCert.crt"); err != nil {
		log.Panicln("加载证书发生错误", err)
		return
	}
	if err = svcContext.AlipayCli.LoadAlipayCertPublicKeyFromFile(certPath + "alipayPublicCert.crt"); err != nil {
		log.Panicln("加载证书发生错误", err)
		return
	}

}
