package config

type Config struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Bucket    string
	Region    string
}

// /api/v1/service-account-credentials
// play.min.io
func LoadConfig() *Config {
	return &Config{
		Endpoint:  "172.17.134.233:9000",                      // Например: "127.0.0.1:9000"
		AccessKey: "PwUeK87TDG10BAJU0fIM",                     // Ваш access key
		SecretKey: "8LwMjYYn43lhYggeesZPNw4LDTJP5HxGh3IiAILS", // Ваш secret key
		Bucket:    "music-bucket",                             // Имя бакета
		Region:    "",                                         // Регион, если есть
	}
}
