package code_gen

type Mysql struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Dbname string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset string `yaml:"charset"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Password string `yaml:"password"`
}

type config struct {
	Host string `yaml:"host"`
	Mysql Mysql `yaml:"mysql"`
	Jaeger Jaeger `yaml:"jaeger"`
	Redis Redis `yaml:"redis"`
}
