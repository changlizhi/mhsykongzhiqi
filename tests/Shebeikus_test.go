package tests

import (
	"github.com/lunny/log"
	"mhsykongzhiqi/kus"
	"mhsykongzhiqi/moxings"
	"testing"
)

func TestCharushebei(t *testing.T) {
	st := &moxings.Shebeis{
		Xuliehao: "11122223333",
		Macdizhi: "ad:ee:bb:aa:12:33",
		Pici:     1,
	}
	kus.Charushebei(st)
}

func TestChaxunshebei(t *testing.T) {
	st := moxings.Shebeis{
		Xuliehao: "14223586423140",
		Pici:     1,
	}
	ret := kus.Chaxunyigeshebei(st)
	log.Println(ret)
}
