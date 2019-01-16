package models

import (
	"Gomgo/middleware/logging"
	"Gomgo/pkg/setting"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Chord struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string 	   `bson:"name"`
	Singer string`bson:"singer"`
	Url string `bson:"url"`
	Kind string `bson:"kind"`
	//CreatedBy string `bson:"created_by"`
	ModifiedBy string `bson:"modified_by"`
	ModifiedOn int64 `bson:"modified_on"`
	CoverUrl string `bson:"cover_url"`
}


func IsExistChordByName(name string) bool{
	var chord Chord
	con := GetDataBase().C("chord")
	con.Find(bson.M{"name": name}).One(&chord)
	if chord.Name != "" {
		return true
	}
	return false
}

func IsExistChordById(Id string) bool{
	var chord Chord
	con := GetDataBase().C("chord")
	objectId := bson.ObjectIdHex(Id)
	con.Find(bson.M{"_id":objectId}).One(&chord)
	fmt.Println(chord)
	if chord.Name != "" {
		return true
	}
	return false
}

func GetChordsCount() int{
	con := GetDataBase().C("chord")
	//var count int
	count , err := con.Count()
	if err != nil{
		logging.Error(err)
	}
	return count
}


func GetChordsList(pageNum int,pageSize int) ([]Chord ) {
	var chords []Chord
	con := GetDataBase().C("chord")
	con.Find(bson.M{}).Limit(pageSize).Batch(pageNum).All(&chords)
	//if err := con.Find(bson.M{}).All(&chords); err != nil {
	//	if err.Error() != GetErrNotFound().Error() {
	//		return chords
	//	}
	//
	//}
	return chords
}

func GetChord(Id string) (Chord) {
	var chord Chord
	objectId := bson.ObjectIdHex(Id)
	con := GetDataBase().C("chord")
	if err := con.Find(bson.M{"_id":objectId}).One(&chord); err != nil {
		return chord
	}
	return chord
}

func InsertChord(name string,singer string, kind string) (err error) {
	//url string,cover_url string,created_by string,modified_by string,modified_on int
	con := GetDataBase().C("chord")
	err = con.Insert(&Chord{ID: bson.NewObjectId(), Name: name,Singer:singer,Kind:kind,Url:name,CoverUrl:name,ModifiedBy:setting.User,ModifiedOn:time.Now().Unix()})
	return
}

func RemoveChord(Id string)(err error){

	if IsExistChordById(Id) == false{
		logging.Fatal("name不存在")
		return
	}
	con := GetDataBase().C("chord")
	err = con.Remove(bson.M{"_id":Id})
	return
}


//func UpdateChord(name string)(err error){
//	if IsExist(name) == false{
//		log.Fatal("name不存在")
//		return
//	}
//	con := GetDataBase().C("chord")
//	err = con.Update()
//}
