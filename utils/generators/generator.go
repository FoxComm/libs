package generators

import "github.com/manveru/faker"

var fake *faker.Faker

func init() {
	var err error
	fake, err = faker.New("en")
	if err != nil {
		panic(err)
	}
}
