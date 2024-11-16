package code_gen

type PhotoCaptcha struct {
	MaxSkew  float64 `yaml:"maxSkew"`
	DotCount int     `yaml:"dotCount"`
	Expire   string  `yaml:"expire"`
	Height   int     `yaml:"height"`
	Width    int     `yaml:"width"`
	Length   int     `yaml:"length"`
}

type Jwt struct {
	RefreshExpireTime string `yaml:"refreshExpireTime"`
	AccessSecret      string `yaml:"accessSecret"`
	RefreshSecret     string `yaml:"refreshSecret"`
	Issuer            string `yaml:"issuer"`
	AccessExpireTime  string `yaml:"accessExpireTime"`
}

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type Phone struct {
	ExpirationTime  string `yaml:"expirationTime"`
	SendInterval    string `yaml:"sendInterval"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	RegionId        string `yaml:"regionId"`
}

type Email struct {
	Host           interface{} `yaml:"host"`
	Name           string      `yaml:"name"`
	Password       string      `yaml:"password"`
	Port           int         `yaml:"port"`
	ExpirationTime string      `yaml:"expiration_time"`
	SendInterval   string      `yaml:"sendInterval"`
	Addresses      interface{} `yaml:"addresses"`
	Email          interface{} `yaml:"email"`
}

type Mysql struct {
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Password struct {
	ErrorLimit    int    `yaml:"ErrorLimit"`
	ErrorLockTime string `yaml:"ErrorLockTime"`
}

type config struct {
	PhotoCaptcha PhotoCaptcha `yaml:"photoCaptcha"`
	Jwt          Jwt          `yaml:"jwt"`
	Service      Service      `yaml:"service"`
	Redis        Redis        `yaml:"redis"`
	Phone        Phone        `yaml:"phone"`
	Email        Email        `yaml:"email"`
	Mysql        Mysql        `yaml:"mysql"`
	Jaeger       Jaeger       `yaml:"jaeger"`
	Password     Password     `yaml:"password"`
}
