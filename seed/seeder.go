package seed

import (
	"database/sql"
	"log"
	"reflect"
)

type Seed struct {
	db *sql.DB
}

func Execute(db *sql.DB, seedMethodNames ...string)  {
	s := Seed{db}

	seedType := reflect.TypeOf(s)

	// execute all seeders if no method name is given
	if len(seedMethodNames) == 0 {
		log.Printf("Running all seeder ...")
		// we are looping over the method on a seed struct
		for i := 0; i < seedType.NumMethod(); i++ {
			// get the method in the current iteration
			method := seedType.Method(i)
			// execute seeder
			seed(s, method.Name)
		}
	}

	// execute only the given method names
	for _, item := range seedMethodNames{
		seed(s, item)
	}
}

func seed(s Seed, seedMethodName string) {
	// Get the reflect value of the method
	m := reflect.ValueOf(s).MethodByName(seedMethodName)
	// Exit if the method doesn't exist
	if !m.IsValid() {
		log.Fatal("No Method called ", seedMethodName)
	}

	// execute the method
	log.Println("Seeding ", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed ", seedMethodName," Succeed")
}
