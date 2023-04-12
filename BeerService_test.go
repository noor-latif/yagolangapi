package main

//OBS: Svårt? Helt OK
// Vi är på en programmeringsteknisk NÖRD-nivå
//det finns inga krav på att ni ska lära er skriva tester
//jag kommer ge er tester ifall det är så
// det viktiga är inse att det är en del i PROCESSEN

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenPromilleIsGreaterThanOnePointFiveThenNotAllowed(t *testing.T) {
	//arrange
	age := 50
	location := "Systemet"
	promille := float32(1.55)
	//act
	result := CanIBuyBeer(age, location, promille)
	//assert
	//assert.Equal(t,123, i)
	assert.False(t, result)

}

func TestWhenAgeIs19AndLocationIsSystemetThenNotAllowed(t *testing.T) {
	//arrange
	age := 19
	location := "Systemet"
	promille := float32(0)
	//act
	result := CanIBuyBeer(age, location, promille)
	//assert
	//assert.Equal(t,123, i)
	assert.False(t, result)

}

func TestWhenAgeIs50AndLocationIsSystemetThenAllowed(t *testing.T) {
	//arrange
	age := 50
	location := "Systemet"
	promille := float32(0)
	//act
	result := CanIBuyBeer(age, location, promille)
	//assert
	//assert.Equal(t,123, i)
	assert.True(t, result)

}
