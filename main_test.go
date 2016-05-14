package payfriendz

import (
	"fmt"
	"testing"

	flatbuffers "github.com/google/flatbuffers/go"
	types "github.com/secmask/contact"
)

func TestProcess(t *testing.T) {
	message := newTestMessage()

	b := flatbuffers.NewBuilder(0)

	data := Serialize(b, message)

	deserMessage, err := Deserialize(data)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if deserMessage.Id != "1" {
		t.Fatalf("Wrong Deserialize")
	}

	fmt.Printf("%+v \n", deserMessage)
}

func newTestMessage() (message *types.Message) {
	message = &types.Message{
		Id:        "1",
		Receivers: []string{"test1", "test2", "test3"},
		Contacts: []types.Contact{
			types.Contact{
				Id:          "1",
				FirstName:   "foo",
				LastName:    "bar",
				Description: "foobar",
				Phones: []types.Phone{
					types.Phone{
						PhoneType: "home",
						Number:    "123456789",
					},
					types.Phone{
						PhoneType: "mobile",
						Number:    "987654321",
					},
				},
			},
		},
	}
	return
}
