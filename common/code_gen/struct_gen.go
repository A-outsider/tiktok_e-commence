package code_gen

type Service struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type Mysql struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	Dbname string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset string `yaml:"charset"`
}

type ElasticSearch struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	CertificateFingerprint string `yaml:"certificate_fingerprint"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type Jaeger struct {
	Host string `yaml:"host"`
	Port int `yaml:"port"`
}

type config struct {
	Service Service `yaml:"service"`
	Mysql Mysql `yaml:"mysql"`
	ElasticSearch ElasticSearch `yaml:"elasticSearch"`
	Jaeger Jaeger `yaml:"jaeger"`
}
