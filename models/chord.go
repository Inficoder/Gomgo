package models

import (
	"Gomgo/middleware/logging"
	"Gomgo/pkg/setting"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Chord struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string 	   `bson:"name"`
	Singer string`bson:"singer"`
	Url string `bson:"url"`
	LocateWay int `bson:"locate_way"`
	Count int `bson:"count"`
	Kind string `bson:"kind"`
	//CreatedBy string `bson:"created_by"`
	Key string `bson:"key"`
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

func GetChordsCountByCondition(kind string,key string) int{
	var chords []Chord
	con := GetDataBase().C("chord")
	//var count int
	if kind == "thrum"{
		kind = "指弹"
	}else{
		kind = "弹唱"
	}
	if key != "All" {
		con.Find(bson.M{"kind":kind,"key":key}).All(&chords)
	}else{
		con.Find(bson.M{"kind":kind}).All(&chords)
	}
	return len(chords)
}


func GetChordsList(pageNum int,pageSize int) []Chord  {
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

func GetChordsListWithCondition(pageNum int,pageSize int,kind string,key string) ([]Chord ) {
	var chords []Chord
	con := GetDataBase().C("chord")
	if kind == "thrum"{
		kind = "指弹"
	}else{
	kind = "弹唱"
	}
	if key != "All" {
		con.Find(bson.M{"kind":kind,"key":key}).Limit(pageSize).Skip(pageSize*(pageNum-1)).All(&chords)
	}else{
		con.Find(bson.M{"kind":kind}).Limit(pageSize).Skip(pageSize*(pageNum-1)).All(&chords)
	}
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

func InsertChord(name string,singer string, locate_way int,kind string,key string,count int) (err error) {
	//url string,cover_url string,created_by string,modified_by string,modified_on int
	con := GetDataBase().C("chord")
	err = con.Insert(&Chord{ID: bson.NewObjectId(), Name: name,Singer:singer,Kind:kind,Url:name,LocateWay:locate_way,Count:count,Key:key,CoverUrl:"chords_cover/"+name,ModifiedBy:setting.User,ModifiedOn:time.Now().Unix()})
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
