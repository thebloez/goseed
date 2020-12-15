package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"seeder/seed"
)

func main() {
	godotenv.Load()
	//handleArgs()
	//fmt.Println(os.Getenv("DB_USERNAME"))
	withoutArgs("PersonSeed")
}

func handleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			connString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&multiStatements=true", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
			// connect DB
			db, err := sql.Open("mysql", connString)
			if err != nil {
				log.Fatal("Error Opening DB: %v", err)
			}

			seed.Execute(db, args[1:]...)
			os.Exit(0)
		}
	}
}

func withoutArgs(methodNames string){
	connString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&multiStatements=true", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	// connect DB
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal("Error Opening DB: %v", err)
	}

	seed.Execute(db, methodNames)
	os.Exit(0)
}
