package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var number string = ""
	var power_ball_num string = ""
	var power []string

	for iterate := 0; iterate <= 4; iterate++ {
		number = generate_lotto_nums(52)
		power = append(power, number)
	}

	power_ball_num = generate_lotto_nums(20)
	if power_ball_num != "0" {
		power_ball_num = "Powerball number: " + power_ball_num
		power = append(power, power_ball_num)
		fmt.Println(power)
	}

	filename := "numbers.txt"
	chosen := strings.Join(power, ",")

	err := os.WriteFile(filename, []byte(chosen), 0644)
	check(err)

	f, err := os.Create("numbers.txt")
	check(err)

	f.WriteString(chosen)

	f.Sync()
}

func generate_lotto_nums(max_num int) string {
	var numbers = single_num_generator(max_num)
	if numbers == 0 {
		numbers = single_num_generator(max_num)
	}
	return strconv.Itoa(numbers)
}

func single_num_generator(max_num int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max_num)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
