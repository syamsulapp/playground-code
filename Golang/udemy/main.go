package main

import "fmt"

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

	//interfaces with struct of update url
	// url := HitUrl{Url: "https://ecif.eng.ui.ac.id"}
	// url.UpdateUrl("https://www.google.com")
	// PrintUrl(url)

	// Url := StoreUrl{
	// 	Url: "https://www.google.com",
	// }
	// Url.ChangeUrl("https://ecif.eng.ui.ac.id")
	// Url.PrintResponseUrl()

	//assigned interfaces
	// t := Triangle{Base: 10, Height: 10}
	// s := Square{SideLength: 10}

	// PrintArea(t)
	// PrintArea(s)

	//channel and go routine
	// ListUrl()

	//project calculator
	var (
		x    int
		y    int
		tipe string
	)
	fmt.Print("masukan nilai x:")
	fmt.Scan(&x)
	fmt.Print("masukan nilai y:")
	fmt.Scan(&y)
	fmt.Print("masukan tipe calculator:")
	fmt.Scan(&tipe)

	tc := InputOfTypeCalculator{x: x, y: y}
	switch tipe {
	case "penjumlahan":
		PrintPenjumlahan(tc)
		break
	case "pengurangan":
		PrintPengurangan(tc)
		break
	case "perkalian":
		PrintPerkalian(tc)
		break
	case "pembagian":
		PrintPembagian(tc)
		break
	}

}
