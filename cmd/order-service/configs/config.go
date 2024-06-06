package configs

type Config struct {
	Database struct {
		User     string `toml:"user"`
		Password string `toml:"password"`
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		DBName   string `toml:"dbname"`
	}
	Server struct {
		Port int `toml:"port"`
	}
	UserService struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
	}
}
