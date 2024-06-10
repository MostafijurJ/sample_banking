package serializer

import (
	"encoding/json"
	"fmt"
	"github.com/pelletier/go-toml/v2"
)

func TestJson() {
	fmt.Println(" ~~~~~~~~~~~~~~~~~~ JSON SERIALIZER START ~~~~~~~~~~~~~~~~~~ ")
	serializeJson := serializeJson()
	fmt.Println("Json String is - >>>> " + serializeJson)
	user := deSerializerJson(serializeJson)

	fmt.Println("Deserialized Object -->>> "+user.Username, user.FullName, user.Email, user.Dob)

	fmt.Println(" ~~~~~~~~~~~~~~~~~~ JSON SERIALIZER END ~~~~~~~~~~~~~~~~~~ ")
}

func serializeJson() string {
	user := User{
		Username: "johndoe",
		FullName: "John Doe",
		Email:    "john@yopmail.com",
		Dob:      toml.LocalDate{Year: 1980, Month: 1, Day: 1},
	}

	marshal, err := json.Marshal(user)

	if err != nil {
		return "Error while marshaling json  " + err.Error()
	}
	return string(marshal)
}

func deSerializerJson(jsonData string) User {
	var user User
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		return User{}
	}
	return user
}
