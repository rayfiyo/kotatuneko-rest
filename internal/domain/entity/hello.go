package entity

import (
	_ "github.com/rayfiyo/kotatuneko-rest/docs"
)

type Response struct {
	Int64  int64  `json:"int64"`
	String string `json:"string"`
	World  *Item  `json:"world"`
}

type Item struct {
	Text string `json:"text"`
}
