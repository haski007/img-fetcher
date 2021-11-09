package model

type XKom struct {
	ProdTitle string   `json:"prod_title"`
	Images    []string `json:"images"`
}

func (rcv XKom) GetTitle() string {
	return rcv.ProdTitle
}

func (rcv XKom) GetImages() []string {
	return rcv.Images
}
