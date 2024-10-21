package main

func f4(n int) {
	var s any
	for i := 0; i < n; i += 15 {
		s = i + 1
		_ = s
		s = i + 2
		_ = s
		s = "Fizz"
		_ = s
		s = i + 4
		_ = s
		s = "Buzz"
		_ = s
		s = "Fizz"
		_ = s
		s = i + 7
		_ = s
		s = i + 8
		_ = s
		s = "Fizz"
		_ = s
		s = "Buzz"
		_ = s
		s = i + 11
		_ = s
		s = "Fizz"
		_ = s
		s = i + 13
		_ = s
		s = i + 14
		_ = s
		s = "FizzBuzz"
		_ = s
	}
}

func f2(n int) {
	var s any
	for i := 0; i < n; i++ {
		if i%3 == 0 && i%5 == 0 {
			s = "FizzBuzz"
		} else if i%3 == 0 {
			s = "Fizz"
		} else if i%5 == 0 {
			s = "Buzz"
		} else {
			s = i
		}
		_ = s
	}
}

func f1(n int) {
	var s any
	fizzCounter, buzzCounter := 2, 4
	for i := 1; i <= n; i++ {
		isFizz := fizzCounter == 0
		isBuzz := buzzCounter == 0

		if isFizz && isBuzz {
			s = "FizzBuzz"
			fizzCounter, buzzCounter = 2, 4
		} else if isFizz {
			s = "Fizz"
			fizzCounter = 2
		} else if isBuzz {
			s = "Buzz"
			buzzCounter = 4
		} else {
			s = i
		}

		fizzCounter--
		buzzCounter--
		_ = s
	}
}

const (
	mod15 = "\x01\x02\x01\x03\x01\x01\x02\x01\x03\x01\x01\x02\x01\x01\x03\x01"
)

func f3(n int) {

	var (
		i uint8
	)
	var s any
	for j := 0; j < n; j++ {
		switch mod15[i] {
		case 1:
			s = j + 1
		case 2:
			s = "Fizz"
		case 3:
			s = "Buzz"
		default:
			s = "FizzBuzz"
		}
		i = (i + 1) & 15
		_ = s
	}
}
