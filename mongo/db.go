package mongo

import "github.com/globalsign/mgo"

type DB struct {
	session *mgo.Session
	Url     string
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) Open() error {
	s, err := mgo.Dial(db.Url)
	if err != nil {
		return err
	}
	db.session = s
	return nil
}

func (db *DB) Close() {
	if db.session != nil {
		db.session.Close()
	}
}
