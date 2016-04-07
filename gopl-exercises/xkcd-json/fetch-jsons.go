// Fetches a sequence of XKCD JSON files from xckd.com and places them in a
// designated directory.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strconv"
	"time"
)

var baseurl = flag.String("baseurl", "http://xkcd.com", "XKCD base fetching URL")
var startnum = flag.Int("startnum", 400, "fetch episodes starting with this number")
var endnum = flag.Int("endnum", 420, "fetch episodes ending with this number")
var infopath = flag.String("infopath", "info.0.json", "info path for JSON files")
var targetdir = flag.String("targetdir", "/tmp", "target directory for JSON files")

func main() {
	flag.Parse()

	for i := *startnum; i <= *endnum; i++ {
		fmt.Println(i)
		url := *baseurl + "/" + strconv.Itoa(i) + "/" + *infopath
		fmt.Println(url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Printf("Get %s... status = %v\n", url, resp.Status)

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("  error reading body: %v\n", err)
			continue
		}
		fmt.Printf("  read %d bytes\n", len(body))

		target_filename := path.Join(*targetdir, strconv.Itoa(i)+".json")
		err = ioutil.WriteFile(target_filename, body, 0644)
		if err != nil {
			fmt.Printf("failed WriteFile: %v\n", err)
			continue
		}
		fmt.Printf("Wrote %s\n", target_filename)

		time.Sleep(time.Millisecond * 500)
	}
}
