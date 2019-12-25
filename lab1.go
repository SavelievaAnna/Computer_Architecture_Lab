package main

import
    ( 
//        "fmt"
        //"math/big"
      //  "sync"
    )
func check(MorzeString string) bool {
	i := 0
	for i < len(MorzeString) {
		if MorzeString[i] != '.' && MorzeString[i] != '-' && MorzeString[i] != ' ' {
			return false
		}
		i++
	}
	return true
}

func searchdict(s string, dict map[string]string) string {
	for key, value := range dict {
		if value == s {
			return key
		}
	}
	return "false"
}

func checkarr(MorzeArr []string, dict map[string]string) bool {
	i := 0
	for i < len(MorzeArr) {
		if searchdict(MorzeArr[i], dict) == "false" {
			return false
		}
		i++
	}
	return true
}

func makearr(MorzeString string) []string {
	helpStr := ""
	var morzeArr []string
	i := 0
	spaceCounter := 0
	for i < len(MorzeString) {
		if MorzeString[i] == ' ' {
			spaceCounter++
			if spaceCounter == 1 {
				morzeArr = append(morzeArr, helpStr)
				helpStr = ""
				i++
				continue
			} else if spaceCounter == 2 {
				morzeArr = append(morzeArr, "  ")
				helpStr = ""
				i++
				continue
			}
		}
		spaceCounter = 0
		helpStr += string(MorzeString[i])
		i++
	}
	if spaceCounter == 0 {
		morzeArr = append(morzeArr, helpStr)
	}
	return morzeArr
}

func decode(MorzeArray []string, dict map[string]string) string {
	i := 0
	decoded := ""
	if checkarr(MorzeArray, dict) == true {
		for i < len(MorzeArray) {
			decoded += searchdict(MorzeArray[i], dict)
			i++
		}
		return decoded
	}
	return " - Wrong morze code"
}

func encode(s string, dict map[string]string) string {
	i := 0
	encoded := ""
	for i < len(s) {
		if s[i] == ' ' {
			encoded += " "
			i++
			continue
		}
		encoded += dict[string(s[i])]
		if s[i] != ' ' {
			encoded += " "
		}
		i++
	}
	return encoded
}

func encode_fast(s string, dict map[string]string) string {
    type empty struct{}
    N := 8
    len_one_part := len(s) / N
    
    data := make([]string, N);
    data[0] = s[0*len_one_part:1*len_one_part]
    data[1] = s[1*len_one_part:2*len_one_part]
    data[2] = s[2*len_one_part:3*len_one_part]
    data[3] = s[3*len_one_part:4*len_one_part]
    data[4] = s[4*len_one_part:5*len_one_part]
    data[5] = s[5*len_one_part:6*len_one_part]
    data[6] = s[6*len_one_part:7*len_one_part]
    data[7] = s[7*len_one_part:len(s)]
    
    res := make([]string, N);
    
    sem := make(chan empty, N);  // semaphore pattern
        
    for i,xi := range data {
        go func (i int, xi string) {
            res[i] = encode(data[i], dict);
            sem <- empty{};
        } (i, xi);
    }
    // wait for goroutines to finish
    for i := 0; i < N; i++ {
        <-sem 
    }
    
    encoded := ""
    for i := 0; i < N; i++ {
        encoded += res[i] 
    }
    
    return encoded
}


func main() {
    
}




