package og

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"

	"golang.org/x/net/html"
)

type Article struct {
	PublishedTime  *time.Time `json:"published_time"`
	ModifiedTime   *time.Time `json:"modified_time"`
	ExpirationTime *time.Time `json:"expiration_time"`
	Authors        []*Author  `json:"authors"`
	Section        string     `json:"section"`
	Tags           []string   `json:"tags"`
}

type Author struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
}

type Image struct {
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	URL       string `json:"url"`
	SecureUrl string `json:"secure_url"`
	MimeType  string `json:"type"`
	Alt       string `json:"alt"`
}
type Video struct {
	Width     int      `json:"width"`
	Height    int      `json:"height"`
	URL       string   `json:"url"`
	SecureUrl string   `json:"secure_url"`
	MimeType  string   `json:"type"`
	Tag       []string `json:"tags"`
}

type Audio struct {
	Url       string `json:"url"`
	SecureUrl string `json:"secure_url"`
	MimeType  string `json:"type"`
}

type TwitterCard struct {
	Title       string        `json:"title"`
	Image       string        `json:"image"`
	ImageSource string        `json:"image_src"`
	ImageAlt    string        `json:"image_alt"`
	Url         string        `json:"url"`
	Card        string        `json:"card"`
	Site        string        `json:"site"`
	SiteId      string        `json:"site_id"`
	Creator     string        `json:"creator"`
	CreatorId   string        `json:"creator_id"`
	Description string        `json:"description"`
	Player      *Player       `json:"player"`
	Device      []*DeviceCard `json:"device"`
}

type Player struct {
	Url    string `json:"player"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Stream string `json:"stream"`
}

type DeviceCard struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"` //IPHONE, IPAD , GOOGLEPLAY
	Url  string `json:"url"`
}

type MetaInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Keywords    string `json:"keywords"`
	Author      string `json:"author"`
}

type OpenGraph struct {
	MetaInfo    *MetaInfo         `json:"meta_info"`
	Title       string            `json:"title"`
	Type        string            `json:"type"`
	Url         string            `json:"url"`
	Site        string            `json:"site"`
	SiteName    string            `json:"site_name"`
	Description string            `json:"description"`
	Locale      string            `json:"locale"`
	Article     *Article          `json:"article"`
	Images      []*Image          `json:"images"`
	Videos      []*Video          `json:"videos"`
	Audios      []*Audio          `json:"audios"`
	Twitter     *TwitterCard      `json:"twitter_card"`
	AlDevices   []*DeviceCard     `json:"device"`
	Others      map[string]string `json:"others"`
}

func isUrlValid(u string) bool {
	_, err := url.Parse(u)
	if err != nil {
		return false
	}
	return true
}

func GetOpenGraphFromUrl(u string) (*OpenGraph, error) {
	if !isUrlValid(u) {
		return nil, errors.New("Given Url is not valid")
	}

	return extractFromUrl(context.Background(), u)
}

func GetOpenGraphFromUrlContext(ctx context.Context, u string) (*OpenGraph, error) {
	if !isUrlValid(u) {
		return nil, errors.New("Given Url is not valid")
	}
	if ctx == nil {
		return nil, errors.New("Context is nil")
	}

	return extractFromUrl(ctx, u)
}

type meta struct {
	Name    string
	Content string
}

func extractFromUrl(ctx context.Context, u string) (*OpenGraph, error) {
	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req.WithContext(ctx)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return extractOpenGraph(resp.Body)
}

func extractOpenGraph(data io.Reader) (*OpenGraph, error) {

	htmlTokenized := html.NewTokenizer(data)
	metaContent := []meta{}
	running := true
	for running {
		token := htmlTokenized.Next()

		switch token {
		case html.ErrorToken:
			if htmlTokenized.Err() == io.EOF {
				running = false
				break
			}
			return nil, htmlTokenized.Err()
		case html.StartTagToken, html.EndTagToken, html.SelfClosingTagToken:
			name, hasAttributes := htmlTokenized.TagName()

			if string(name) != "meta" || !hasAttributes {
				continue
			}

			if string(name) == "body" {
				running = false // <body></body> skipped. We just need <head>
				break
			}

			m := meta{}
			for hasAttributes {
				key, val, moreAttributes := htmlTokenized.TagAttr()
				hasAttributes = moreAttributes
				keyString := string(key)
				switch keyString {
				case "name", "property":
					m.Name = string(val)
				case "content":
					m.Content = string(val)

				}

			}
			metaContent = append(metaContent, m)

		}

	}

	return marshallFromMap(metaContent), nil
}

