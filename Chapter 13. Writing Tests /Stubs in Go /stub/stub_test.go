package stub

import (
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

type GetPetNamesStub struct {
	Entities
}

func (ps GetPetNamesStub) GetPets(userId string) ([]Pet, error) {
	switch userId {
	case "1":
		return []Pet{{"Bubbles"}}, nil
	case "2":
		return []Pet{{"Stampy"}, {"Snowball II"}}, nil
	default:
		return nil, fmt.Errorf("invalid id : %s", userId)
	}
}

type EntitiesStub struct {
	getUser     func(string) (User, error)
	getPets     func(string) ([]Pet, error)
	getChildren func(string) ([]Person, error)
	getFriends  func(string) ([]Person, error)
	saveUser    func(User) error
}

func (e EntitiesStub) GetUser(s string) (User, error) {
	return e.getUser(s)
}

func (e EntitiesStub) GetPets(s string) ([]Pet, error) {
	return e.getPets(s)
}

func (e EntitiesStub) GetChildren(s string) ([]Person, error) {
	return e.getChildren(s)
}

func (e EntitiesStub) GetFriends(s string) ([]Person, error) {
	return e.getFriends(s)
}

func (e EntitiesStub) SaveUser(user User) error {
	return e.saveUser(user)
}

func TestLogic_GetPetNames(t *testing.T) {
	type fields struct {
		Entities Entities
	}
	type args struct {
		userID string
	}
	stub := GetPetNamesStub{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{name: "user 1", fields: fields{stub}, args: args{"1"}, want: []string{"Bubbles"}, wantErr: false},
		{name: "user 2", fields: fields{stub}, args: args{"2"}, want: []string{"Stampy", "Snowball II"}, wantErr: false},
		{name: "user 3", fields: fields{stub}, args: args{"3"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		l := Logic{
			Entities: tt.fields.Entities,
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := l.GetPetNames(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPetNames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPetNames() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogic_GetPetNames2(t *testing.T) {
	data := []struct {
		name     string
		getPets  func(string) ([]Pet, error)
		userID   string
		petNames []string
		errMsg   string
	}{
		{"user 1", func(userId string) ([]Pet, error) {
			return []Pet{{"Bubbles"}}, nil
		}, "1", []string{"Bubbles"}, ""},
		{"user 3", func(userId string) ([]Pet, error) {
			return nil, errors.New("invalid id " + userId)
		}, "3", nil, "invalid id 3"},
	}
	logic := Logic{}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			logic.Entities = EntitiesStub{getPets: d.getPets}
			petNames, err := logic.GetPetNames(d.userID)
			if diff := cmp.Diff(d.petNames, petNames); diff != "" {
				t.Error(diff)
			}
			errMsg := ""
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("expected error %s, got %s", d.errMsg, errMsg)
			}
		})
	}
	fmt.Println(data)
}
