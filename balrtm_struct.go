package package_rtm

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type list_rapat struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_ListRapat    string             `bson:"nama_listrapat,omitempty" json:"nama_listrapat,omitempty"`
	Tanggal_ListRapat string             `bson:"tanggal_listrapat,omitempty" json:"tanggal_listrapat,omitempty"`
	Lokasi_ListRapat  string             `bson:"lokasi_listrapat,omitempty" json:"lokasi_listrapat,omitempty"`
	Agenda_ListRapat  string             `bson:"agenda_listrapat,omitempty" json:"agenda_listrapat,omitempty"`
}