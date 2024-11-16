package code_gen

type Mysql struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Dbname string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset string `yaml:"charset"`
}

type Phone struct {
	AccessKeyId string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	RegionId string `yaml:"regionId"`
	ExpirationTime string `yaml:"expiration_time"`
	SendInterval string `yaml:"sendInterval"`
}

type Email struct {
	Port int `yaml:"port"`
	ExpirationTime string `yaml:"expiration_time"`
	SendInterval string `yaml:"sendInterval"`
	Addresses interface{} `yaml:"addresses"`
	Email interface{} `yaml:"email"`
	Host interface{} `yaml:"host"`
	Name string `yaml:"name"`
	Password string `yaml:"password"`
}

type Password struct {
	ErrorLimit int `yaml:"ErrorLimit"`
	ErrorLockTime string `yaml:"ErrorLockTime"`
}

type Service struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type Redis struct {
	Port int `yaml:"port"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type PhotoCaptcha struct {
	DotCount int `yaml:"dotCount"`
	Expire string `yaml:"expire"`
	Height int `yaml:"height"`
	Width int `yaml:"width"`
	Length int `yaml:"length"`
	MaxSkew float64 `yaml:"maxSkew"`
}

type Jwt struct {
	RefreshSecret string `yaml:"refreshSecret"`
	Issuer string `yaml:"issuer"`
	AccessExpireTime string `yaml:"accessExpireTime"`
	RefreshExpireTime string `yaml:"refreshExpireTime"`
	AccessSecret string `yaml:"accessSecret"`
}

type config struct {
	Mysql Mysql `yaml:"mysql"`
	Phone Phone `yaml:"phone"`
	Email Email `yaml:"email"`
	Password Password `yaml:"password"`
	Service Service `yaml:"service"`
	Redis Redis `yaml:"redis"`
	Jaeger Jaeger `yaml:"jaeger"`
	PhotoCaptcha PhotoCaptcha `yaml:"photoCaptcha"`
	Jwt Jwt `yaml:"jwt"`
}
