package db

import (
	"encoding/binary"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

type DB struct {
	Connection *bolt.DB
}

type Task struct {
	Key   int
	Value string
}

var taskBucket = []byte("tasks")

func Initialize(path string) *DB {
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: time.Second * 10})
	if err != nil {
		log.Fatalf("Error: cannot open database %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
	if err != nil {
		log.Printf("Can not create bucket %s:%v", taskBucket, err)
	}
	return &DB{
		Connection: db,
	}
}

func (db *DB) CreateTask(v string) (int, error) {
	var id int
	err := db.Connection.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		key := Itob(id64)
		id = int(id64)
		return b.Put(key, []byte(v))
	})
	if err != nil {
		return -1, err
	} else {
		return id, nil
	}
}

func (db *DB) ListTasks() ([]Task, error) {
	t := make([]Task, 0)
	db.Connection.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil && v != nil; k, v = c.Next() {
			t = append(t, Task{
				Key:   Btoi(k),
				Value: string(v),
			})
		}

		// or maybe use this:
		// b.ForEach(func(k, v []byte) error {

		// })
		return nil
	})
	return t, nil
}

func (db *DB) DoTask(ids []int) []error {
	var errs []error
	db.Connection.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		for _, id := range ids {
			err := b.Delete(Itob(uint64(id)))
			if err != nil {
				errs = append(errs, err)
			}
		}
		return nil
	})
	return errs
}

func Itob(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}

func Btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
