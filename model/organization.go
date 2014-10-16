package model

import ()

type Organization struct {
	Id       int            `json:"id"`
	Name     string         `json:"name"`
	Currency string         `json:"currency"`
	Shipping string         `json:"shipping"`
	Address  Address        `json:"address"`
	Bank     Bank           `json:"bank"`
	Sites    []Organization `json:"sites"`
}

func (o Organization) String() string {
	return ""
}

func (o Organization) AddSite(site Organization) {
	o.Sites = append(o.Sites, site)
}
