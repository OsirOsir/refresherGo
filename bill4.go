package main

import "fmt"

type invoice struct {
	name string
	item map[string]float64
	tip  float64
}

// make new invoice
func newInvoice(name string) invoice {
	i := invoice{
		name: name,
		item: map[string]float64{},
		tip:  0,
	}
	return i
}

// format the invoice
func (i invoice) format() string {
	fs := "Invoice summary\n"
	var total float64 = 0

	for k, v := range i.item {
		fs = fs + fmt.Sprintf("%-25s   ...$%.1f\n", k, v)
		total += v
	}
	fs = fs + fmt.Sprintf("%-25s   ...$%.1f\n", "Tip:", i.tip)
	fs = fs + fmt.Sprintf("%-25s   ...$%.1f", "Total:", total+i.tip)
	return fs
}

//addditem
func (i invoice) updateItem(item string, price float64) {
	i.item[item] = price
}

//addtip

func (i *invoice) updateTip(tip float64) {
	i.tip = tip
}
