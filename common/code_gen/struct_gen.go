package code_gen

type Service struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset string `yaml:"charset"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Dbname string `yaml:"dbname"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type Redis struct {
	Port int `yaml:"port"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
}

type Config struct {
	Service Service `yaml:"service"`
	Mysql Mysql `yaml:"mysql"`
	Jaeger Jaeger `yaml:"jaeger"`
	Redis Redis `yaml:"redis"`
}
