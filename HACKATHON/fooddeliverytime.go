package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	time := food{0}

	if !(order == "burger") || !(order == "chips") || !(order == "nuggets") {
		time.preptime = 404
	}

	if order == "burger" {
		time.preptime = 15
	}

	if order == "chips" {
		time.preptime = 10
	}
	if order == "nuggets" {
		time.preptime = 12
	}
	return time.preptime
}
