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
		termbox.SetCell(h.X, 1, header_info_site_page, h.Color, h.BackgroundColor)
		h.X++
	}
}

func (h *Header_Info) WriteNowPage() {
	termbox.SetCell(h.X, 1, ' ', h.Color, h.BackgroundColor)
	h.X++

	h.WriteBanner("| Now Page:")
	for _, header_info_now_page := range h.Now_Page {
		termbox.SetCell(h.X, 1, header_info_now_page, h.Color, h.BackgroundColor)
		h.X++
	}
}

func (h *Header_Info) WriteMaxPage() {
	termbox.SetCell(h.X, 1, ' ', h.Color, h.BackgroundColor)
	h.X++

	h.WriteBanner("| Max Page:")
	for _, header_info_max_page := range h.Max_Page {
		termbox.SetCell(h.X, 1, header_info_max_page, h.Color, h.BackgroundColor)
		h.X++
	}
}

func (h *Header_Info) WriteBanner(banner string) {
	for _, b := range banner {
		termbox.SetCell(h.X, 1, b, h.Color, h.BackgroundColor)
		h.X++
	}
}

func (h *Header_Info) WriteHeaderInfo() {
	termWidth, _ := termbox.Size()

	for col := 0; col < termWidth; col++ {
		termbox.SetCell(col, 1, ' ', h.Color, h.BackgroundColor)
	}

	h.WriteSitePage()
	h.WriteNowPage()
	h.WriteMaxPage()
}

func (h *Header_Info) GetNowPage() int {
	now_page, _ := strconv.Atoi(h.Now_Page)

	return now_page
}
