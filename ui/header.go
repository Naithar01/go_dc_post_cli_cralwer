package ui

import (
	"strconv"

	"github.com/nsf/termbox-go"
)

type Header_Info struct {
	Site_Page       string
	Now_Page        string
	Max_Page        string
	X               int
	Color           termbox.Attribute
	BackgroundColor termbox.Attribute
}

func (h *Header_Info) WriteSitePage() {
	h.WriteBanner("Site Page:")
	for _, header_info_site_page := range h.Site_Page {
		termbox.SetChar(h.X, 1, header_info_site_page)
		termbox.SetFg(h.X, 1, termbox.ColorBlue)
		h.X++
	}
}

func (h *Header_Info) WriteNowPage() {
	termbox.SetChar(h.X, 1, ' ')
	termbox.SetBg(h.X, 1, termbox.ColorCyan)
	h.X++

	h.WriteBanner("| Now Page:")
	for _, header_info_now_page := range h.Now_Page {
		termbox.SetChar(h.X, 1, header_info_now_page)
		termbox.SetFg(h.X, 1, termbox.ColorBlue)
		h.X++
	}
}

func (h *Header_Info) WriteMaxPage() {
	termbox.SetChar(h.X, 1, ' ')
	termbox.SetBg(h.X, 1, termbox.ColorCyan)
	h.X++

	h.WriteBanner("| Max Page:")
	for _, header_info_max_page := range h.Max_Page {
		termbox.SetChar(h.X, 1, header_info_max_page)
		termbox.SetFg(h.X, 1, termbox.ColorBlue)
		h.X++
	}
}

func (h *Header_Info) WriteBanner(banner string) {
	for _, b := range banner {
		termbox.SetChar(h.X, 1, b)
		termbox.SetBg(h.X, 1, h.BackgroundColor)
		h.X++
	}
}

func (h *Header_Info) WriteHeaderInfo() {
	h.WriteSitePage()
	h.WriteNowPage()
	h.WriteMaxPage()
}

func (h *Header_Info) GetNowPage() int {
	now_page, _ := strconv.Atoi(h.Now_Page)

	return now_page
}
