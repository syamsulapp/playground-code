package main

func main() {
	//save file
	// cards := newDeck()
	// cards.saveToFile("my_cards")

	//read file
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// cards := newDeck()
	// cards.shuffle()
	// cards.print()

	//(declare struct basic)
	// alex := Person{FirstName: "alex", LastName: "anderson"}
	// fmt.Println(alex)

	//(struct with handle via variable)
	// var alex Person
	// var contact ContactInfo
	// fmt.Println(alex)
	// alex.FirstName = "alex"
	// alex.LastName = "anderson"
	// contact.Email = "alex@gmail.com"
	// contact.ZipCode = 93233
	// fmt.Printf("%+v", alex)
	// fmt.Printf("%+v", contact)

	// (embedding struct)
	// Alex := Person{
	// 	FirstName: "alex",
	// 	LastName:  "anderson",
	// 	ContactInfo: ContactInfo{
	// 		Email:   "alex@gmail.com",
	// 		ZipCode: 93233,
	// 	},
	// }
	// Alex.ChangeName("alexander")
	// Alex.print()

	//map
	// Mapping()

	//interfaces
	// eb := EnglishBot{}
	// sb := SpanishBot{}

	// PrintGreeting(eb)
	// PrintGreeting(sb)

	url := HitUrl{}
	PrintUrl(url)
}
