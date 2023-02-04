package stub

type User struct {
}

type Pet struct {
	Name string
}

type Person struct {
}

type Entities interface {
	GetUser(string) (User, error)
	GetPets(string) ([]Pet, error)
	GetChildren(string) ([]Person, error)
	GetFriends(string) ([]Person, error)
	SaveUser(User) error
}

type Logic struct {
	Entities Entities
}

func (l Logic) GetPetNames(userID string) ([]string, error) {
	pets, err := l.Entities.GetPets(userID)
	if err != nil {
		return nil, err
	}
	r := make([]string, 0, len(pets))
	for _, pet := range pets {
		r = append(r, pet.Name)
	}
	return r, nil

}
