package serializer

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"google.golang.org/protobuf/proto"
	"log"
)

func TestProtoBuff() {
	fmt.Println(" ~~~~~~~~~~~~~~~~~~ PROTOBUFF SERIALIZER START ~~~~~~~~~~~~~~~~~~ ")
	jsonString := serializeProtoBuff()
	fmt.Println("Encoded Serialized String ->>>> " + jsonString)

	user := deSerializerProtoBuff(jsonString)
	fmt.Println("DeSerialized Object --->>> "+user.Username, user.FullName, user.Email, user.Dob)
	fmt.Println(" ~~~~~~~~~~~~~~~~~~ PROTOBUFF SERIALIZER END ~~~~~~~~~~~~~~~~~~ ")
}

func serializeProtoBuff() string {
	user := User{
		Username: "ProtoBuff",
		FullName: "Proto Buff",
		Email:    "protobuff@yopmail.com",
		Dob:      toml.LocalDate{Year: 2000, Month: 1, Day: 1},
	}
	// Serialization
	data, err := proto.Marshal(nil)
	if err != nil {
		log.Fatal("Marshaling error: ", user)
	}
	fmt.Println("Serialized Protobuf:", data)
	return string(data)
}

func deSerializerProtoBuff(jsonString string) User {

	user := User{}
	// Deserialization
	err := proto.Unmarshal([]byte(jsonString), nil)
	if err != nil {
		log.Fatal("Unmarshaling error: ", err)
	}
	return user

}
