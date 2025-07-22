package blockchain

import (
	"fmt"
	"math/rand"
	"time"
)

type Validator struct {
	Name  string
	Stake int
}

func SelectValidator(validators []Validator) Validator {
	totalStake := 0
	for _, v := range validators {
		totalStake += v.Stake
	}
	rand.New(rand.NewSource(time.Now().Unix()))
	r := rand.Intn(totalStake)
	sum := 0
	for _, v := range validators {
		sum += v.Stake
		if r < sum {
			return v
		}
	}
	return validators[0]
}
func StakingCoins() {
	validators := []Validator{
		{"Validator A", 1000},
		{"Validator B", 3000},
		{"Validator C", 6000},
	}
	for i := 1; i <= 5; i++ {
		selected := SelectValidator(validators)
		fmt.Printf("Round %d: %s (Stake:%d) \n", i, selected.Name, selected.Stake)
	}
}
