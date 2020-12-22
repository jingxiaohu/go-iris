package storage
//
//import (
//	"context"
//	"github.com/kataras/golog"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"go.mongodb.org/mongo-driver/mongo/readpref"
//	"goiris/common"
//	"time"
//)
//
//var G_Mongodb *Mongodb
//
//func InitMongo() {
//	var (
//		client *mongo.Client
//		err    error
//	)
//	opts := options.Client().
//		ApplyURI(common.G_DBConfig.Mongodb.Addr).
//		SetConnectTimeout(time.Duration(common.G_DBConfig.Mongodb.ConnectTimeout) * time.Second).
//		SetMaxPoolSize(uint64(common.G_DBConfig.Mongodb.MaxPoolSize))
//	if client, err = mongo.Connect(context.TODO(), opts); err != nil {
//		golog.Fatalf("~~> Mongodb初始化错误,原因:%s", err.Error())
//	}
//
//	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
//		golog.Fatalf("~~> Mongodb服务不可用,原因:%s", err.Error())
//	}
//
//	database := client.Database(common.G_DBConfig.Mongodb.Database)
//	G_Mongodb = &Mongodb{
//		AdminCollection: database.Collection("admin"),
//		AreaCollection:  database.Collection("area"),
//		ClassCollection: database.Collection("class"),
//		ItemCollection:  database.Collection("item"),
//		TradeCollection: database.Collection("trade"),
//	}
//}
//
//type Mongodb struct {
//	// 对应mongo所有表的映射
//	AdminCollection *mongo.Collection
//	AreaCollection  *mongo.Collection
//	ClassCollection *mongo.Collection
//	ItemCollection  *mongo.Collection
//	TradeCollection *mongo.Collection
//}
