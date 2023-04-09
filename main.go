package main

import (
	"log"
	"math"
	"strconv"

	"github.com/Naithar01/dc_cli_crawler/crawler"
	"github.com/Naithar01/dc_cli_crawler/ui"
	"github.com/nsf/termbox-go"
)

func InitTermBox() error {
	err := termbox.Init()
	if err != nil {
		return err
	}

	return nil
}

func InitWritePostHeader() *ui.Header_Info {
	return &ui.Header_Info{
		Site_Page: "1",
		Now_Page:  "1",
		Max_Page:  "",
		X:         0,
	}
}

func InitWritePost() *ui.Post_Info {
	return &ui.Post_Info{Posts: []crawler.Post{}}
}

func InitApp() (*ui.Header_Info, *ui.Post_Info) {
	err := InitTermBox()
	header_info := InitWritePostHeader()
	posts_info := InitWritePost()

	if err != nil {
		log.Panic(err.Error())
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Draw Header
	header_info.Color = termbox.ColorRed

	// Get now page (string -> integer)
	now_page, _ := strconv.Atoi(header_info.Now_Page)
	site_page, _ := strconv.Atoi(header_info.Site_Page)

	// Get Posts Data ( Crawler Data )
	posts_info.GetPosts(site_page)
	posts_info.Post_Length = int(math.Ceil(float64(float32(len(posts_info.Posts)) / float32(7.0))))

	// Draw Posts
	// Set Post Lenth ( Site Post Lenth (0 ~ 51) )
	posts_info.WritePosts(termbox.ColorWhite, now_page)

	// Post Length: 51, Cloude: 2 || 51 - 2 = 49; => 7 * 7 == 49
	header_info.Max_Page = strconv.Itoa(posts_info.Post_Length)

	header_info.WriteHeaderInfo()

	return header_info, posts_info
}

func main() {
	header_info, posts_info := InitApp()

	defer termbox.Close()

	for {
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			// Exit
			if ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC {
				return
			}

			// Change Page
			if ev.Ch == 'q' || ev.Ch == 'e' {
				termWidth, _ := termbox.Size()

				for col := 0; col < termWidth; col++ {
					termbox.SetCell(col, 0, ' ', termbox.ColorDefault, termbox.ColorDefault)
				}

				header_info.X = 0
				now_page, _ := strconv.Atoi(header_info.Now_Page)

				if ev.Ch == 'q' {
					if now_page != 1 {
						header_info.Now_Page = strconv.Itoa(now_page - 1)
					} else {
						continue
					}
				} else if ev.Ch == 'e' {
					max_page, _ := strconv.Atoi(header_info.Max_Page)
					if now_page < max_page {
						header_info.Now_Page = strconv.Itoa(now_page + 1)
					} else {
						continue
					}
				}
				now_page, _ = strconv.Atoi(header_info.Now_Page)
				header_info.WriteHeaderInfo()
				posts_info.WritePosts(termbox.ColorWhite, now_page)
			}
		}
	}
}
