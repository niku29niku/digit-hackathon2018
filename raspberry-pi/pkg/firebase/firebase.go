package firebase

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	firego "gopkg.in/zabawaba99/firego.v1"
)

// NewFirebaseClient create firebase client instance
func NewFirebaseClient() (*firego.Firebase, error) {
	credPath := filepath.Join(os.Getenv("HOME"), "digit-hackathon2018-niku29-firebase.json")
	glog.V(2).Infof("credential file path : %s", credPath)
	cred, err := ioutil.ReadFile(credPath)
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON(cred, "https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/firebase.database")
	firebase := firego.New("https://digit-hackathon2018-niku29.firebaseio.com", conf.Client(oauth2.NoContext))
	return firebase, nil
}
