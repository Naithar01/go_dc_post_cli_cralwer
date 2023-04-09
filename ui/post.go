package ui

import (
	"math"

	"github.com/Naithar01/dc_cli_crawler/crawler"
	"github.com/nsf/termbox-go"
)

type Post_Info struct {
	Posts                []crawler.Post
	Post_Length          int
	Post_Line_Space      int
	IDColor              termbox.Attribute
	IDBackgroundColor    termbox.Attribute
	TITLEColor           termbox.Attribute
	TITLEBackgroundColor termbox.Attribute
}

func (p *Post_Info) GetPosts(site_page int) {
	p.Posts = crawler.Page(site_page)
	p.Post_Length = int(math.Ceil(float64(float32(len(p.Posts)) / float32(7.0))))
}

func (p *Post_Info) WritePosts(page int) {
	for index, post := range p.Posts[(page-1)*7 : page*7] {
		x := 2
		for _, ID := range post.Id {
			termbox.SetChar(x, index+p.Post_Line_Space, ID)
			termbox.SetFg(x, index+p.Post_Line_Space, p.IDColor)
			termbox.SetBg(x, index+p.Post_Line_Space, p.IDBackgroundColor)
			x++
		}

		x += 5

		for _, TITLE := range post.Title {
			termbox.SetChar(x, index+p.Post_Line_Space, TITLE)
			termbox.SetFg(x, index+p.Post_Line_Space, p.TITLEColor)
			termbox.SetBg(x, index+p.Post_Line_Space, p.TITLEBackgroundColor)
			x++
		}
	}
}
