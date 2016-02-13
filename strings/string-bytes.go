package main

import "fmt"
import "unicode/utf8"

func ShowString(sample string) {
	fmt.Println("Println:", sample)

	fmt.Printf("Byte loop:\n  ")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Rune loop:")
	for index, runeValue := range sample {
		fmt.Printf("  %#U starts at byte position %d\n", runeValue, index)
		fmt.Printf("  %d\n", runeValue)

	}

	fmt.Printf("Byte loop decoding with utf8:\n")
	for i, w := 0, 0; i < len(sample); i += w {
		runeValue, width := utf8.DecodeRuneInString(sample[i:])
		fmt.Printf("  %#U starts at byte position %d\n", runeValue, i)
		w = width
	}

	fmt.Printf("Printf with %%x %x\n", sample)
	fmt.Printf("Printf with %% x % x\n", sample)
	fmt.Printf("Printf with %%q %q\n", sample)
	fmt.Printf("Printf with %%+q %+q\n", sample)
}

func main() {
	//const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	const sample = "⌘日本語"
	ShowString(sample)
}
