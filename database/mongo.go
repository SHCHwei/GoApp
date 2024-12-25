package database

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx         context.Context
	err         error
)

type MongoDB struct{
    Collection *mongo.Collection
    Client *mongo.Client
    Ctx context.Context
}

func(m *MongoDB) connect() {

    conn := options.Client().ApplyURI("mongodb://root:minmin@localhost:27017")
    m.Client, err = mongo.Connect(ctx, conn)

    if err != nil {
        log.Printf("DB Failed : %v \n", err)
        panic("mongo database connection Failed")
    }else{
        log.Println("mongo database connection Successful")
    }

    m.Collection = m.Client.Database("Lesson").Collection("session")
    m.Ctx = ctx
}


func(m *MongoDB) disConnect(){

    m.Client.Disconnect(m.Ctx)
    log.Println("disConnect func.")
}


// func GetObjId(obj ObjectID)(string){
//
//
//
//
//     return idString
// }