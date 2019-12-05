package data_layer

import (
	"github.com/globalsign/mgo"
)

type IMgoDataLayer interface {
	Insert(something ...interface{}) error
	Update(condition, something interface{}) error
	UpdateAll(condition, something interface{}) error
	Count(condition interface{}) (int, error)
	FindAll(condition, result interface{}, offset, limit *int, sort []string) error
	FindOne(condition, result interface{}) error
	Delete(something interface{}) error
	DeleteAll(condition interface{}) error
	Upsert(condition, something interface{}) error
	UpdateBulk(conditions []interface{}, somethings []interface{}) (*mgo.BulkResult, error)
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

func (dl *mgoDataLayer) UpdateAll(condition, something interface{}) error {
	s, close := dl.getConn()
	defer close()
	_, err := s.UpdateAll(condition, something)
	if err != nil {
		return err
	}
	return nil
}

func (dl *mgoDataLayer) Upsert(condition, something interface{}) error {
	s, close := dl.getConn()
	defer close()
	_, err := s.Upsert(condition, something)
	return err
}

func (dl *mgoDataLayer) UpdateBulk(conditions []interface{}, somethings []interface{}) (*mgo.BulkResult, error) {
	s, close := dl.getConn()
	defer close()

	bulk := s.Bulk()
	for i, _ := range somethings {
		bulk.UpdateAll(conditions[i], somethings[i])
	}

	result, err := bulk.Run()
	if err != nil {
		return result, err
	}

	return result, nil
}

func (dl *mgoDataLayer) Count(condition interface{}) (int, error) {
	s, close := dl.getConn()
	defer close()
	return s.Find(condition).Count()
}

func (dl *mgoDataLayer) FindAll(condition, result interface{}, offset, limit *int, sort []string) error {
	s, close := dl.getConn()
	defer close()
	query := s.Find(condition)
	if offset != nil {
		query = query.Skip(*offset)
	}
	if limit != nil {
		query = query.Limit(*limit)
	}
	if len(sort) > 0 {
		query = query.Sort(sort...)
	}
	return query.All(result)
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

func (dl *mgoDataLayer) DeleteAll(condition interface{}) error {
	s, close := dl.getConn()
	defer close()
	_, err := s.RemoveAll(condition)
	if err != nil {
		return err
	}
	return nil
}
