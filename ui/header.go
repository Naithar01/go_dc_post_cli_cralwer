package ui

import "github.com/nsf/termbox-go"

type Header_Info struct {
	Site_Page string
	Now_Page  string
	Max_Page  string
	X         int
	Color     termbox.Attribute
}

func (h *Header_Info) WriteSitePage() {
	h.WriteBanner("Site Page:")
	for _, header_info_site_page := range h.Site_Page {
		termbox.SetCell(h.X, 1, header_info_site_page, h.Color, termbox.ColorDefault)
		h.X++
	}
}

func (h *Header_Info) WriteNowPage() {
	termbox.SetCell(h.X, 1, ' ', h.Color, termbox.ColorDefault)
	h.X++

	h.WriteBanner("| Now Page:")
	for _, header_info_now_page := range h.Now_Page {
		termbox.SetCell(h.X, 1, header_info_now_page, h.Color, termbox.ColorDefault)
		h.X++
	}
}

func (h *Header_Info) WriteMaxPage() {
	termbox.SetCell(h.X, 1, ' ', h.Color, termbox.ColorDefault)
	h.X++

	h.WriteBanner("| Max Page:")
	for _, header_info_max_page := range h.Max_Page {
		termbox.SetCell(h.X, 1, header_info_max_page, h.Color, termbox.ColorDefault)
		h.X++
	}
}

func (h *Header_Info) WriteBanner(banner string) {
	for _, b := range banner {
		termbox.SetCell(h.X, 1, b, h.Color, termbox.ColorDefault)
		h.X++
	}
}

func (h *Header_Info) WriteHeaderInfo() {
	h.WriteSitePage()
	h.WriteNowPage()
	h.WriteMaxPage()
}