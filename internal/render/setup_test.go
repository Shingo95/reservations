package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/elorenzotti/bookings/internal/config"
	"github.com/elorenzotti/bookings/internal/models"
)

var TestApp config.AppConfig
var session *scs.SessionManager

func TestMain(m *testing.M) {
	// what am I goin to put in my session
	gob.Register(models.Reservation{})

	// Change to true when in Production
	TestApp.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	TestApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	TestApp.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	TestApp.Session = session
	app = &TestApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (myW *myWriter) Header() http.Header {
	var h http.Header

	return h
}

func (myW *myWriter) WriteHeader(i int) {

}

func (myW *myWriter) Write(b []byte) (int, error) {
	length := len(b)

	return length, nil
}
