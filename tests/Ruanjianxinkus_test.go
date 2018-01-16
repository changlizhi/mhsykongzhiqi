package tests

import (
	"testing"
	"mhsykongzhiqi/kus"
	"github.com/lunny/log"
)

func TestSuoyouruanjianxin(t *testing.T) {
	ret := kus.Suoyouruanjianxin("111")
	log.Println(ret)
}