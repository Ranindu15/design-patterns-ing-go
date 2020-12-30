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
