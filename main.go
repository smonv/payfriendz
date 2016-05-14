package payfriendz

import (
	"errors"

	"github.com/tthanh/payfriendz/model"

	flatbuffers "github.com/google/flatbuffers/go"
	types "github.com/secmask/contact"
)

func Serialize(b *flatbuffers.Builder, message *types.Message) []byte {
	b.Reset()

	id := b.CreateString(message.Id)

	revs := []flatbuffers.UOffsetT{}
	for _, receiver := range message.Receivers {
		revs = append(revs, b.CreateString(receiver))
	}

	model.MessageStartReceiversVector(b, len(message.Receivers))
	for _, rev := range revs {
		b.PrependUOffsetT(rev)
	}
	receivers := b.EndVector(len(message.Receivers))

	cOffsets := []flatbuffers.UOffsetT{}
	for _, contact := range message.Contacts {
		cOffsets = append(cOffsets, serializeContact(b, contact))
	}

	model.MessageStartContactsVector(b, len(message.Contacts))
	for _, cOffset := range cOffsets {
		b.PrependUOffsetT(cOffset)
	}
	contacts := b.EndVector(len(message.Contacts))

	model.MessageStart(b)

	model.MessageAddId(b, id)

	model.MessageAddReceivers(b, receivers)

	model.MessageAddContacts(b, contacts)

	m := model.MessageEnd(b)

	b.Finish(m)

	return b.Bytes[b.Head():]
}

func Deserialize(buf []byte) (*types.Message, error) {
	if len(buf) <= 0 {
		err := errors.New("Empty Buffer")
		return nil, err
	}

	message := &types.Message{}
	m := model.GetRootAsMessage(buf, 0)
	message.Id = string(m.Id())

	message.Receivers = []string{}
	for i := 0; i < m.ReceiversLength(); i++ {
		message.Receivers = append(message.Receivers, string(m.Receivers(i)))
	}

	message.Contacts = deserializeContact(m)

	return message, nil
}

func serializeContact(b *flatbuffers.Builder, c types.Contact) (contact flatbuffers.UOffsetT) {

	id := b.CreateString(c.Id)
	firstName := b.CreateString(c.FirstName)
	lastName := b.CreateString(c.LastName)
	description := b.CreateString(c.Description)

	pOffsets := []flatbuffers.UOffsetT{}
	for _, phone := range c.Phones {
		pOffsets = append(pOffsets, serializePhone(b, phone))
	}

	model.ContactStartPhonesVector(b, len(c.Phones))
	for _, pOffset := range pOffsets {
		b.PrependUOffsetT(pOffset)
	}
	phones := b.EndVector(len(c.Phones))

	model.ContactStart(b)
	model.ContactAddId(b, id)
	model.ContactAddFirstName(b, firstName)
	model.ContactAddLastName(b, lastName)
	model.ContactAddDescription(b, description)
	model.ContactAddPhones(b, phones)

	contact = model.ContactEnd(b)
	return
}

func serializePhone(b *flatbuffers.Builder, p types.Phone) flatbuffers.UOffsetT {
	phoneType := b.CreateString(p.PhoneType)
	number := b.CreateString(p.Number)

	model.PhoneStart(b)
	model.PhoneAddPhoneType(b, phoneType)
	model.PhoneAddNumber(b, number)

	return model.PhoneEnd(b)
}

func deserializeContact(message *model.Message) []types.Contact {
	contacts := []types.Contact{}
	for i := 0; i < message.ContactsLength(); i++ {
		c := &model.Contact{}
		result := message.Contacts(c, i)
		if result {
			contact := types.Contact{}
			contact.Id = string(c.Id())
			contact.FirstName = string(c.FirstName())
			contact.LastName = string(c.LastName())
			contact.Description = string(c.Description())
			contact.Phones = deserializePhone(c)

			contacts = append(contacts, contact)
		}
	}

	return contacts
}

func deserializePhone(contact *model.Contact) []types.Phone {
	phones := []types.Phone{}
	for j := 0; j < contact.PhonesLength(); j++ {
		p := &model.Phone{}
		r := contact.Phones(p, j)
		if r {
			phone := types.Phone{}
			phone.PhoneType = string(p.PhoneType())
			phone.Number = string(p.Number())

			phones = append(phones, phone)
		}
	}
	return phones
}
