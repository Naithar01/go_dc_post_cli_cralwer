package ui

import (
	"math"

	"github.com/Naithar01/dc_cli_crawler/crawler"
	"github.com/nsf/termbox-go"
)

type Post_Info struct {
	Posts       []crawler.Post
	Post_Length int
}

func (p *Post_Info) GetPosts(site_page int) {
	p.Posts = crawler.Page(site_page)
	p.Post_Length = int(math.Ceil(float64(float32(len(p.Posts)) / float32(7.0))))
}

func (p *Post_Info) WritePosts(color termbox.Attribute, page int) {
	for index, post := range p.Posts[(page-1)*7 : page*7] {
		x := 2
		for _, ID := range post.Id {
			termbox.SetCell(x, index+2, ID, color, termbox.ColorDefault)
			x++
		}

		x += 5

		termbox.SetCell(x, index+2, ' ', color, termbox.ColorDefault)
		x++

		for _, TITLE := range post.Title {
			termbox.SetCell(x, index+2, TITLE, color, termbox.ColorDefault)
			x++
		}
	}
}
