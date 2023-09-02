package main

import "fmt"

// What items do you should steal to maximize the value stolen
// Exist 3 items and you have a maximum 4kg of space in the bag
// Radio $3000 = 4kg
// Notebook $2000 = 3kg
// Guitar $1500 = 1kg
// you'll need to steal max of items without passing off bag weight and with max cost

type Item struct {
	Name   string
	Value  int
	Weight int
}

func main() {
	var items = []Item{
		{Name: "Guitar", Value: 1500, Weight: 1},
		{Name: "Radio", Value: 3000, Weight: 4},
		{Name: "Notebook", Value: 2000, Weight: 3},
	}

	fmt.Printf("%+v\n", DynamicAlgorithms(4, items))

}

func DynamicAlgorithms(bagWeight int, items []Item) Item {
	// Matrix X,Y
	// X line = count of items
	// Y column = bagWeight
	bag := make([][]Item, len(items))

	for i := 0; i < len(items); i++ {
		bag[i] = make([]Item, bagWeight)
	}

	var bestChoose Item

	for i := 0; i < len(bag); i++ {
		for j := 0; j < len(bag[i]); j++ {
			if j+1 > items[i].Weight && i != 0 {
				if oldItem := bag[i-1][j-items[i].Weight]; oldItem.Weight+items[i].Weight <= j+1 {
					bag[i][j] = Item{
						Name:   fmt.Sprintf("%s, %s", oldItem.Name, items[i].Name),
						Value:  oldItem.Value + items[i].Value,
						Weight: oldItem.Weight + items[i].Weight,
					}
				}
			} else if items[i].Weight <= j+1 {
				bag[i][j] = items[i]
			} else if v := bag[i-1][j]; v.Value > 0 {
				bag[i][j] = v
			}

			if bag[i][j].Value > bestChoose.Value {
				bestChoose = bag[i][j]
			}
		}
	}

	return bestChoose
}
