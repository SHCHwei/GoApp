package session

import(
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
    "time"
    "context"
    "fmt"

)


type SessionUser struct {
    ID          primitive.ObjectID `bson:"_id"`
	UserID      string
	Mfg         time.Time
	Exp         time.Time
}


func InitSession(id string)(*SessionUser){

    start := time.Now()

    return &SessionUser{
        ID: primitive.NewObjectID(),
        UserID: id,
        Mfg: start,
        Exp: start.Add(time.Hour * 1),
    }
}


func DeleteSession(mc *mongo.Collection){

    nowTime := time.Now
    filter := bson.D{{"exp", bson.D{{"$gte", nowTime}} }}

    res, err := mc.DeleteMany(context.TODO(), filter)

    if err != nil {
        fmt.Println("session err",  err)
    }else{
        fmt.Println("res", res.DeletedCount)
    }

}
