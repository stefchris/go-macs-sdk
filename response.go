// Copyright (C) 2023 Stefan Christen <s.christen@dycom.ch>.
//
// Use of this source code is prohibited without
// written permission.

package sdk

import (
	"sort"
)

type Response struct {
	Version  int
	Code     string
	Text     string   `json:",omitempty"`
	Title    string   `json:",omitempty"`
	Searches []Search `json:",omitempty"`
	View     *View    `json:",omitempty"`
}

type Search struct {
	Code  string
	Title string
	Args  map[string]string
	Hits  []Hit
	More  int
}

type Hit struct {
	Icon string
	Text string
	Info string
	PKey string
}

type View struct {
	Title string
	Pages []Page
}

type Page struct {
	Title    string
	Info     string
	Icon     string
	More     string
	Level    int
	Sections []Section
}

type Section struct {
	Title string
	Rows  []Row
}

type Row struct {
	Label string
	Value string
	Type  string
	Name  string
}

func (r *Response) AddSearch(title string) *Search {
	r.Searches = append(r.Searches, Search{
		Code:  "OK",
		Title: title,
		Args:  map[string]string{},
	})
	return &r.Searches[len(r.Searches)-1]
}

func (s *Search) AddHit(icon, text, info, pkey string) {
	s.Hits = append(s.Hits, Hit{
		Icon: icon,
		Text: text,
		Info: info,
		PKey: pkey,
	})
}

func (s *Search) SortHits() {
	sort.Slice(s.Hits, func(i, j int) bool {
		if s.Hits[i].Text != s.Hits[j].Text {
			return s.Hits[i].Text < s.Hits[j].Text
		}
		return s.Hits[i].Info > s.Hits[j].Info
	})
}

func (r *Response) SetView(title string) *View {
	r.View = &View{
		Title: title,
	}
	return r.View
}

func (v *View) AddPage(title string) *Page {
	v.Pages = append(v.Pages, Page{
		Title: title,
	})
	return &v.Pages[len(v.Pages)-1]
}

func (v *View) AddSubPage(title, info, icon, more string, level int) *Page {
	v.Pages = append(v.Pages, Page{
		Title: title,
		Info:  info,
		Icon:  icon,
		More:  more,
		Level: level,
	})
	return &v.Pages[len(v.Pages)-1]
}

func (p *Page) AddSection(title string) *Section {
	p.Sections = append(p.Sections, Section{
		Title: title,
	})
	return &p.Sections[len(p.Sections)-1]
}

func (s *Section) AddRowLabelText(label, value string) {
	s.Rows = append(s.Rows, Row{
		Label: label,
		Value: value,
		Type:  "text/plain",
	})
}

func (s *Section) AddRowLabelPhoto(label, photo string) {
	s.Rows = append(s.Rows, Row{
		Label: label,
		Value: photo,
		Type:  "image/jpeg",
	})
}

func (s *Section) AddRowAlert(text string) {
	s.Rows = append(s.Rows, Row{
		Value: text,
		Type:  "text/alert",
	})
}

func (s *Section) AddRowWarning(text string) {
	s.Rows = append(s.Rows, Row{
		Value: text,
		Type:  "text/warning",
	})
}

func (s *Section) AddRowInputText(label, name, value string) {
	s.Rows = append(s.Rows, Row{
		Label: label,
		Value: value,
		Type:  "input/text",
		Name:  name,
	})
}

func (s *Section) AddRowInputTextNumber(label, name, value string) {
	s.Rows = append(s.Rows, Row{
		Label: label,
		Value: value,
		Type:  "input/number",
		Name:  name,
	})
}

func (s *Section) AddRowInputDate(label, name, value string) {
	s.Rows = append(s.Rows, Row{
		Label: label,
		Value: value,
		Type:  "input/date",
		Name:  name,
	})
}
