package db

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/google/uuid"
	"log"
)

const DbPath = "item.db"
const ItemBucket = "item"

var db *bolt.DB

type Item struct {
	ID          string `json:"id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	InStock     int    `json:"in_stock"`
}

func (i *Item) GenerateUniqueId() {
	newID := uuid.New()
	item, _ := GetItem(newID.String())
	if item == nil {
		i.ID = newID.String()
	} else {
		i.GenerateUniqueId()
	}
}

func Open() {
	var err error
	db, err = bolt.Open(DbPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	item := Item{
		ID:          "1",
		Image:       "<img src=\"images/16.jpg\" />",
		Name:        "MacBook Pro",
		Price:       1110,
		Description: "Test",
		InStock:     1,
	}
	_ = item.Save()
}

func Close() {
	_ = db.Close()
}

func GetItem(id string) (*Item, error) {
	var i *Item
	err := db.View(func(tx *bolt.Tx) error {
		var err error
		b := tx.Bucket([]byte(ItemBucket))
		k := []byte(id)
		i, err = decode(b.Get(k))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Could not get Item ID %s", id)
		return nil, err
	}
	return i, nil
}

func (i Item) Save() error {
	err := db.Update(func(tx *bolt.Tx) error {
		item, err := tx.CreateBucketIfNotExists([]byte(ItemBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		enc, err := i.encode()
		if err != nil {
			return fmt.Errorf("could not encode Item %s: %s", i.ID, err)
		}
		err = item.Put([]byte(i.ID), enc)
		return err
	})
	return err
}

func Delete(id string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		item, err := tx.CreateBucketIfNotExists([]byte(ItemBucket))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		err = item.Delete([]byte(id))
		return err
	})
	return err
}

func (i *Item) encode() ([]byte, error) {
	enc, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func decode(data []byte) (*Item, error) {
	var p *Item
	err := json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func List(bucket string) []*Item {
	var items []*Item
	_ = db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(bucket)).Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
			i, _ := decode(v)
			items = append(items, i)
		}
		return nil
	})
	return items
}
