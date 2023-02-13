package cache

import (
	"Minimalist_TikTok/model/websocket"
	"Minimalist_TikTok/server/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"time"
)

type SendSortMsg struct {
	Content  string `json:"content"`
	Read     string `json:"read"`
	CreateAt int64  `json:"create_at"`
}

func InsertMsg(database string, id string, content string, read uint, expire int64) (err error) {
	collection := config.MongoDBClient.Database(database).Collection(id)
	comment := websocket.Trainer{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expire,
		Read:      read,
	}
	_, err = collection.InsertOne(context.TODO(), comment)
	return
}

func FindMany(database, sendID, id string, time int64, pageSize int) (results []websocket.Result, err error) {
	//fil:=bson.D{{"read",
	//	bson.D{{
	//		"$in",
	//		bson.A{0, 1},
	//	}},}}
	var resultMe []websocket.Trainer  //ID
	var resultYou []websocket.Trainer //sendID
	sendIDCollection := config.MongoDBClient.Database(database).Collection(sendID)
	idCollection := config.MongoDBClient.Database(database).Collection(id)
	fmt.Println("idCollection", idCollection)
	sendIDTimeCurcor, err := sendIDCollection.Find(context.TODO(), bson.D{{}},
		options.Find().SetSort(bson.D{{"startTime", -1}}), options.Find().SetLimit(int64(pageSize)))

	/*.options.Find()SetSort(bson.D{{"startTime",-1}}),
	options.Find().SetLimit(int64(pageSize))*/
	idTimeCurcor, err := idCollection.Find(context.TODO(), bson.D{{}},
		options.Find().SetSort(bson.D{{"startTime", -1}}), options.Find().SetLimit(int64(pageSize)))

	fmt.Println("idTimeCurcor", idTimeCurcor)
	err = sendIDTimeCurcor.All(context.TODO(), &resultYou)
	fmt.Println("resultYou", resultYou)

	err = idTimeCurcor.All(context.TODO(), &resultMe)
	fmt.Println("resultMe", resultMe)

	results, _ = AppendAndSort(resultMe, resultYou)
	return
}

func AppendAndSort(resultMe, resultYou []websocket.Trainer) (results []websocket.Result, err error) {
	for _, r := range resultMe {
		sendSort := SendSortMsg{
			Content:  r.Content,
			Read:     string(r.Read),
			CreateAt: r.StartTime,
		}
		result := websocket.Result{
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "me",
		}
		results = append(results, result)
	}
	for _, r := range resultYou {
		sendSort := SendSortMsg{
			Content:  r.Content,
			Read:     string(r.Read),
			CreateAt: r.StartTime,
		}
		result := websocket.Result{
			StartTime: r.StartTime,
			Msg:       fmt.Sprintf("%v", sendSort),
			From:      "you",
		}
		results = append(results, result)
	}
	return
}
