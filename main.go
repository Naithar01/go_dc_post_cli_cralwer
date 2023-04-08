package main

import (
	"log"
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
		Now_Page:  "2",
		Max_Page:  "31",
		X:         0,
	}
}

func InitWritePost() *ui.Post_Info {
	return &ui.Post_Info{Posts: []crawler.Post{}}
}

func main() {
	err := InitTermBox()

	if err != nil {
		log.Panic(err.Error())
	}

	defer termbox.Close()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	header_info := InitWritePostHeader()
	header_info.Color = termbox.ColorRed
	header_info.WriteHeaderInfo()

	posts_info := InitWritePost()
	posts_info.WritePosts(termbox.ColorWhite)

	for {
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC {
				return
			}

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
						header_info.WriteHeaderInfo()
					} else {
						header_info.WriteHeaderInfo()
					}
				} else if ev.Ch == 'e' {
					max_page, _ := strconv.Atoi(header_info.Max_Page)
					if now_page < max_page {
						header_info.Now_Page = strconv.Itoa(now_page + 1)
						header_info.WriteHeaderInfo()
					} else {
						header_info.WriteHeaderInfo()
					}
				}
			}

		}
	}
}
