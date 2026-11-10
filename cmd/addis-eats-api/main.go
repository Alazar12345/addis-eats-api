package main

import "fmt"

type Dish struct {
	Name  string
	Price int
}

func main() {
	dishes := []Dish{
		{Name: "Doro Wat", Price: 18500},
		{Name: "Shiro", Price: 12000},
		{Name: "Tibs", Price: 24000},
		{Name: "Injera Basket", Price: 4500},
	}
	for _, dish := range dishes {
		fmt.Printf("%-15s %.2f ETB\n", dish.Name, float64(dish.Price)/100)
	}

	subtotal := 0
	for _, dish := range dishes {
		subtotal += dish.Price
	}

	serviceFee := int(float64(subtotal) * 0.05)

	deliveryArea := "Piassa"
	deliveryFee := 0
	switch deliveryArea {
	case "Bole":
		deliveryFee = 3000

	case "Kazanchis":
		deliveryFee = 3500

	case "Piassa":
		deliveryFee = 4000

	case "Megengna":
		deliveryFee = 4500

	default:
		deliveryFee = 5000
	}
	if subtotal >= 50000 {
		deliveryFee = 0
	}

	grandTotal := subtotal + serviceFee + deliveryFee

	if grandTotal > 60000 {
		fmt.Println("Phone: +251911234567")
	}

	fmt.Printf("\nSubtotal: %.2f ETB\n", float64(subtotal)/100)
	fmt.Printf("Service Fee: %.2f ETB\n", float64(serviceFee)/100)
	fmt.Printf("Delivery Fee: %.2f ETB\n", float64(deliveryFee)/100)
	fmt.Printf("Grand Total: %.2f ETB\n", float64(grandTotal)/100)
}
