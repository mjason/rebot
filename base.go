package rebot

import (
	"code.google.com/p/mahonia"
	"exp/html"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
)

func GbkToUtf8(s string) string {
	decoder := mahonia.NewDecoder("gb18030")
	return Clear_space(decoder.ConvertString(s))
}

func Clear_space(s string) string {
	return strings.TrimSpace(s)
}

func PostForm(uri string, data url.Values) (d *goquery.Document, e error) {
	resp, e := http.PostForm(uri, data)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()

	root, e := html.Parse(resp.Body)
	if e != nil {
		return nil, e
	}

	doc := goquery.NewDocumentFromNode(root)
	return doc, nil
}
