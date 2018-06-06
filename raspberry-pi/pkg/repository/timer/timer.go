package timer

import (
	"time"

	"gopkg.in/zabawaba99/firego.v1"
)

// Repository is interface to save timer information
type Repository interface {
	SetTimer(duration time.Duration) error
	Remove() error
}

func NewFirebaseRepository(firebase *firego.Firebase) Repository {
	return &firebaseTimerRepository{firebase}
}

type firebaseTimerRepository struct {
	firebase *firego.Firebase
}

func (rep *firebaseTimerRepository) SetTimer(duration time.Duration) error {
	d := int(duration / time.Second)
	value := map[string]map[string]int{
		"timer": {
			"duration": d,
		},
	}

	return rep.firebase.Update(value)
}

func (rep *firebaseTimerRepository) Remove() error {
	return rep.firebase.Child("timer").Remove()
}
