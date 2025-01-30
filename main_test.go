package main

import (
	"fmt"
	"testing"
)

func Test_calculateTotal(t *testing.T) {
	tests := []struct {
		name            string
		menu            []foodSet
		discountedItems map[string]bool
		quantities      map[string]int
		wantNormal      float64
		wantDiscount    float64
		wantDiscountAmt float64
		wantAfterDisc   float64
		wantSubtotal    float64
	}{
		{
			name: "Ordering Red set and Green set (without membership)",
			menu: []foodSet{
				{"Red set", 50},
				{"Green set", 40},
			},
			discountedItems: map[string]bool{
				"Green set": true,
			},
			quantities:      map[string]int{"Red set": 1, "Green set": 1},
			wantNormal:      90, 
			wantDiscount:    0,  
			wantDiscountAmt: 0,  
			wantAfterDisc:   0,  
			wantSubtotal:    90, 
		},
		{
			name: "Ordering 2 Orange sets (5% discount applies)",
			menu: []foodSet{
				{"Orange set", 120},
			},
			discountedItems: map[string]bool{
				"Orange set": true,
			},
			quantities:      map[string]int{"Orange set": 2},
			wantNormal:      0,
			wantDiscount:    240, 
			wantDiscountAmt: 12,  
			wantAfterDisc:   228, 
			wantSubtotal:    228,
		},
		{
			name: "Ordering 0 item",
			menu: []foodSet{
				{"Red set", 50},
				{"Green set", 40},
			},
			discountedItems: map[string]bool{
				"Green set": true,
			},
			quantities:      map[string]int{},
			wantNormal:      0,
			wantDiscount:    0,
			wantDiscountAmt: 0,
			wantAfterDisc:   0,
			wantSubtotal:    0,
		},
		{
			name: "Ordering 2 Pink sets and one Orange set",
			menu: []foodSet{
				{"Pink set", 80},
				{"Orange set", 120},
			},
			discountedItems: map[string]bool{
				"Pink set": true,
				"Orange set": true,
			},
			quantities: map[string]int{
				"Pink set": 2,
				"Orange set": 1,
			  },
			wantNormal:      120,
			wantDiscount:    160,
			wantDiscountAmt: 8,
			wantAfterDisc:   152,
			wantSubtotal:    272,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNormal, gotDiscount, gotDiscountAmt, gotAfterDisc, gotSubtotal := calculateTotalTestHelper(tt.menu, tt.discountedItems, tt.quantities)

			if gotNormal != tt.wantNormal {
				t.Errorf("Normal total = %v, want %v", gotNormal, tt.wantNormal)
			}
			if gotDiscount != tt.wantDiscount {
				t.Errorf("Discounted total = %v, want %v", gotDiscount, tt.wantDiscount)
			}
			if gotDiscountAmt != tt.wantDiscountAmt {
				t.Errorf("Discount amount = %v, want %v", gotDiscountAmt, tt.wantDiscountAmt)
			}
			if gotAfterDisc != tt.wantAfterDisc {
				t.Errorf("Total after discount = %v, want %v", gotAfterDisc, tt.wantAfterDisc)
			}
			if gotSubtotal != tt.wantSubtotal {
				t.Errorf("Subtotal = %v, want %v", gotSubtotal, tt.wantSubtotal)
			}
		})
	}
}

func calculateTotalTestHelper(menu []foodSet, discountedItems map[string]bool, quantities map[string]int) (float64, float64, float64, float64, float64) {
	var normalTotal, discountTotal float64

	for _, food := range menu {
		quantity := quantities[food.Name]
		if quantity > 0 {
			if discountedItems[food.Name] && quantity >= 2 {
				discountTotal += float64(quantity * food.Price)
			} else {
				normalTotal += float64(quantity * food.Price)
			}
		}
	}

	afterDiscountTotal := discountTotal * 0.95
	discount := discountTotal * 0.05
	subtotal := afterDiscountTotal + normalTotal

	return normalTotal, discountTotal, discount, afterDiscountTotal, subtotal
}

func Test_calculateMemberDiscount(t *testing.T) {
	tests := []struct {
		name		string
		subtotal	float64
		haveMember	bool
		want		float64
	}{
		{"No membership", 100, false, 100}, //without membership
		{"With membership", 100, true, 90},
		{"Red and Green with membership", 90, true, 81},
		{"Zero subtotal", 0, true, 0},
		{"Large subtotal", 1000, true, 900},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateMemberDiscountTest(tt.subtotal, tt.haveMember)
			if got != tt.want {
				t.Errorf("calculateMemberDiscountTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func calculateMemberDiscountTest(subtotal float64, haveMember bool) float64 {
	var total float64
	var memberDiscount float64
	if haveMember {
		total = subtotal*0.9
		memberDiscount = subtotal*0.1
		fmt.Printf("10 percent member discount: %.2f THB", memberDiscount)
	} else {
		total = subtotal
	}
	return total
}
