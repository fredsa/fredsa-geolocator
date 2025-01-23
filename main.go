package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/appengine/v2"
)

// https://cloud.google.com/appengine/docs/standard/go/runtime#environment_variables
const PORT = "PORT"

const GOOGLE_CLOUD_PROJECT = "GOOGLE_CLOUD_PROJECT" // The Google Cloud project ID associated with your application.
const GAE_APPLICATION = "GAE_APPLICATION"           // App id, with prefix.
const GAE_ENV = "GAE_ENV"                           // `standard` in production.
const GAE_RUNTIME = "GAE_RUNTIME"                   // Runtime in `app.yaml`.
const GAE_VERSION = "GAE_VERSION"                   // App version.
const DUMMY_APP_ID = "my-app-id"

// X-Appengine-Country: US
const X_APPENGINE_COUNTRY = "X-AppEngine-Country"

// X-Appengine-Region: ca
const X_APPENGINE_REGION = "X-AppEngine-Region"

// X-Appengine-City: sunnyvale
const X_APPENGINE_CITY = "X-AppEngine-City"

// X-Appengine-Citylatlong: 37.368830,-122.036350
const X_APPENGINE_CITYLATLONG = "X-AppEngine-CityLatLong"

// X-Appengine-User-Ip: 2620:15c:2d1:206:969:d104:6bee:4832
const X_APPENGINE_USER_IP = "X-Appengine-User-Ip"

var headers = []string{
	X_APPENGINE_COUNTRY,
	X_APPENGINE_REGION,
	X_APPENGINE_CITY,
	X_APPENGINE_CITYLATLONG,
	X_APPENGINE_USER_IP,
}

func init() {
	// Register handlers in init() per `appengine.Main()` documentation.
	http.HandleFunc("/", indexHandler)
}

func main() {
	if isDev() {
		_ = os.Setenv(X_APPENGINE_COUNTRY, "US")
		_ = os.Setenv(X_APPENGINE_REGION, "ca")
		_ = os.Setenv(X_APPENGINE_CITY, "sunnyvale")
		_ = os.Setenv(X_APPENGINE_CITYLATLONG, "37.368830,-122.036350")
		_ = os.Setenv(X_APPENGINE_USER_IP, "2620:15c:2d1:206:969:d104:6bee:4832")

		_ = os.Setenv(GAE_APPLICATION, DUMMY_APP_ID)
		_ = os.Setenv(GAE_RUNTIME, "go123456")
		_ = os.Setenv(GAE_VERSION, "my-version")
		_ = os.Setenv(GAE_ENV, "standard")
		_ = os.Setenv(PORT, "4200")

		log.Printf("appengine.Main() will listen: %s", defaultVersionOrigin())
	}

	// Standard App Engine APIs require `appengine.Main` to have been called.
	appengine.Main()
}

func defaultVersionOrigin() string {
	if isDev() {
		return "http://localhost:" + os.Getenv(PORT)
	} else {
		return fmt.Sprintf("https://%s.appspot.com", projectID())
	}
}

func projectID() string {
	return os.Getenv(GOOGLE_CLOUD_PROJECT)
}

func isDev() bool {
	appid := os.Getenv(GAE_APPLICATION)
	return appid == "" || appid == DUMMY_APP_ID
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// App Engine context for the in-flight HTTP request.
	// ctx := appengine.NewContext(r)

	w.Header().Set("Content-Type", "text/plain")

	for _, header := range headers {
		fmt.Fprintf(w, "%s: %s\n", header, r.Header.Get(header))
	}
}
