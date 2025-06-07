package main

type bill struct {
	name  string
	items map[string]float64
	tip   float64
	news  []map[string]any
}

func getBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{
			"ages":   22,
			"points": 22.1,
		},
		news: []map[string]any{
			{
				"title":   "Golang in Web Development",
				"author":  "Jane Smith",
				"content": "Using Go for backend services...",
				"views":   750,
			},
		},
		tip: 0,
	}

	return b
}

// func (bill bill) getName() {
	// return bill.name
// }
