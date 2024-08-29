package main

import (
	"fmt"
	"os"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	err := m.OutputFileAndClose("pdfs/example-pdf.pdf")
	if err != nil {
		fmt.Println("could not save the PDF: ", err)
		os.Exit(1)
	}

	fmt.Println("PDF save successfully.")
}
