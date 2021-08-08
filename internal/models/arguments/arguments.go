package arguments

import (
	"encoding/json"
	"log"
)

type Arguments struct {
	Int1  int
	Int2  int
	Limit int
	Str1  string
	Str2  string
}

func New(int1 int, int2 int, limit int, str1 string, str2 string) Arguments {
	return Arguments{int1, int2, limit, str1, str2}
}

func NewFromJson(s string) (Arguments, error) {
	args := &Arguments{}
	err := json.Unmarshal([]byte(s), args)
	if err != nil {
		log.Printf("Failed to decode arguments from JSON: %s\n", err.Error())
		return *args, err
	}
	return *args, nil
}

func (a Arguments) AsJson() (string, error) {
	bytes, err := json.Marshal(a)
	if err != nil {
		log.Printf("Failed to encode arguments as JSON: %s\n", err.Error())
		return "", err
	}
	return string(bytes), nil
}
