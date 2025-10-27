package config 
import(
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct{
	DbUser string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbHost string `env:"DB_HOST"`
	DbPort string `env:"DB_PORT"`
	DbName string `env:"DB_NAME"`
	Port   string `env:"PORT"`
}
var AppConfig Config

func Initconfig(){
	err:= godotenv.Load()
	if err!= nil{
		log.Fatalf("Error loading .env file")
	}

	AppConfig = Config{
		DbUser: os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbName: os.Getenv("DB_NAME"),
		Port: os.Getenv("PORT"),
	}

	InitDB()
}

func GetConfig() *Config{
	return &AppConfig
}