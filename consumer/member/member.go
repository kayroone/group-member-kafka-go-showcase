package main

import "fmt"

type Member struct {
	Name  string `json:"name"`
	Group string `json:"group"`
}

func (m Member) String() string {

	return fmt.Sprintf("{Name: %s Group: %s}", m.Name, m.Group)
}
