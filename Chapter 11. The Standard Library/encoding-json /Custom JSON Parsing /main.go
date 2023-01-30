package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	Id          string      `json:"id"`
	DateOrdered RFC822ZTime `json:"date_ordered"`
	CustomerId  string      `json:"customer_id"`
	Items       []Item      `json:"items"`
}
type RFC822ZTime struct {
	time.Time
}

func (r RFC822ZTime) MarshalJSON() ([]byte, error) {
	out := r.Time.Format(time.RFC822Z)
	return []byte(`"` + out + `"`), nil
}
func (r *RFC822ZTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	t, err := time.Parse(fmt.Sprintf(`"%s"`, time.RFC822Z), string(b))
	if err != nil {
		return err
	}
	*r = RFC822ZTime{t}
	return nil
}

func main() {
	data := `{
  "id":"12345",
  "date_ordered":"01 May 20 13:01 +0800",
  "customer_id":"3",
  "items":[{"id":"xyz123","name":"Thing 1"},{"id":"abc789","name":"Thing 2"}]
}`
	var order Order
	err := json.Unmarshal([]byte(data), &order)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", order)

	bytes, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))

}
