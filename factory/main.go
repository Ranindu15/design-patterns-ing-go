package main

import "fmt"

func main() {
	var phoneBrand string
	fmt.Scanln(&phoneBrand)
	result, err := getPhone(phoneBrand)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("Your Phone Details :",result.price, result.name)
}

type phone struct {
	name string
	price string
}
var err error = fmt.Errorf("not found")

type fPhone phone

func apple() phone{
	return phone{
		name: "Apple",
		price: "999$",
	}
}
func samsung() phone{
	return phone{
		name: "Samsung",
		price: "799$",
	}
}
func xiomi() phone{
	return phone{
		name: "Xiomi",
		price: "299$",
	}
}
func defultPhone() phone{
	return phone{
		name: "Not found",
		price: "000$",
	}
}

func getPhone(name string) (fPhone, error){
	if name == "Apple"{
		return fPhone(apple()),nil
	}else if name == "Samsung"{
		return fPhone(samsung()),nil
	}else if name == "xiomi"{
		return fPhone(xiomi()),nil
	}else{
		return fPhone(defultPhone()), err
	}
}