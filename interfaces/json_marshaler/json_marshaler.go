package json_marshaler

import (
	"encoding/json"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	NickNames []string `json:"nicknames"`
}

func (p *Person) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
}

func (p *Person) UnmarshalJSON(data []byte) error {
	s := &struct{ Name string }{}
	err := json.Unmarshal(data, s)
	p.Name = s.Name + "!!!"
	return err
}
