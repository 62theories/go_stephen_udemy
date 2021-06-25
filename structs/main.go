package main

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "jim",
		lastName:  "good",
		contact: contactInfo{
			email:   "email",
			zipCode: 141250,
		},
	}
	jim.updateName("jimmy")
}

func (pointerToPerson *person) updateName(newFirstname string) {
	(*pointerToPerson).firstName = newFirstname
}
