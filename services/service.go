package services

import (
	"math/rand"
	"time"
)

var length int = 3

//var DataSet map[string]string

var DataSet = make(map[string]string)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func Rand_String(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func UrlsDataSet(url string) string {
	rand_letters := Rand_String(length, charset)
	for {
		if val, ok := DataSet[url]; ok {
			return val
		} else {
			values := make([]string, 0, len(DataSet))
			for _, v := range DataSet {
				values = append(values, v)
			}
			val := contains(values, rand_letters)
			if !val {
				DataSet[url] = rand_letters
				return rand_letters
			}

		}

	}

}