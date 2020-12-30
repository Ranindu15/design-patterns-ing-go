package main

type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}
type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}
type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}
type house struct {
	windowType string
	doorType   string
	floor      int
}

func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	}
	if builderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}
func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *normalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *normalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *normalBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}