package config

type Config struct {
	DbName     string
	DbHost     string
	DbUser     string //Db - database//
	DbPassword string
}

func GetConfig() Config {
	return Config{
		DbHost:     "localhost",
		DbName:     "Clever-House",
		DbUser:     "postgres",
		DbPassword: "postgres",
	}
}
