package main

import (
	"github.com/KeithAlt/psql-serializer-ums-demo/pkg/errorCheck"
	"github.com/KeithAlt/psql-serializer-ums-demo/pkg/ums_serializer"
	"log"
	"os"
)

func main() {
	log.Println("Starting psql serialization...")
	output := ums_serializer.Generate()

	file, err := os.Create("./data/out/.psql_migration")
	errorCheck.Fatal(err)

	defer file.Close()

	_, err = file.WriteString(output)
	errorCheck.Fatal(err)
	log.Println("Finished psql serialization")
}
