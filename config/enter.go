package config

type Config struct {
	Mysql    Mysql    `yaml:"mysql"`
	Logger   Logger   `yaml:"logger"`
	System   System   `yaml:"system"`
	SiteInfo SiteInfo `yaml:"site_info"`
	QQ       QQ       `yaml:"qq"`
	Qiniu    Qiniu    `yaml:"qi_niu"`
	Email    Email    `yaml:"email"`
	Jwt      Jwt      `yaml:"jwt"`
	Upload   Upload   `yaml:"upload"`
}
