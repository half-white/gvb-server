package config

type Jwt struct {
	Secret  string `json:"secret" yaml:"secret"`   //密钥
	Expries int    `json:"expries" yaml:"expries"` //过期时间
	Issuer  string `json:"issuer" yaml:"issuer"`   //颁发人
}
