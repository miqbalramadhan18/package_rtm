package package_rtm

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "ListRapat",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertDataList(t *testing.T) {
	Nama_ListRapat := "Rapat A Departemen X" 
	Tanggal_ListRapat := "2023-07-20"
	Lokasi_ListRapat := "Ruangan A Gedung X"
	Agenda_ListRapat := "Pembahasan Proyek"
	hasil := InsertDataList(MongoConn, Nama_ListRapat, Tanggal_ListRapat, Lokasi_ListRapat, Agenda_ListRapat)
	fmt.Println(hasil)
}

func TestGetDataList(t *testing.T) {
	Nama_ListRapat := "Rapat A Departemen X"
	hasil := GetDataList(Nama_ListRapat, MongoConn, "list_rapat")
	fmt.Println(hasil)
}

func TestGetDataListFromAganda(t *testing.T) {
	Agenda_ListRapat := "Pembahasan Proyek"
	hasil := GetDataListFromAganda(Agenda_ListRapat, MongoConn, "List_Rapat")
	fmt.Println(hasil)
}

func TestDeleteDataList(t *testing.T) {
	lokasi_listrapat := "Ruangan A Gedung Y"
	hasil := DeleteDataList(lokasi_listrapat, MongoConn, "List_rapat")
	fmt.Println(hasil)
}