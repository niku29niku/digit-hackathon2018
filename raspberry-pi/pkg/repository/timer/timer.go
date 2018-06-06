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

// NewFirebaseRepository create new instance
func NewFirebaseRepository(firebase *firego.Firebase) Repository {
	return &firebaseTimerRepository{firebase}
}

type firebaseTimerRepository struct {
	firebase *firego.Firebase
}

func (rep *firebaseTimerRepository) SetTimer(duration time.Duration) error {
	now := time.Now()
	willEndAt := now.Add(duration)
	formatted := willEndAt.Format(time.RFC3339)
	value := map[string]map[string]interface{}{
		"timer": {
			"willEndAt": formatted,
			"cooking":   true,
		},
	}

	return rep.firebase.Update(value)
}

func (rep *firebaseTimerRepository) Remove() error {
	value := map[string]interface{}{
		"willEndAt": nil,
		"cooking":   false,
	}
	return rep.firebase.Child("timer").Update(value)
}
