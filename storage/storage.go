package storage

import (
	"fmt"

	"github.com/mcmahonliam837/microcache/config"
)

type Storage struct {
	Data map[string][]byte
}

func (s *Storage) GetValue(key string, conf *config.Config) ([]byte, error) {

	fmt.Printf("Request: get key=%v\n", key)

	if conf.DiskMode {
		return s.getValueFromDisk(key, conf)
	}

	if value, ok := s.Data[key]; ok {
		return value, nil
	} else {
		return nil, fmt.Errorf("Key not found")
	}

}

func (s *Storage) Store(key string, value []byte, conf *config.Config) error {

	if conf.DiskMode {
		return s.storeToDisk(key, value, conf)
	} else {
		s.Data[key] = value
		return nil
	}
}

func (s *Storage) Delete(key string, conf *config.Config) error {
	if conf.DiskMode {
		return s.deleteFromDisk(key, conf)
	} else {
		delete(s.Data, key)
		return nil
	}
}

////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////// DiskMode //////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

func (s *Storage) getValueFromDisk(key string, conf *config.Config) ([]byte, error) {
	// TODO: This is not priority. Write when done with in-memory

	return nil, nil
}

func (s *Storage) storeToDisk(key string, value []byte, conf *config.Config) error {
	// TODO
	return nil
}

func (s *Storage) deleteFromDisk(key string, conf *config.Config) error {
	// TODO
	return nil
}
