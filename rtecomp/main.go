package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	inFileName  = flag.String("i", "", "Specifies the name of the source xml file to be compiled.")
	outLocation = flag.String("o", "", "Specifies the name of the directory to put output files. If blank, uses current directory")
	language    = flag.String("l", "xml", "The output language")
)

func main() {
	flag.Parse()

	*outLocation = strings.TrimSuffix(*outLocation, "/")
	*outLocation = strings.TrimSuffix(*outLocation, "\\")

	if *inFileName == "" {
		fmt.Println("You need to specify a file or directory name to compile! Check out -help for options")
		return

	}

	sourceFile, err := ioutil.ReadFile(*inFileName)
	if err != nil {
		fmt.Printf("Error reading file '%s' for conversion: %s\n", *inFileName, err.Error())
		return
	}

	fmt.Println(sourceFile)
	
	// for _, output := range sourceFile {
	// 	fmt.Printf("Writing %s.%s\n", output.Name, output.Extension)

	// 	err = ioutil.WriteFile(fmt.Sprintf("%s%c%s.%s", *outLocation, os.PathSeparator, output.Name, output.Extension), output.Contents, 0644)
	// 	if err != nil {
	// 		fmt.Println("Error during file write:", err.Error())
	// 		return
	// 	}
	// }

}
