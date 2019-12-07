package data_layer_test

import (
	"testing"

	dl "tbox_backend/src/repository/data_layer"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

func TestInsert(t *testing.T) {
	session, _ := mgo.Dial("mongodb://localhost:27017/test")
	layer := dl.NewMgoDataLayer(session, "test")

	err := layer.Insert(map[string]interface{}{
		"field": 1,
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {
	session, _ := mgo.Dial("mongodb://localhost:27017/test")
	layer := dl.NewMgoDataLayer(session, "test")

	err := layer.Update(bson.D{{"field", 1}}, map[string]interface{}{
		"field": 2,
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestFindOne(t *testing.T) {
	session, _ := mgo.Dial("mongodb://localhost:27017/test")
	layer := dl.NewMgoDataLayer(session, "test")

	err := layer.FindOne(bson.D{{"field", 2}}, map[string]interface{}{
		"field": 2,
	})

	if err != nil {
		t.Fatal(err)
	}
}
