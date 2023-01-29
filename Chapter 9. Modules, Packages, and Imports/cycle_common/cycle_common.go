package cycle_common

import (
	"ch9/person"
	"ch9/pet"
)

var pets = map[string]pet.Pet{
	"Hello": {
		Name:      "Hello dog",
		Age:       1,
		OwnerName: "Hello",
	},
}
var owners = map[string]person.Person{
	"Hello": {"Hello", 18, "Hello dog"},
}
