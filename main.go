package main

import (
	"log"
	"strconv"

	"github.com/Naithar01/dc_cli_crawler/crawler"
	"github.com/nsf/termbox-go"
)

type Post_Info struct {
	posts []crawler.Post
}

func (p *Post_Info) WritePosts(color termbox.Attribute) {
	p.posts = crawler.Page()[:5] // Test...

	for index, post := range p.posts {
		x := 0
		for _, ID := range post.Id {
			termbox.SetCell(x, index+1, ID, color, termbox.ColorDefault)
			x++
		}

		termbox.SetCell(x, index+1, ' ', color, termbox.ColorDefault)
		x++

		for _, TITLE := range post.Title {
			termbox.SetCell(x, index+1, TITLE, color, termbox.ColorDefault)
			x++
		}
	}
}

type Header_Info struct {
	Site_Page string
	Now_Page  string
	Max_Page  string
	x         int
	color     termbox.Attribute
}

func (h *Header_Info) WriteSitePage() {
	h.WriteBanner("Site Page:")
	for _, header_info_site_page := range h.Site_Page {
		termbox.SetCell(h.x, 0, header_info_site_page, h.color, termbox.ColorDefault)
		h.x++
	}
}

func (h *Header_Info) WriteNowPage() {
	termbox.SetCell(h.x, 0, ' ', h.color, termbox.ColorDefault)
	h.x++

	h.WriteBanner("| Now Page:")
	for _, header_info_now_page := range h.Now_Page {
		termbox.SetCell(h.x, 0, header_info_now_page, h.color, termbox.ColorDefault)
		h.x++
	}
}

func (h *Header_Info) WriteMaxPage() {
	termbox.SetCell(h.x, 0, ' ', h.color, termbox.ColorDefault)
	h.x++

	h.WriteBanner("| Max Page:")
	for _, header_info_max_page := range h.Max_Page {
		termbox.SetCell(h.x, 0, header_info_max_page, h.color, termbox.ColorDefault)
		h.x++
	}
}

func (h *Header_Info) WriteBanner(banner string) {
	for _, b := range banner {
		termbox.SetCell(h.x, 0, b, h.color, termbox.ColorDefault)
		h.x++
	}
}

func (h *Header_Info) WriteHeaderInfo() {
	h.WriteSitePage()
	h.WriteNowPage()
	h.WriteMaxPage()
}

func InitTermBox() error {
	err := termbox.Init()
	if err != nil {
		return err
	}

	return nil
}

func InitWritePostHeader() *Header_Info {
	return &Header_Info{
		Site_Page: "1",
		Now_Page:  "2",
		Max_Page:  "31",
		x:         0,
	}
}

func InitWritePost() *Post_Info {
	return &Post_Info{posts: []crawler.Post{}}
}

func main() {
	err := InitTermBox()

	if err != nil {
		log.Panic(err.Error())
	}

	defer termbox.Close()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	header_info := InitWritePostHeader()
	header_info.color = termbox.ColorRed

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

				header_info.x = 0
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
