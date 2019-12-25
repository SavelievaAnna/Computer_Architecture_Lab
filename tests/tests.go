package tests

import (
	"math/rand"
)

func RandomOne(n int) string {
	rand.Seed(int64(n))
    ch := rand.Intn(2)
    if ch == 0 {
        return "."
    }
	return "-"
}

func RandomTwo(n int) string {
	var l1, l2 string
	rand.Seed(int64(n+1))
	ch1 := rand.Intn(2)
	rand.Seed(int64(n+2))
	ch2 := rand.Intn(2)
	if ch1 == 0 {
		l1 = "."
	} else {
		l1 = "-"
	}
	if ch2 == 0 {
		l2 = "."
	} else {
		l2 = "-"
	}
	return l1 + l2
}

func RandomThree(n int) string {
	var l1, l2, l3 string
	rand.Seed(int64(n+1))
	ch1 := rand.Intn(2)
	rand.Seed(int64(n+2))
	ch2 := rand.Intn(2)
	rand.Seed(int64(n+3))
	ch3 := rand.Intn(2)
	if ch1 == 0 {
		l1 = "."
	} else {
		l1 = "-"
	}
	if ch2 == 0 {
		l2 = "."
	} else {
		l2 = "-"
	}
	if ch3 == 0 {
		l3 = "."
	} else {
		l3 = "-"
	}
	return l1 + l2 + l3
}

func RandomFive(n int) string {
	var l1, l2, l3, l4, l5 string
	rand.Seed(int64(n+1))
	ch1 := rand.Intn(2)
	rand.Seed(int64(n+2))
	ch2 := rand.Intn(2)
	rand.Seed(int64(n+3))
	ch3 := rand.Intn(2)
	rand.Seed(int64(n+4))
	ch4 := rand.Intn(2)
	rand.Seed(int64(n+5))
	ch5 := rand.Intn(2)
	if ch1 == 0 {
		l1 = "."
	} else {
		l1 = "-"
	}
	if ch2 == 0 {
		l2 = "."
	} else {
		l2 = "-"
	}
	if ch3 == 0 {
		l3 = "."
	} else {
		l3 = "-"
	}
	if ch4 == 0 {
		l4 = "."
	} else {
		l4 = "-"
	}
	if ch5 == 0 {
		l5 = "."
	} else {
		l5 = "-"
	}
	return l1 + l2 + l3 + l4 + l5
}

func RandomFour(n int) string {
	var l1, l2, l3, l4 string
	rand.Seed(int64(n+1))
	ch1 := rand.Intn(2)
	rand.Seed(int64(n+2))
	ch2 := rand.Intn(2)
	rand.Seed(int64(n+3))
	ch3 := rand.Intn(2)
	rand.Seed(int64(n+4))
	ch4 := rand.Intn(2)
	if ch1 == 0 {
		l1 = "."
	} else {
		l1 = "-"
	}
	if ch2 == 0 {
		l2 = "."
	} else {
		l2 = "-"
	}
	if ch3 == 0 {
		l3 = "."
	} else {
		l3 = "-"
	}
	if ch4 == 0 {
		l4 = "."
	} else {
		l4 = "-"
	}
	return l1 + l2 + l3 + l4
}

func RandomWord(n int) string {
	word := ""
	i := 0
	for i < 15 {
		rand.Seed(int64(n + i))
		x := rand.Intn(4)
		if x == 0 {
			rand.Seed(int64(n + i + 7))
			word += RandomOne(rand.Intn(33521))
			if i != 14 {
				word += " "
			}
		} else if x == 1{
			rand.Seed(int64(n + i + 22))
			word += RandomTwo(rand.Intn(14999))
			if i != 14 {
				word += " "
			}
		} else if x == 2{
			rand.Seed(int64(n + i + 11))
			word += RandomThree(rand.Intn(2888))
			if i != 14 {
				word += " "
			}
		} else if x == 3{
			rand.Seed(int64(n + i + 1888))
			word += RandomFour(rand.Intn(54444))
			if i != 14 {
				word += " "
			}
		} else if x == 4{
			rand.Seed(int64(n + i + 123))
			word += RandomFive(rand.Intn(62626))
			if i != 14 {
				word += " "
			}
		}
		i++
	}
	return word
}