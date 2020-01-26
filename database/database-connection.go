package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type User struct{
	Email string `json:"email"`
	Password int `json:"password"`
}

type Config struct{
	DBProperties struct{
		Username      string `yaml:"user"`
		Password      string `yaml:"password"`
		Port          string `yaml:"port"`
		Database_name string `yaml:"database_name"`
		Address       string `yaml:"address"`
	} `yaml:"database"`
}
func GetDatabase() (*sql.DB, error) {

	f, err := os.Open("application.yml")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Panic(err)
	}
	db, err := sql.Open("mysql", cfg.DBProperties.Username+ ":" + cfg.DBProperties.Password+ "@tcp(" + cfg.DBProperties.Address+ ":" + cfg.DBProperties.Port+ ")/" + cfg.DBProperties.Database_name)
	if err != nil {
		log.Panic(err.Error())
	}
	log.Println( "DB Connection Successful")
	return db, nil
}
