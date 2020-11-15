package main

import (
	"encoding/base32"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

// Error codes
const (
	errorReadingStdin int = 1
	errorDecoding
	errorUnknownMode
)

// Modes
const (
	Encoding int = iota
	Decoding
)

func main() {
	mode := Encoding

	decoding := flag.Bool("d", false, "Decode mode")
	base32Mode := flag.Bool("32", false, "base32 mode")

	flag.Parse()

	if *decoding {
		mode = Decoding
	}

	// Read stdin
	input, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		fmt.Println("Error reading stdin : " + err.Error())
		os.Exit(errorReadingStdin)
	}

	var output string
	if mode == Encoding {
		// Switch between encode modes.
		if *base32Mode {
			output = base32.StdEncoding.EncodeToString(input)
		} else {
			output = base64.StdEncoding.EncodeToString(input)
		}

		// Append a newline onto the end of the output string
		// when we are decoding so output terminal isn't messed up
		output += "\n"
	} else if mode == Decoding {
		var buff []byte
		var err error

		// Switch between decode modes
		if *base32Mode {
			buff, err = base32.StdEncoding.DecodeString(string(input))
		} else {
			buff, err = base64.StdEncoding.DecodeString(string(input))
		}

		if err != nil {
			fmt.Println("Error decoding : " + err.Error())
			os.Exit(errorDecoding)
		}

		output = string(buff)
	} else {
		fmt.Println("Error unknown mode")
		os.Exit(errorUnknownMode)
	}

	// Display output.
	fmt.Printf("%v", string(output))
}
