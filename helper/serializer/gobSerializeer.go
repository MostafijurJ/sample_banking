package serializer

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/pelletier/go-toml/v2"
)

func TestGob() {
	fmt.Println(" ~~~~~~~~~~~~~~~~~~ GOB SERIALIZER START ~~~~~~~~~~~~~~~~~~ ")
	jsonString := serializeGob()
	fmt.Println("Encoded Serialized String ->>>> " + jsonString)

	user := deSerializerGob(jsonString)
	fmt.Println("DeSerialized Object --->>> "+user.Username, user.FullName, user.Email, user.Dob)
	fmt.Println(" ~~~~~~~~~~~~~~~~~~ GOB SERIALIZER END ~~~~~~~~~~~~~~~~~~ ")
}

func deSerializerGob(jsonString string) User {
	var user User
	decoder := gob.NewDecoder(bytes.NewReader([]byte(jsonString)))
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println("Error deserializing:", err)
		return User{}
	}
	return user

}

func serializeGob() string {
	user := User{
		Username: "Gob",
		FullName: "Gob Bluth",
		Email:    "gob@yopmail.com",
		Dob:      toml.LocalDate{Year: 1999, Month: 1, Day: 1},
	}

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(user)
	if err != nil {
		fmt.Println("Error serializing:", err)
		return ""
	}

	fmt.Println("Serialized Gob --->>> ", buffer.Bytes())
	return string(buffer.Bytes())
}
