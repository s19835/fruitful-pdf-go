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
	tableOfContent(m)

	err := m.OutputFileAndClose("pdfs/example-pdf.pdf")
	if err != nil {
		fmt.Println("could not save the PDF: ", err)
		os.Exit(1)
	}

	fmt.Println("PDF save successfully.")
}

// Function to build the heading of the pdf
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
	})

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("A vibrant, stylized sunset scene featuring traditional East Asian architecture against a backdrop of a futuristic cityscape.", props.Text{
				Top:   1,
				Style: consts.Bold,
				Align: consts.Center,
				Color: getDarkPurple(),
			})
		})
	})
}

// function to create a table of contents
func tableOfContent(m pdf.Maroto) {
	tableHeadings := []string{"Furit", "Description", "Price"}                                         // demi headings
	contents := [][]string{{"Apple", "Red and Juciy", "2.00"}, {"Orange", "Orange and Juciy", "3.00"}} // demi contents
	lightPurple := getLightPurple()

	m.SetBackgroundColor(getTeal())

	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Table of Content", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})

	m.SetBackgroundColor(color.NewWhite())

	m.TableList(tableHeadings, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{3, 7, 2},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{3, 7, 2},
		},
		Align:                consts.Left,
		HeaderContentSpace:   1,
		Line:                 false,
		AlternatedBackground: &lightPurple,
	})
}

func getLightPurple() color.Color {
	return color.Color{
		Red:   210,
		Green: 200,
		Blue:  230,
	}
}

func getDarkPurple() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}

func getTeal() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}
