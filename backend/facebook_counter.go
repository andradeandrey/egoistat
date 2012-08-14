package egoistat

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type facebookFeed struct {
	Shares int
}

type FacebookCounter struct{}

func (c *FacebookCounter) urlFor(url string) string {
	return fmt.Sprintf("http://graph.facebook.com/?ids=%s", url)
}

func (c *FacebookCounter) Count(r *Request) (count int) {
	var resp *http.Response
	var dec *json.Decoder
	var err error
	var data map[string]facebookFeed
	var url = r.Url()

	if resp, err = http.Get(c.urlFor(url)); err != nil {
		return 0
	}

	dec = json.NewDecoder(resp.Body)
	if err = dec.Decode(&data); err != nil {
		return 0
	}

	if dataForUrl, ok := data[url]; ok {
		count = dataForUrl.Shares
	}
	return
}

func init() {
	counters["facebook"] = new(FacebookCounter)
}
