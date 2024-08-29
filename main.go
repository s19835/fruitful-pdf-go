package main

import (
	"fmt"
	"os"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildingHeader(m)

	err := m.OutputFileAndClose("pdfs/example-pdf.pdf")
	if err != nil {
		fmt.Println("could not save the PDF: ", err)
		os.Exit(1)
	}

	fmt.Println("PDF save successfully.")
}

func buildingHeader(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("images/freepik-jp.jpeg", props.Rect{
					Center:  true,
					Percent: 75,
				})

				if err != nil {
					fmt.Println("images file was not loaded: ", err)
				}
			})
		})

		m.Row(10, func() {
			m.Col(12, func() {
				m.Text("", props.Text{
					Top:   3,
					Style: consts.Bold,
					Align: consts.Center,
					Color: getDarkPurple(),
				})
			})
		})
	})
}

func getDarkPurple() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}
