package config

type DbConfig struct {
	Addr string `json:"addr"`
	Port string `json:"port"`
	Password string `json:"password"`
	UserName string `json:"user_name"`
}
