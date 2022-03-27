package main

import (
	"fmt"

	"github.com/gosimple/slug"
)

func main() {
	fmt.Println(slug.Make(""))                    // Will print: ""
	fmt.Println(slug.Make("Beef 1kg"))            // Will print: ""
	fmt.Println(slug.Make("beef 1 kg"))           // Will print: ""
	fmt.Println(slug.Make("Beef 1kg (Regular)"))  // Will print: ""
	fmt.Println(slug.Make("Beef Regular - 1KG"))  // Will print: ""
	fmt.Println(slug.MakeLang("গরুর মাংস", "en")) // Will print: ""

	text := slug.Make("Hellö Wörld хелло ворлд")
	fmt.Println(text) // Will print: "hello-world-khello-vorld"

	someText := slug.Make("影師")
	fmt.Println(someText) // Will print: "ying-shi"

	enText := slug.MakeLang("This & that", "en")
	fmt.Println(enText) // Will print: "this-and-that"

	deText := slug.MakeLang("Diese & Dass", "de")
	fmt.Println(deText) // Will print: "diese-und-dass"

	slug.Lowercase = false // Keep uppercase characters
	deUppercaseText := slug.MakeLang("Diese & Dass", "de")
	fmt.Println(deUppercaseText) // Will print: "Diese-und-Dass"

	slug.CustomSub = map[string]string{
		"water": "sand",
	}
	textSub := slug.Make("water is hot")
	fmt.Println(textSub) // Will print: "sand-is-hot"
}
