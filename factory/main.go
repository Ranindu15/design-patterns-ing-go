package main

import "fmt"

func main() {
	var result fPhone
	var phoneBrand string
	fmt.Scanln(&phoneBrand)
	result = getPhone(phoneBrand)
	fmt.Println("Your Phone Details :")
	fmt.Println("Name: ",result.getName())
	fmt.Println("Price: ", result.getPrice())
}

func getPhone(name string) (fPhone){
	if name == "Apple"{
		return newApple()
	}else if name == "Samsung"{
		return newSamsung()
	}else if name == "Xiomi"{
		return newXiomi()
	}else{
		return nil
	}
}

type fPhone interface {
	setName(name string)
	setPrice(price string)
	getName() string
	getPrice() string
}

func (fptr *phone) setName(pName string){
	fptr.name = pName
}
func (fptr *phone) setPrice(pPrice string){
	fptr.price = pPrice
}
func (fptr *phone) getName() string{
	return fptr.name
}
func (fptr *phone) getPrice() string{
	return fptr.price
}

type phone struct {
	name string
	price string
}

type apple struct {
	phone
}
type samsung struct {
	phone
}
type xiomi struct {
	phone
}

func newApple() fPhone{
	return &apple{
		phone{
			name: "Apple",
			price: "999$",
		},
	}
}
func newSamsung() fPhone{
	return &samsung{
		phone{
			name: "Samsung",
			price: "499$",
		},
	}
}
func newXiomi() fPhone{
	return &xiomi{
		phone{
			name: "Xiomi",
			price: "299$",
		},
	}
}



