package valueObject

type Contact struct {
	Email string
}

func MakeContact(email string) *Contact {

	return &Contact{
		Email: email,
	}
}
