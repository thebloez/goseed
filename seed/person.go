package seed

import (
	"github.com/bxcodec/faker/v3"
)

func (s Seed) PersonSeed() {
	for i := 0; i < 10; i++ {
		//prepare the statement
		stmt, _ := s.db.Prepare(`INSERT INTO person(first_name, last_name) VALUES (?,?)`)
		// execute query
		_, err := stmt.Exec(faker.AmountWithCurrency(), faker.TitleMale())
		if err != nil {
			panic(err)
		}
	}
}