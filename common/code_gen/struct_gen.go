package code_gen

type Service struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Password string `yaml:"password"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type Config struct {
	Service Service `yaml:"service"`
	Redis Redis `yaml:"redis"`
	Jaeger Jaeger `yaml:"jaeger"`
}
