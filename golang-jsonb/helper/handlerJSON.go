package helper

import (
	"encoding/json"
	"golang-jsonb/model/entity"
)

type Message struct {
	Content string `json:"message"`
}

func WriteJSONPeople(data []entity.Person) []byte {
	output, err := json.MarshalIndent(data, "", "   ")
	CheckErrorFatal(err)

	return output
}

func WriteJSONPerson(data entity.Person) []byte {
	output, err := json.MarshalIndent(data, "", "   ")
	CheckErrorFatal(err)

	return output
}

func WriteJSONMessage(message string) []byte {
	var msg Message
	msg.Content = message

	output, err := json.MarshalIndent(msg, "", "   ")
	CheckErrorFatal(err)

	return output
}
