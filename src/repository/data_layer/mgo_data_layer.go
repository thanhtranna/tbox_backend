package data_layer

import (
	"github.com/globalsign/mgo"
)

type IMgoDataLayer interface {
	Insert(something ...interface{}) error
	Update(condition, something interface{}) error
	FindOne(condition, result interface{}) error
	Delete(something interface{}) error
	Upsert(condition, something interface{}) error
}

type closeFunc func()

type mgoDataLayer struct {
	s              *mgo.Session
	collectionName string
}

func NewMgoDataLayer(s *mgo.Session, collectionName string) IMgoDataLayer {
	return &mgoDataLayer{s: s, collectionName: collectionName}
}

func (dl *mgoDataLayer) getConn() (*mgo.Collection, closeFunc) {
	s := dl.s.New()
	return s.DB("").C(dl.collectionName), s.Close
}

func (dl *mgoDataLayer) Update(condition, something interface{}) error {
	s, close := dl.getConn()
	defer close()
	return s.Update(condition, something)
}

func (dl *mgoDataLayer) Upsert(condition, something interface{}) error {
	s, close := dl.getConn()
	defer close()
	_, err := s.Upsert(condition, something)
	return err
}

func (dl *mgoDataLayer) FindOne(condition, result interface{}) error {
	s, close := dl.getConn()
	defer close()
	return s.Find(condition).One(result)
}

func (dl *mgoDataLayer) Insert(something ...interface{}) error {
	s, close := dl.getConn()
	defer close()

	return s.Insert(something...)
}

func (dl *mgoDataLayer) Delete(condition interface{}) error {
	s, close := dl.getConn()
	defer close()
	return s.Remove(condition)
}
