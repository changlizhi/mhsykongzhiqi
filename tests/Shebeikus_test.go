package tests

import (
	"mhsykongzhiqi/kus"
	"testing"
	"mhsykongzhiqi/moxings"
)

func TestCharushebei(t *testing.T) {
	st := &moxings.Shebeis{
		Xuliehao:"11122223333",
		Macdizhi:"ad:ee:bb:aa:12:33",
		Pici:1,
	}
	kus.Charushebei(st)
}
