package model

import ()

type Proposol struct {
	Number       string       `json:"number"`
	Mark         Address      `json:"mark"`
	Bank         Bank         `json:"bank"`
	Organization Organization `json:"org"`
}
