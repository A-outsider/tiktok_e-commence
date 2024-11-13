package code_gen

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
	Policy string `yaml:"policy"`
	Model  string `yaml:"model"`
}

type Jwt struct {
	Issuer            string `yaml:"issuer"`
	AccessExpireTime  string `yaml:"accessExpireTime"`
	RefreshExpireTime string `yaml:"refreshExpireTime"`
	AccessSecret      string `yaml:"accessSecret"`
	RefreshSecret     string `yaml:"refreshSecret"`
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
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
}

type config struct {
	Redis      Redis      `yaml:"redis"`
	Jaeger     Jaeger     `yaml:"jaeger"`
	Role       Role       `yaml:"role"`
	Jwt        Jwt        `yaml:"jwt"`
	VisitLimit VisitLimit `yaml:"visitLimit"`
	Service    Service    `yaml:"service"`
	Mysql      Mysql      `yaml:"mysql"`
}