//Reff: http://ogp.me/
func marshallFromMap(data []meta) *OpenGraph {
	og := &OpenGraph{}
	imgs := []*Image{}
	videos := []*Video{}
	info := &MetaInfo{}
	article := &Article{}
	articleTags := []string{}
	authors := []*Author{}
	twitterCard := &TwitterCard{}
	player := &Player{}
	iphone := &DeviceCard{Type: "iphone"}
	android := &DeviceCard{Type: "android"}
	ipad := &DeviceCard{Type: "ipad"}

	alIos := &DeviceCard{Type: "ios"}
	alIphone := &DeviceCard{Type: "iphone"}
	alIpad := &DeviceCard{Type: "ipad"}
	alAndroid := &DeviceCard{Type: "android"}

	otherMetas := map[string]string{}
	for index := 0; index < len(data); index++ {
		m := data[index]

		switch m.Name {
		case "author":
			info.Author = m.Content
		case "title":
			info.Title = m.Content
		case "description":
			info.Description = m.Content
		case "keywords":
			info.Keywords = m.Content
		case "og:title":
			if og.Title == "" {
				og.Title = m.Content
			}
		case "og:description":
			if og.Description == "" {
				og.Description = m.Content
			}
		case "og:type":
			og.Type = m.Content
		case "og:url":
			og.Url = m.Content
		case "og:locale":
			og.Locale = m.Content
		case "og:site":
			og.Site = m.Content
		case "og:site_name":
			og.SiteName = m.Content
		case "og:image":
			img := &Image{}
			img.URL = m.Content
			imgs = append(imgs, img)
		case "og:image:url":
			if len(imgs) > 0 && imgs[len(imgs)-1].URL == "" {
				imgs[len(imgs)-1].URL = m.Content
			}
		case "og:image:width":
			c := removeWhiteSpace(m.Content)
			w, _ := strconv.ParseInt(c, 10, 64)
			if len(imgs) > 0 {
				imgs[len(imgs)-1].Width = int(w)
			}
		case "og:image:height":
			c := removeWhiteSpace(m.Content)
			h, _ := strconv.ParseInt(c, 10, 64)
			if len(imgs) > 0 {
				imgs[len(imgs)-1].Height = int(h)
			}
		case "og:image:type":
			if len(imgs) > 0 {
				imgs[len(imgs)-1].MimeType = m.Content
			}
		case "og:image:secure_url":
			if len(imgs) > 0 {
				imgs[len(imgs)-1].SecureUrl = m.Content
			}
		case "og:image:alt":
			if len(imgs) > 0 {
				imgs[len(imgs)-1].Alt = m.Content
			}
		case "og:video":
			v := &Video{}
			v.URL = m.Content
			videos = append(videos, v)
		case "og:video:tag":
			if len(videos) == 0 {
				v := &Video{}
				videos = append(videos, v)
			}
			videos[len(videos)-1].Tag = append(videos[len(videos)-1].Tag, m.Content)

		case "og:video:url":
			if len(videos) > 0 && videos[len(videos)-1].URL == "" {
				videos[len(videos)-1].URL = m.Content
			}
		case "og:video:secure_url":
			if len(videos) > 0 {
				videos[len(videos)-1].SecureUrl = m.Content
			}
		case "og:video:type":
			if len(videos) > 0 {
				videos[len(videos)-1].MimeType = m.Content
			}
		case "og:video:width":
			if len(videos) > 0 {
				c := removeWhiteSpace(m.Content)
				w, _ := strconv.ParseInt(c, 10, 64)

				videos[len(videos)-1].Width = int(w)
			}
		case "og:video:height":
			if len(videos) > 0 {
				c := removeWhiteSpace(m.Content)
				h, _ := strconv.ParseInt(c, 10, 64)
				videos[len(videos)-1].Height = int(h)
			}
		case "article:tag":
			articleTags = append(articleTags, m.Content)
		case "article:published_time":
			published, err := time.Parse(time.RFC3339, m.Content)
			if err == nil {
				article.PublishedTime = &published
			}
		case "article:modified_time":
			modified, err := time.Parse(time.RFC3339, m.Content)
			if err == nil {
				article.ModifiedTime = &modified
			}
		case "article:expiration_time":
			expiration, err := time.Parse(time.RFC3339, m.Content)
			if err == nil {
				article.ExpirationTime = &expiration
			}
		case "article:section":
			article.Section = m.Content
		case "article:author":
			a := &Author{FirstName: m.Content}
			authors = append(authors, a)
		case "article:author:firstname":
			if len(authors) == 0 {
				a := &Author{}
				authors = append(authors, a)
			}
			authors[len(authors)-1].FirstName = m.Content
		case "article:author:lastname":
			if len(authors) == 0 {
				a := &Author{}
				authors = append(authors, a)
			}
			authors[len(authors)-1].LastName = m.Content
		case "article:author:username":
			if len(authors) == 0 {
				a := &Author{}
				authors = append(authors, a)
			}
			authors[len(authors)-1].Username = m.Content
		case "article:author:gender":
			if len(authors) == 0 {
				a := &Author{}
				authors = append(authors, a)
			}
			authors[len(authors)-1].Gender = m.Content
		case "twitter:site":
			twitterCard.Site = m.Content
		case "twitter:site:id":
			twitterCard.SiteId = m.Content
		case "twitter:url":
			twitterCard.Url = m.Content
		case "twitter:title":
			twitterCard.Title = m.Content
		case "twitter:description":
			twitterCard.Description = m.Content
		case "twitter:image":
			twitterCard.Image = m.Content
		case "twitter:image:alt":
			twitterCard.ImageAlt = m.Content
		case "twitter:image:src":
			twitterCard.ImageAlt = m.Content
		case "twitter:card":
			twitterCard.Card = m.Content
		case "twitter:creator":
			twitterCard.Creator = m.Content
		case "twitter:creator:id":
			twitterCard.CreatorId = m.Content
		case "twitter:player":
			player.Url = m.Content
		case "twitter:player:width":
			c := removeWhiteSpace(m.Content)
			w, _ := strconv.ParseInt(c, 10, 64)
			player.Width = int(w)
		case "twitter:player:height":
			c := removeWhiteSpace(m.Content)
			h, _ := strconv.ParseInt(c, 10, 64)
			player.Height = int(h)

		case "twitter:app:name:iphone":
			iphone.Name = m.Content
		case "twitter:app:id:iphone":
			iphone.Id = m.Content
		case "twitter:app:url:iphone":
			iphone.Url = m.Content
		case "twitter:app:name:ipad":
			ipad.Name = m.Content
		case "twitter:app:id:ipad":
			ipad.Id = m.Content
		case "twitter:app:url:ipad":
			ipad.Url = m.Content
		case "twitter:app:name:googleplay":
			android.Name = m.Content
		case "twitter:app:id:googleplay":
			android.Id = m.Content
		case "twitter:app:url:googleplay":
			android.Url = m.Content
		case "al:android:app_name":
			alAndroid.Name = m.Content
		case "al:android:package":
			alAndroid.Id = m.Content
		case "al:android:url":
			alAndroid.Url = m.Content
		case "al:ios:app_store_id":
			alIos.Id = m.Content
		case "al:ios:app_name":
			alIos.Name = m.Content
		case "al:ios:url":
			alIos.Url = m.Content
		case "al:ipad:app_store_id":
			alIpad.Id = m.Content
		case "al:ipad:app_name":
			alIpad.Name = m.Content
		case "al:ipad:url":
			alIpad.Url = m.Content
		case "al:iphone:app_store_id":
			alIphone.Id = m.Content
		case "al:iphone:app_name":
			alIphone.Name = m.Content
		case "al:iphone:url":
			alIphone.Url = m.Content
		default:
			if m.Name == "" {
				continue
			}
			otherMetas[m.Name] = m.Content
		}
	}
	twitterCard.Device = []*DeviceCard{iphone, ipad, android}
	twitterCard.Player = player
	article.Tags = articleTags
	article.Authors = authors
	og.AlDevices = []*DeviceCard{alAndroid, alIos, alIpad, alIphone}
	og.Twitter = twitterCard
	og.Images = imgs
	og.Videos = videos
	og.MetaInfo = info
	og.Article = article
	og.Others = otherMetas

	return og
}

func removeWhiteSpace(s string) string {
	res := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, s)

	return res
}

func GetOpenGraphFromHtml(h string) (*OpenGraph, error) {
	data := strings.NewReader(h)
	return extractOpenGraph(data)
}
