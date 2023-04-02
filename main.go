package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Post struct {
	Id     string
	Title  string
	Writer string
}

var (
	URL   string = "https://gall.dcinside.com/board/lists/?id=programming&page=%d"
	mutex sync.Mutex
)

func RequestCrawlerSite(page int, docChan chan *goquery.Document) error {
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

// 페이지 * 10 페이지의 데이터를 가져오는 함수
// page_count 1이면 1 페이지 ~ 9 페이지의 글을 가져옴
func Crawler_Pages(posts *[]Post, page_count int) {
	docChan := make(chan *goquery.Document, 10)

	go func() {
		var StartPage, EndPage int
		if page_count == 1 {
			StartPage = 1
			EndPage = 9
		} else {
			StartPage = (page_count - 1) * 10
			EndPage = page_count*10 - 1
		}
		for page := StartPage; page <= EndPage; page++ {
			RequestCrawlerSite(page, docChan)
		}
		close(docChan)
	}()

	for doc := range docChan {
		doc.Find("tr.ub-content").Each(func(i int, s *goquery.Selection) {
			id := s.Find("td.gall_num").Text()
			title := s.Find("td.gall_tit > a").Text()
			writer := s.Find("td.gall_writer").AttrOr("data-nick", "ㅇㅇ")
			post := Post{Id: id, Title: title, Writer: writer}
			mutex.Lock()
			*posts = append(*posts, post)
			mutex.Unlock()
		})
	}
}

// 매개변수로 온 페이지의 게시물만 가져오는 함수
// 위의 함수랑 절대 다름
func Crawler_Page(posts *[]Post, page_count int) {
	docChan := make(chan *goquery.Document, 10)

	go func() {
		RequestCrawlerSite(page_count, docChan)
		close(docChan)
	}()

	for doc := range docChan {
		doc.Find("tr.ub-content").Each(func(i int, s *goquery.Selection) {
			id := s.Find("td.gall_num").Text()
			title := s.Find("td.gall_tit > a").Text()
			writer := s.Find("td.gall_writer").AttrOr("data-nick", "ㅇㅇ")
			post := Post{Id: id, Title: title, Writer: writer}
			mutex.Lock()
			*posts = append(*posts, post)
			mutex.Unlock()
		})
	}
}

func Pages() {
	// 실행속도를 확인하기 위한
	start := time.Now()

	// 한 페이지에 게시글은 총 51개
	var posts []Post

	page_count := 2

	for i := 1; i <= page_count; i++ {
		Crawler_Pages(&posts, i)
	}

	log.Println("Crawler_Pages Get Data length:: ", len(posts))

	log.Println(time.Now().Sub(start).Seconds(), "/s")
}

func Page() {
	// 실행속도를 확인하기 위한
	start := time.Now()

	// 한 페이지에 게시글은 총 51개
	var posts []Post

	page_count := 1

	Crawler_Page(&posts, page_count)

	log.Println("Crawler_Page Get Data length:: ", len(posts))

	log.Println(time.Now().Sub(start).Seconds(), "/s")
}

func main() {
	Pages()
	Page()
}
