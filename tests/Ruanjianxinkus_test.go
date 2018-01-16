package tests

import (
	"testing"
	"mhsykongzhiqi/kus"
	"log"
)

func TestSuoyouruanjianxin(t *testing.T) {
	ret := *kus.Suoyouruanjianxin("111")
	for _,r := range ret{
		log.Println(r.Id)
	}
}