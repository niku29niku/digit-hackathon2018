package phone

import (
	firego "gopkg.in/zabawaba99/firego.v1"
)

// Repository is interface to get phone number
type Repository interface {
	PhoneNumbers() ([]string, error)
}

// NewFirebaseRepository create new instance
func NewFirebaseRepository(firebase *firego.Firebase) Repository {
	return &firebasePhoneRepository{firebase: firebase}
}

type firebasePhoneRepository struct {
	firebase *firego.Firebase
}

func (rep *firebasePhoneRepository) PhoneNumbers() (numbers []string, err error) {
	var values map[string]map[string]string
	err = rep.firebase.Child("tel").Value(&values)
	if err != nil {
		return nil, err
	}
	numbers = make([]string, 0)
	for _, col := range values {
		num := col["telNum"]
		numbers = append(numbers, num)
	}
	return
}
