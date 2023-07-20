package package_rtm

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
func InsertDataList(db *mongo.Database, Nama_ListRapat, Tanggal_ListRapat string, Lokasi_ListRapat string, Agenda_ListRapat string) (InsertedID interface{}) {
	var listrapat list_rapat
	listrapat.Nama_ListRapat = Nama_ListRapat
	listrapat.Tanggal_ListRapat = Tanggal_ListRapat
	listrapat.Lokasi_ListRapat = Lokasi_ListRapat
	listrapat.Agenda_ListRapat = Agenda_ListRapat
	return InsertOneDoc(db, "list_rapat", listrapat)
}

func GetDataListFromAgenda(Agenda_ListRapat string, db *mongo.Database, col string) (data list_rapat) {
	agdlr := db.Collection(col)
	filter := bson.M{"agendalistrapat": Agenda_ListRapat}
	err := agdlr.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdataListbyagdlr: %v\n", err)
	}
	return data
}

func GetDataList(nama_listrapat string, db *mongo.Database, col string) (data []list_rapat) {
	rapat := db.Collection(col)
	filter := bson.M{"namalistrapat": nama_listrapat}
	cursor, err := rapat.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("getDataList: %v\n", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func DeleteDataList(lokasi_listrapat string, db *mongo.Database, col string) (data list_rapat) {
	lkslr := db.Collection(col)
	filter := bson.M{"lokasilistrapat": lokasi_listrapat}
	err, _ := lkslr.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("DeleteDataList : %v\n", err)
	}
	fmt.Println("Succes Delete data")
	return data
}