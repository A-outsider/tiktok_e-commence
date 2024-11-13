package code_gen

type Mysql struct {
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type PhotoCaptcha struct {
	Height   int     `yaml:"height"`
	Width    int     `yaml:"width"`
	Length   int     `yaml:"length"`
	MaxSkew  float64 `yaml:"maxSkew"`
	DotCount int     `yaml:"dotCount"`
	Expire   string  `yaml:"expire"`
}

type Password struct {
	ErrorLimit    int    `yaml:"ErrorLimit"`
	ErrorLockTime string `yaml:"ErrorLockTime"`
}

type Jwt struct {
	AccessExpireTime  string `yaml:"accessExpireTime"`
	RefreshExpireTime string `yaml:"refreshExpireTime"`
	AccessSecret      string `yaml:"accessSecret"`
	RefreshSecret     string `yaml:"refreshSecret"`
	Issuer            string `yaml:"issuer"`
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
	SendInterval string `yaml:"sendInterval"`
}

type Email struct {
	Password       string      `yaml:"password"`
	Port           int         `yaml:"port"`
	ExpirationTime string      `yaml:"expiration_time"`
	SendInterval   string      `yaml:"sendInterval"`
	Addresses      interface{} `yaml:"addresses"`
	Email          interface{} `yaml:"email"`
	Host           interface{} `yaml:"host"`
	Name           string      `yaml:"name"`
}

type config struct {
	Mysql        Mysql        `yaml:"mysql"`
	Jaeger       Jaeger       `yaml:"jaeger"`
	PhotoCaptcha PhotoCaptcha `yaml:"photoCaptcha"`
	Password     Password     `yaml:"password"`
	Jwt          Jwt          `yaml:"jwt"`
	Service      Service      `yaml:"service"`
	Redis        Redis        `yaml:"redis"`
	Phone        Phone        `yaml:"phone"`
	Email        Email        `yaml:"email"`
}
