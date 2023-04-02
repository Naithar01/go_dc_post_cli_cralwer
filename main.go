package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Post struct {
	Id     string
	Title  string
	Writer string
}

var (
	URL string = "https://gall.dcinside.com/board/lists/?id=programming&page=%d"
)

func CrawlerSite(page int, docChan chan *goquery.Document) error {
	url := fmt.Sprintf(URL, page)
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return err
	}

	docChan <- doc

	return nil
}

func Crawler(posts *[]Post, page_count int) {
	docChan := make(chan *goquery.Document)

	go func() {
		if page_count == 1 {
			for page := 1; page <= 9; page++ {
				CrawlerSite(page, docChan)
				fmt.Print(page)
			}
		} else {
			for page := (page_count - 1) * 10; page <= page_count*10-1; page++ {
				CrawlerSite(page, docChan)
				fmt.Print(page)
			}
		}
		close(docChan)
	}()

	for doc := range docChan {
		doc.Find("tr.ub-content").Each(func(i int, s *goquery.Selection) {
			id := s.Find("td.gall_num").Text()
			title := s.Find("td.gall_tit > a").Text()
			writer := s.Find("td.gall_writer").AttrOr("data-nick", "ㅇㅇ")
			post := Post{Id: id, Title: title, Writer: writer}
			*posts = append(*posts, post)
		})
	}
}

func main() {
	// 실행속도를 확인하기 위한
	start := time.Now()

	// 한 페이지에 게시글은 총 51개
	var posts []Post

	page_count := 2
	for i := 1; i <= page_count; i++ {
		Crawler(&posts, i)
	}

	fmt.Println("    ", len(posts))

	fmt.Println(time.Now().Sub(start).Seconds(), "/s")
}
