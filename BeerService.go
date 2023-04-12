package main

func CanIBuyBeer(age int, location string, promille float32) bool {
	if promille > 1.5 {
		return false
	}
	if age >= 18 && location == "Krogen" {
		return true
	}
	if age >= 20 && location == "Systemet" {
		return true
	}
	return false
}
