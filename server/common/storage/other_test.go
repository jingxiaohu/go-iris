package storage

import (
	"strconv"
	"strings"
	"testing"
)

func Test_o1(t *testing.T)  {
	a := "1,4,5,"
	a1 := strings.Split(a, ",")

	var (
		err error
		id int64
	)
	for _, v := range a1 {
		if id, err = strconv.ParseInt(v, 10, 64); err != nil {
			t.Log(err)
		}
		t.Logf("id=%d", id)
	}
	//t.Logf("len=%d, val=%s", len(a1), a1)
}
