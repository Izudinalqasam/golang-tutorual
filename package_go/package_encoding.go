package packagego

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func MainPackageEncoding() {
	base64Encoding()
	csvReaderEncoding()
}

func base64Encoding() {
	value := "Hello world"
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Println(string(decoded))
}

func csvReaderEncoding() {
	csvString := "suha,25,jakarta \n" +
		"ajeng,26,bandung \n" +
		"joko,27,solo "

	reader := csv.NewReader(strings.NewReader(csvString))
	writer := csv.NewWriter(os.Stdout)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	_ = writer.Write([]string{"suha", "25", "jakarta"})
	_ = writer.Write([]string{"ajeng", "26", "bandung"})
	_ = writer.Write([]string{"joko", "27", "solo"})

	writer.Flush()
}
