package config

type Config struct {
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
}
type Mysql struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
}

type Redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password`
	DB       int    `json:"db"`
}
