package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

// func InitEnv() *AppConfig {
// 	if err := godotenv.Load(".env"); err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	return &AppConfig{
// 		DBUSERNAME: os.Getenv("DBUSERNAME"),
// 		DBPASS:     os.Getenv("DBPASS"),
// 		DBHOST:     os.Getenv("DBHOST"),
// 		DBPORT:     os.Getenv("DBPORT"),
// 		DBNAME:     os.Getenv("DBNAME"),
// 		JWTKEY:     os.Getenv("SECRETKEY"),
// 	}
// }

func InitEnv() *AppConfig {
	fmt.Println("MASUK")
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("SECRETKEY"); found {
		app.JWTKEY = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBUSERNAME"); found {
		app.DBUSERNAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DBPASS = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHOST = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		app.DBPORT = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBNAME = val
		isRead = false
	}

	// err2 := viper.Unmarshal(&app)
	// if err2 != nil {
	// 	log.Println("error parse config : ", err2.Error())
	// 	return nil
	// }

	if isRead {
		viper.AddConfigPath("./")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}

		app.JWTKEY = viper.Get("SECRETKEY").(string)
		app.DBUSERNAME = viper.Get("DBUSERNAME").(string)
		app.DBPASS = viper.Get("DBPASS").(string)
		app.DBHOST = viper.Get("DBHOST").(string)
		app.DBPORT, _ = viper.Get("DBPORT").(string)
		app.DBNAME = viper.Get("DBNAME").(string)
	}

	return &app
}
