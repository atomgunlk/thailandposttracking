package thailandposttracking_test

import (
	"os"
	"testing"
	"time"

	"github.com/atomgunlk/thailandposttracking"
)

const (
	token = "OHB:FdY3EsU^F6F^EAU5MQJmP2XUKQB7O&Q-VcYOGwE+DwKBM?T=K+CSMBHdJ=DpMSMoHhPHQjB0YYZbYFCxBHK#T$H!PQMpQWAI"
)

var thpostClient *thailandposttracking.Client

func setup() {
	// Setup service
	thpostClient = thailandposttracking.New(token)
	thpostClient.SetApiToken(thailandposttracking.Token{
		Expire: time.Now().AddDate(0, 0, 1),
		Token:  "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpc3MiOiJzZWN1cmUtYXBpIiwiYXVkIjoic2VjdXJlLWFwcCIsInN1YiI6IkF1dGhvcml6YXRpb24iLCJleHAiOjE2ODE2NDg2NTcsInJvbCI6WyJST0xFX1VTRVIiXSwiZCpzaWciOnsicCI6InpXNzB4IiwicyI6bnVsbCwidSI6ImY5ZmIwY2EzM2IyNTQxODJkMWYzOWFiYjliNTEzNmJlIiwiZiI6InhzeiM5In19.Qp6G7AgwhPzo6qaEkh4XaUWeJp4GTavWg3v2KzqkL1UdQp8-ad_LqZ8ODGeuWM6Q__bAlPHIz0GYFe9eBpDuGA",
	})
}
func shutdown() {
	// close client
}
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
