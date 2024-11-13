package code_gen

type Jwt struct {
	RefreshExpireTime string `yaml:"refreshExpireTime"`
	AccessSecret      string `yaml:"accessSecret"`
	RefreshSecret     string `yaml:"refreshSecret"`
	Issuer            string `yaml:"issuer"`
	AccessExpireTime  string `yaml:"accessExpireTime"`
}

type VisitLimit struct {
	RateLimitInterval string `yaml:"rateLimitInterval"`
	RateLimitCap      int    `yaml:"rateLimitCap"`
}

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Mysql struct {
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Redis struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Role struct {
	Model  string `yaml:"model"`
	Policy string `yaml:"policy"`
}

type config struct {
	Jwt        Jwt        `yaml:"jwt"`
	VisitLimit VisitLimit `yaml:"visitLimit"`
	Service    Service    `yaml:"service"`
	Mysql      Mysql      `yaml:"mysql"`
	Redis      Redis      `yaml:"redis"`
	Jaeger     Jaeger     `yaml:"jaeger"`
	Role       Role       `yaml:"role"`
}
