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

type Static struct {
	AvatarPath string `yaml:"avatar_path"`
}

type Config struct {
	Service Service `yaml:"service"`
	Mysql Mysql `yaml:"mysql"`
	Static Static `yaml:"static"`
}
