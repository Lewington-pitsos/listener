package assist

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

func PathToPackage() string {
	return os.Getenv("GOPATH") + "/src/github.com/lewington/listener"
}

func PathToAsset() string {
	return PathToPackage() + "/public/static/assets/"
}

func Timestamp() time.Time {
	return time.Now().In(TimeLocation)
}

var AEST = time.FixedZone("AEST", -60*60*4)
var TimeLocation, _ = time.LoadLocation("Australia/Melbourne")
var BetfairLocation, _ = time.LoadLocation("GMT")

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Panicf(message string, args ...interface{}) {
	panic(fmt.Sprintf(message, args...))
}

func StrictBytes(body io.ReadCloser) []byte {
	defer body.Close()
	bytes, err := ioutil.ReadAll(body)
	Check(err)

	return bytes
}
