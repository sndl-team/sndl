package lib

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// This struct may honestly not be necessary, we need to do some discovery work
// on what the best practices may be and what we can do to minimize our code
// footprint.
type SenturionClient struct {
	Client http.Client
}

// Golang's OOP support is minimal and thus I am adopting nim's practice of using
// custom initializers that are functions. I believe this is a go best practice but
// I am not sure of that fact.
func NewSenturionClient() SenturionClient {

	// create a jar object that is used to do 90% of the heavy lifting when it comes to
	// http headers. The public suffix list is a security option and should default
	// to being present.
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}

	// Return the initialized object
	return SenturionClient{Client: http.Client{Jar: jar}}
}

// This signature is how "methods" work in go, under the hood there is not much
// differnce between:
// 		foo(object, "bar")
// and
// 		object.foo("bar")
// ---
// Regardless this function initializes some of the headers by performing a
// get request to the index url of the server and retrieving the default headers
func (client SenturionClient) Init() {

	// We simply need to make sure the request didnt raise an error
	// the error will not be nil if something bad happened
	// the error will nbe nil even if we get a non-2xx status code
	// thus we check the code here.
	resp, err := client.Client.Get("https://senturion.to")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Request to initialize status code != 200")
	}
}

// This method must be called after Init above and will fail otherwise.
// We could chain the calls but we would need to tweak the function names
func (client SenturionClient) Login(username string, password string) {

	// the login endpoint takes formdata and this is how that is done.
	formData := url.Values{
		"username": {username},
		"password": {password},
		"remember": {"1"}}

	// Post the form data to the endpoint
	resp, err := client.Client.PostForm(
		"https://senturion.to/session/userlogin",
		formData)

	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("Request to Login status code != 200")
	}

	// Strangely, this endpoint always returns a 200 status code.
	// The only way to tell it failed is by checking the body for a given
	// string. This is obnoxious and slows down our program a tad but I dont
	// know what else to do.
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Checking for that failed string. blah.
	if strings.Contains(string(body),
		"You are now required to login before viewing the website.") {

		log.Fatal("Username or password incorrect")
	}
}
