package main

import "time"

type topic struct {
	id        uint64
	title     string
	text      string
	createdAt time.Time
	comments  []string
}

type topicsApi interface {
	AddTopic(newTopic topic)
	GetTopic(id uint64) topic
	RemoveTopic(id uint16)
}

type myApi struct {
	storage map[uint64]topic
}

// Пример реализации метода GetTopic
func (api *myApi) GetTopic(id uint64) topic {
	if topic, ok := api.storage[id]; ok {
		return topic
	}

	return topic{}
}

// здесь необходимо реализовать метод AddTopic
func (api *myApi) AddTopic(newTopic topic) bool {
	if _, ok  := api.storage[newTopic.id]; ok {
		return false
	}
	api.storage[newTopic.id] = newTopic
	return true
}

// здесь необходимо реализовать метод RemoveTopic
func (api *myApi) RemoveTopic(id uint64) bool {
	if _, ok  := api.storage[id]; ok {
		delete(api.storage, id)
		return true
	}
	return false
}

func main() {

}
