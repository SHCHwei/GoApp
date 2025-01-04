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
    Database *mongo.Database
    Client *mongo.Client
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

    m.Database = m.Client.Database("Lesson")
}


func(m *MongoDB) disConnect(){

    m.Client.Disconnect(ctx)
    log.Println("disConnect func.")
}
