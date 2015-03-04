package sim

import (
	"container/list"
	"math"

	"github.com/go-gl/mathgl/mgl32"
)

// PheromoneStorageItem models an item in the pheromone storage
type PheromoneStorageItem struct {
	Pheromone         *Pheromone
	AllElement        *list.Element
	PartitionElements [9]*list.Element
}

// PheromoneStorage models a pheromone storage
type PheromoneStorage struct {
	All        *list.List
	Partitions []*list.List
}

// NewPheromoneStorage build and returns a pheromone storage
func NewPheromoneStorage(max int) *PheromoneStorage {
	storage := &PheromoneStorage{
		All:        list.New(),
		Partitions: make([]*list.List, pheromoneIndexParts*pheromoneIndexParts),
	}
	for i := 0; i < pheromoneIndexParts*pheromoneIndexParts; i++ {
		storage.Partitions[i] = list.New()
	}
	return storage
}

// Len returns the storage length
func (storage *PheromoneStorage) Len() int {
	return storage.All.Len()
}

// Add inserts a pheromone in the storage
func (storage *PheromoneStorage) Add(pheromone *Pheromone) {
	item := &PheromoneStorageItem{
		Pheromone: pheromone,
	}
	item.AllElement = storage.All.PushBack(item)
	for i := 0; i < 9; i++ {
		partition := storage.Partition(pheromone.Position, (i%3)-1, (i/3)-1)
		if partition != nil {
			item.PartitionElements[i] = partition.PushBack(item)
		}
	}

	pheromone.PheromoneStorageItem = item
}

// Remove inserts a pheromone in the storage
func (storage *PheromoneStorage) Remove(pheromone *Pheromone) {
	storage.All.Remove(pheromone.PheromoneStorageItem.AllElement)
	for i := 0; i < 9; i++ {
		partition := storage.Partition(pheromone.Position, (i%3)-1, (i/3)-1)
		if partition != nil {
			partition.Remove(pheromone.PheromoneStorageItem.PartitionElements[i])
		}
	}

	pheromone.PheromoneStorageItem = nil
}

// Front inserts a pheromone in the storage
func (storage *PheromoneStorage) Front() *list.Element {
	return storage.All.Front()
}

// Partition returns the short list for the given pheromone
func (storage *PheromoneStorage) Partition(point mgl32.Vec2, offsetX, offsetY int) *list.List {
	bucketX := int(math.Min(math.Floor(float64(point.X())*pheromoneIndexParts), pheromoneIndexParts-1))
	bucketX += offsetX
	if bucketX < 0 || bucketX > pheromoneIndexParts-1 {
		return nil
	}
	bucketY := int(math.Min(math.Floor(float64(point.Y())*pheromoneIndexParts), pheromoneIndexParts-1))
	bucketY += offsetY
	if bucketY < 0 || bucketY > pheromoneIndexParts-1 {
		return nil
	}
	return storage.Partitions[(bucketX*pheromoneIndexParts)+bucketY]
}
