package tests

import (
	"log"
	"mhsykongzhiqi/kus"
	"testing"
)

func TestSuoyouruanjianxin(t *testing.T) {
	ret := *kus.Suoyouruanjianxin("111")
	for _, r := range ret {
		log.Println(r.Id)
	}
}
