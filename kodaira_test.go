package og_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/katera/og"
)

func TestGetOpenGraphFromUrl(t *testing.T) {
	// u := "https://medium.com/"
	// u := "https://techcrunch.com/2018/03/25/towards-a-world-without-facebook"
	u := "https://www.theguardian.com/technology/2018/mar/25/facebook-logs-texts-and-calls-users-find-as-they-delete-accounts-cambridge-analytica"
	// u := "https://www.youtube.com/watch?v=DelhLppPSxY&list=RDDelhLppPSxY"
	// u := "https://github.com/domenic/package-name-maps"
	// u := "https://hackernoon.com/golang-clean-archithecture-efd6d7c43047"
	// u := "https://www.theverge.com/2018/3/25/17160944/facebook-call-history-sms-data-collection-android"
	// u := "https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html"
	// u := "https://blog.kissmetrics.com/open-graph-meta-tags/"
	// u := "https://www.technologyreview.com/s/610576/how-to-manipulate-facebook-and-twitter-instead-of-letting-them-manipulate-you/"
	// u := "https://www.wired.com/story/europes-new-privacy-law-will-change-the-web-and-more/"
	// u := "https://www.economist.com/news/leaders/21739151-how-it-and-wider-industry-should-respond-facebook-faces-reputational-meltdown"
	// u := "https://www.nytimes.com/2018/03/21/technology/facebook-zucktown-willow-village.html"
	res, err := og.GetOpenGraphFromUrl(u)
	if err != nil {
		t.Error("Expected not Err but got Err: ", err)
	}

	jybt, _ := json.Marshal(res)

	fmt.Println(string(jybt))
}
