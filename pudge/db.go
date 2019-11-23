package pudge

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os/exec"
)

const DbPath = "item.db"

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
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	i.ID = string(out)
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

func (i Item) Save() error {
	err := db.Update(func(tx *bolt.Tx) error {
		people, err := tx.CreateBucketIfNotExists([]byte("item"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		enc, err := i.encode()
		if err != nil {
			return fmt.Errorf("could not encode Person %s: %s", i.ID, err)
		}
		err = people.Put([]byte(i.ID), enc)
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
