package code_gen

type Service struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type Mysql struct {
	Charset string `yaml:"charset"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Dbname string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type Config struct {
	Service Service `yaml:"service"`
	Mysql Mysql `yaml:"mysql"`
	Jaeger Jaeger `yaml:"jaeger"`
}
