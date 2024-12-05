package code_gen

type Mysql struct {
	Port int `yaml:"port"`
	Dbname string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset string `yaml:"charset"`
	Host string `yaml:"host"`
}

type ElasticSearch struct {
	Password string `yaml:"password"`
	CertificateFingerprint string `yaml:"certificate_fingerprint"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Username string `yaml:"username"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Password string `yaml:"password"`
}

type Static struct {
	ProductPath string `yaml:"product_path"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type Service struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type Config struct {
	Mysql Mysql `yaml:"mysql"`
	ElasticSearch ElasticSearch `yaml:"elasticSearch"`
	Redis Redis `yaml:"redis"`
	Static Static `yaml:"static"`
	Jaeger Jaeger `yaml:"jaeger"`
	Service Service `yaml:"service"`
}
