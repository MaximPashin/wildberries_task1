package storage

import (
	"wb_task1/db"
	"wb_task1/entity"
	"wb_task1/validator"
)

type Storage struct {
	conn *db.Conn
	kv   *map[string]entity.Order
}

func New(conn *db.Conn) (*Storage, error) {
	return &Storage{conn, &map[string]entity.Order{}}, nil
}

func (s *Storage) Get(key string) (entity.Order, bool) {
	v, ok := (*s.kv)[key]
	return v, ok
}

func (s *Storage) Put(data entity.Order) error {
	err := validator.Validate(data)
	if err != nil {
		return err
	}
	err = s.conn.Write(data)
	if err != nil {
		return err
	}
	(*s.kv)[data.ID] = data
	return nil
}

func (s *Storage) Recovery() error {
	records, err := s.conn.LoadAll()
	if err != nil {
		return err
	}
	s.kv = records
	return nil
}
