package utils

import (
	"log"
	"strconv"
)

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("cannot convert %s to int", str)
	}

	return num
}
