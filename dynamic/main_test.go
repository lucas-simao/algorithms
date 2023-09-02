package main

import "testing"

func TestDynamicAlgorithms(t *testing.T) {
	var items = []Item{
		{Name: "Guitar", Value: 1500, Weight: 1},
		{Name: "Radio", Value: 3000, Weight: 4},
		{Name: "Notebook", Value: 2000, Weight: 3},
	}

	tests := map[string]struct {
		Items     []Item
		Expect    Item
		BagWeight int
	}{
		"should return {Name:Guitar, Notebook Value:3500 Weight:4}": {
			Items:     items,
			Expect:    Item{Name: "Guitar, Notebook", Value: 3500, Weight: 4},
			BagWeight: 4,
		},
		"should return {Name:Notebook, Iphone Value:4000 Weight:4}": {
			Items:     append(items, Item{Name: "Iphone", Value: 2000, Weight: 1}),
			Expect:    Item{Name: "Notebook, Iphone", Value: 4000, Weight: 4},
			BagWeight: 4,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			resp := DynamicAlgorithms(test.BagWeight, test.Items)

			if resp.Name != test.Expect.Name {
				t.Errorf("expect %s but got %s", test.Expect.Name, resp.Name)
			}

			if resp.Value != test.Expect.Value {
				t.Errorf("expect %d but got %d", test.Expect.Value, resp.Value)
			}

			if resp.Weight != test.Expect.Weight {
				t.Errorf("expect %d but got %d", test.Expect.Weight, resp.Weight)
			}
		})
	}
}
