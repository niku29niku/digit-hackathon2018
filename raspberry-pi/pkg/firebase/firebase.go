package firebase

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	firego "gopkg.in/zabawaba99/firego.v1"
)

// NewFirebaseClient create firebase client instance
func NewFirebaseClient() (*firego.Firebase, error) {
	credPath := filepath.Join(os.Getenv("HOME"), "digit-hackathon2018-niku29-firebase.json")
	cred, err := ioutil.ReadFile(credPath)
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON(cred, "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/firebase.database")
	firebase := firego.New("https://digit-hackathon2018-niku29.firebaseio.com", conf.Client(oauth2.NoContext))
	return firebase, nil
}
