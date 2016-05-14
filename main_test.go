package payfriendz

import (
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

	if deserMessage.Id != message.Id {
		t.Fatalf("Wrong Message Id")
	}

	if len(deserMessage.Receivers) != len(message.Receivers) {
		t.Fatalf("Wrong Message Receivers Length")
	}

	if deserMessage.Receivers[0] != message.Receivers[0] {
		t.Fatalf("Wrong Message Receiver")
	}

	if len(deserMessage.Contacts) != len(message.Contacts) {
		t.Fatalf("Wrong Message Contacts Length")
	}

	if deserMessage.Contacts[0].Id != message.Contacts[0].Id {
		t.Fatalf("Wrong Message Contact")
	}

	if len(deserMessage.Contacts[0].Phones) != len(message.Contacts[0].Phones) {
		t.Fatalf("Wrong Contact Phones Length")
	}

	if deserMessage.Contacts[0].Phones[0].Number != message.Contacts[0].Phones[0].Number {
		t.Fatalf("Wrong Contact Phone")
	}
}

func newTestMessage() (message *types.Message) {
	p1 := types.Phone{
		PhoneType: "home",
		Number:    "123456789",
	}
	p2 := types.Phone{
		PhoneType: "mobile",
		Number:    "987654321",
	}

	c1 := types.Contact{
		Id:          "1",
		FirstName:   "foo",
		LastName:    "bar",
		Description: "foobar",
		Phones:      []types.Phone{p1, p2},
	}

	c2 := types.Contact{
		Id:          "2",
		FirstName:   "foo",
		LastName:    "bar",
		Description: "foobar",
		Phones:      []types.Phone{p1, p2},
	}

	message = &types.Message{
		Id:        "1",
		Receivers: []string{"test1", "test2", "test3"},
		Contacts:  []types.Contact{c1, c2},
	}
	return
}
