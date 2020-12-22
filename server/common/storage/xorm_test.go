package storage

import (
	"goiris/admin/app/bindata/conf"
	"goiris/common"
	"testing"
	"time"
)

func Test_1(t *testing.T) {
	common.InitConfig(conf.Asset)
	InitXorm()

	var result []int64
	if err := G_Xorm.Table("sub_class").Cols("id").Find(&result); err != nil {
		t.Log(err)
		return
	}
	t.Logf("%v", result)
}

func Test_2(t *testing.T) {
	t.Log(time.Now().Format("2006-01-02 15:04:05"))
}