package todo

type Item struct {
	Text        string `json:"text"`
	Description string `json:"description"`
}

type Todos struct {
	Items []Item
}

func New() *Todos {
	return &Todos{}
}

// func (r *Todos) Add(item Item) {
// 	r.Items = append(r.Items, item)
// }

func (r *Todos) GetAll() []Item {
	return r.Items
}
