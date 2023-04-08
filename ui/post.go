package ui

import (
	"github.com/Naithar01/dc_cli_crawler/crawler"
	"github.com/nsf/termbox-go"
)

type Post_Info struct {
	Posts       []crawler.Post
	Post_Length int
}

func (p *Post_Info) WritePosts(color termbox.Attribute, page int) {
	p.Posts = crawler.Page() // Test...

	for index, post := range p.Posts {
		x := 0
		for _, ID := range post.Id {
			termbox.SetCell(x, index+2, ID, color, termbox.ColorDefault)
			x++
		}

		termbox.SetCell(x, index+2, ' ', color, termbox.ColorDefault)
		x++

		for _, TITLE := range post.Title {
			termbox.SetCell(x, index+2, TITLE, color, termbox.ColorDefault)
			x++
		}
	}
}
