package main

import (
//	"fmt"
	"testing"
        "io/ioutil"
//        "runtime"
)

/*func TestCheck(t *testing.T) {
    i := 0
    for i < 100000 {
        var str = tests.RandomWord(i)
        var st = true
        if st != check(str) {
            t.Errorf("%s != %t", str, st)
        }
        i++
    }
}

func TestSearchDict(t *testing.T) {
    MorzeDict := map[string]string{
		"A": ".-",
		"B": "-...",
		"C": "-.-.",
		"D": "-..",
		"E": ".",
		"F": "..-.",
		"G": "--.",
		"H": "....",
		"I": "..",
		"J": ".---",
		"K": "-.-",
		"L": ".-..",
		"M": "--",
		"N": "-.",
		"O": "---",
		"P": ".--.",
		"Q": "--.-",
		"R": ".-.",
		"S": "...",
		"T": "-",
		"U": "..-",
		"V": "...-",
		"W": ".--",
		"X": "-..-",
		"Y": "-.--",
		"Z": "--..",
		"1": ".----",
		"2": "..---",
		"3": "...--",
		"4": "....-",
		"5": ".....",
		"6": "-....",
		"7": "--...",
		"8": "---..",
		"9": "----.",
		"0": "-----",
		" ": "  ",
	}
    for key, str := range(MorzeDict) {
        if key != searchdict(str, MorzeDict) {
            t.Errorf("%s != %s", str, key)
        }
    }
}

func arrComp(arr1, arr2 []string) bool {
    if len(arr1) != len(arr2) {
        return false
    }
    i := 0
    for i < len(arr1) {
        if arr1[i] != arr2[i] {
            return false
        }
        i++
    }
    return true
}

func TestMakeArr(t *testing.T) {
    i := 0
    for i < 100000 {
        var str = tests.RandomWord(i)
        arr := makearr(str)
        if !arrComp(arr, makearr(str)) {
            t.Errorf("arrs are not equal")
        }
        i++
    }
}

func TestCheckArr(t *testing.T) {
    MorzeDict := map[string]string{
		"A": ".-",
		"B": "-...",
		"C": "-.-.",
		"D": "-..",
		"E": ".",
		"F": "..-.",
		"G": "--.",
		"H": "....",
		"I": "..",
		"J": ".---",
		"K": "-.-",
		"L": ".-..",
		"M": "--",
		"N": "-.",
		"O": "---",
		"P": ".--.",
		"Q": "--.-",
		"R": ".-.",
		"S": "...",
		"T": "-",
		"U": "..-",
		"V": "...-",
		"W": ".--",
		"X": "-..-",
		"Y": "-.--",
		"Z": "--..",
		"1": ".----",
		"2": "..---",
		"3": "...--",
		"4": "....-",
		"5": ".....",
		"6": "-....",
		"7": "--...",
		"8": "---..",
		"9": "----.",
		"0": "-----",
		" ": "  ",
	}
    i := 0
    for i < 100000 {
        var str = tests.RandomWord(i)
        var arr = makearr(str)
        var b = checkarr(arr, MorzeDict)
        if b != checkarr(arr, MorzeDict) {
            t.Errorf("Failed to make an arr")
        }
        i++
    }
}

func TestDecode(t *testing.T) {
    MorzeDict := map[string]string{
		"A": ".-",
		"B": "-...",
		"C": "-.-.",
		"D": "-..",
		"E": ".",
		"F": "..-.",
		"G": "--.",
		"H": "....",
		"I": "..",
		"J": ".---",
		"K": "-.-",
		"L": ".-..",
		"M": "--",
		"N": "-.",
		"O": "---",
		"P": ".--.",
		"Q": "--.-",
		"R": ".-.",
		"S": "...",
		"T": "-",
		"U": "..-",
		"V": "...-",
		"W": ".--",
		"X": "-..-",
		"Y": "-.--",
		"Z": "--..",
		"1": ".----",
		"2": "..---",
		"3": "...--",
		"4": "....-",
		"5": ".....",
		"6": "-....",
		"7": "--...",
		"8": "---..",
		"9": "----.",
		"0": "-----",
		" ": "  ",
	}
    i := 0
    for i < 100000 {
        var str = tests.RandomWord(i)
        var arr = makearr(str)
        var st = " - Wrong morze code"
        if checkarr(arr, MorzeDict) == false {
            if decode(arr, MorzeDict) != st {
                t.Errorf("Decode failed")
            }
        }
        i++
    }
}
*/
func TestDecode(t *testing.T) {
	MorzeDict := map[string]string{
		"A": ".-",
		"B": "-...",
		"C": "-.-.",
		"D": "-..",
		"E": ".",
		"F": "..-.",
		"G": "--.",
		"H": "....",
		"I": "..",
		"J": ".---",
		"K": "-.-",
		"L": ".-..",
		"M": "--",
		"N": "-.",
		"O": "---",
		"P": ".--.",
		"Q": "--.-",
		"R": ".-.",
		"S": "...",
		"T": "-",
		"U": "..-",
		"V": "...-",
		"W": ".--",
		"X": "-..-",
		"Y": "-.--",
		"Z": "--..",
		"1": ".----",
		"2": "..---",
		"3": "...--",
		"4": "....-",
		"5": ".....",
		"6": "-....",
		"7": "--...",
		"8": "---..",
		"9": "----.",
		"0": "-----",
		" ": "  ",
	}
        contents,_ := ioutil.ReadFile("tests.txt")
	var str = string(contents)
        
        ioutil.WriteFile("output.txt", []byte(encode(str, MorzeDict)), 0644)
}/*

func TestDecodeFast(t *testing.T) {
        runtime.GOMAXPROCS(8)
	MorzeDict := map[string]string{
		"A": ".-",
		"B": "-...",
		"C": "-.-.",
		"D": "-..",
		"E": ".",
		"F": "..-.",
		"G": "--.",
		"H": "....",
		"I": "..",
		"J": ".---",
		"K": "-.-",
		"L": ".-..",
		"M": "--",
		"N": "-.",
		"O": "---",
		"P": ".--.",
		"Q": "--.-",
		"R": ".-.",
		"S": "...",
		"T": "-",
		"U": "..-",
		"V": "...-",
		"W": ".--",
		"X": "-..-",
		"Y": "-.--",
		"Z": "--..",
		"1": ".----",
		"2": "..---",
		"3": "...--",
		"4": "....-",
		"5": ".....",
		"6": "-....",
		"7": "--...",
		"8": "---..",
		"9": "----.",
		"0": "-----",
		" ": "  ",
	}
        contents,_ := ioutil.ReadFile("tests.txt")
	var str = string(contents)
        ioutil.WriteFile("output.txt", []byte(encode_fast(str, MorzeDict)), 0644)
}*/