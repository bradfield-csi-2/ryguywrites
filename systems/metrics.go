package metrics

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
)

// changed data representation to be column oriented
var cents [1000000]uint32
var ages [100000]uint8

func AverageAge() float64 { // can't change float64 to 32 because math stdlib assumes float64
	var i, sum1, sum2, sum3, sum4 uint32 = 0, 0, 0, 0, 0
	for i < 99997 {
		sum1 += uint32(ages[i])
		sum2 += uint32(ages[i+1])
		sum3 += uint32(ages[i+2])
		sum4 += uint32(ages[i+3])
		i += 4
	}
	// sum1 + sum2 + sum3 + sum4 won't overflow
	return float64(sum1 + sum2 + sum3 + sum4)/float64(i)
}

func AveragePaymentAmount() float64 {
	var i uint32 = 0
	var sum1, sum2, sum3, sum4 uint64 = 0, 0, 0, 0
	for i < 999997 {
		sum1 += uint64(cents[i])
		sum2 += uint64(cents[i+1])
		sum3 += uint64(cents[i+2])
		sum4 += uint64(cents[i+3])
		i += 4
	}
	return float64(sum1 + sum2 + sum3 + sum4)/100000000 // denominator is length of array * 100
}

// Compute the standard deviation of payment amounts
func StdDevPaymentAmount() float64 {
	var diff1, diff2 float64
	var mean, acc1, acc2 float64 = AveragePaymentAmount()*100, 0.0, 0.0 // keep units in cents to avoid division
	var i uint32 = 0
	for i < 999999 {
		diff1 = float64(cents[i]) - mean
		acc1 += diff1 * diff1
		diff2 = float64(cents[i+1]) - mean
		acc2 += diff2 * diff2
		i += 2
	}
	// denominator = size of array * 100 * 100, 100 * 100 accounts for units in cents
	return math.Sqrt((acc1 + acc2) / 10000000000)
}

// edited csv and LoadData to ignore all fields that aren't used
func LoadData() {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	reader := csv.NewReader(f)
	userLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse users.csv as csv", err)
	}

	for i, line := range userLines {
		age, _ := strconv.Atoi(line[0])
		ages[i] = uint8(age)
	}


	f, err = os.Open("payments.csv")
	if err != nil {
		log.Fatalln("Unable to read payments.csv", err)
	}
	reader = csv.NewReader(f)
	paymentLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse payments.csv as csv", err)
	}

	for i, line := range paymentLines {
		paymentCents, _ := strconv.Atoi(line[0])
		cents[i] = uint32(paymentCents)
	}
}
