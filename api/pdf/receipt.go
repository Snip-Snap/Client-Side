package pdf

import (
	"api/model"
	"fmt"
	"os/exec"
	"strconv"
	"time"

	"os"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func Receipt(rdata []*model.ReceiptData) string {

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	//m.SetPageMargins(10, 15, 10)

	grayish := getGrayishColor()

	m.RegisterHeader(func() {

		m.Row(40, func() {
			m.Col(4, func() {
				_ = m.FileImage("images/snipsnap.png", props.Rect{
					Left:    20,
					Center:  false,
					Percent: 80,
				})
			})

			m.ColSpace(3)

			m.Col(4, func() {
				current := time.Now().UTC()
				daterec := "Receipt #" + rdata[0].ApptID + " * " + current.Format("April 02 2006")
				m.Text(daterec, props.Text{
					Top:         8,
					Size:        10,
					Style:       consts.BoldItalic,
					Align:       consts.Right,
					Extrapolate: false,
				})

				addr := rdata[0].Shopstreetaddr + " " + rdata[0].ShopCity + ", California"
				m.Text(rdata[0].ShopName, props.Text{
					Top: 12,

					Size:  10,
					Align: consts.Right,
				})
				m.Text(addr, props.Text{
					Top: 15,

					Size:  10,
					Align: consts.Right,
				})
			})
		})
		m.Row(20, func() {
			m.Col(3, func() {
				m.Text("Prepared For", props.Text{
					Size:  10,
					Top:   12,
					Align: consts.Right,
				})
				client := rdata[0].Clientfirstname + " " + rdata[0].Clientlastname
				m.Text(client, props.Text{
					Size:  12,
					Top:   18,
					Align: consts.Right,
					Style: consts.Bold,
				})
			})

			m.ColSpace(6)

			m.Col(3, func() {
				m.Text("Payment To", props.Text{
					Size:  10,
					Top:   12,
					Align: consts.Left,
					//Extrapolate: true,
				})
				barber := rdata[0].Barberfirstname + " " + rdata[0].Barberlastname
				m.Text(barber, props.Text{
					Size:  12,
					Top:   18,
					Align: consts.Left,
					Style: consts.Bold,
				})
			})
		})
	})

	m.RegisterFooter(func() {
		m.Row(20, func() {
			m.Col(12, func() {

			})
		})

	})
	m.Line(10)
	total := 0.0
	header := []string{"DESCRIPTION", "QTY", "PRICE", "SUBTOTAL"}
	content := [][]string{}
	for _, element := range rdata {
		row := []string{element.ServiceName, "1", element.Price, element.Price}
		content = append(content, row)
		value, err := strconv.ParseFloat(element.Price, 64)
		if err == nil {
			total += value
		}

	}
	m.TableList(header, content, props.TableList{
		HeaderProp: props.TableListContent{
			Size: 9,
		},
		ContentProp: props.TableListContent{
			Size: 8,
		},
		Align:                consts.Center,
		AlternatedBackground: &grayish,
		HeaderContentSpace:   1,
		Line:                 false,
	})

	m.Row(20, func() {
		m.ColSpace(7)
		m.Col(2, func() {
			m.Text("Total:", props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Right,
			})
		})
		m.Col(3, func() {
			temp := fmt.Sprintf("$ %.2f", total)
			m.Text(temp, props.Text{
				Top:   5,
				Style: consts.Bold,
				Size:  8,
				Align: consts.Center,
			})
		})
	})

	m.Line(10)
	m.Row(15, func() {
		m.ColSpace(2)
		m.Col(7, func() {
			_ = m.Barcode("5123.151231.512314.1251251.123215", props.Barcode{
				Percent: 100,
				Left:    3,
				Proportion: props.Proportion{
					Width:  20,
					Height: 2,
				},
			})
			m.Text("5123.151231.512314.1251251.123215", props.Text{
				Top:    12,
				Family: "",
				Style:  consts.Bold,
				Size:   9,
				Align:  consts.Center,
			})
		})

	})

	err := m.OutputFileAndClose("./billing.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}
	_, _ = exec.Command("sh", "-c", "aws s3 cp ./"+rdata[0].ApptID+".pdf s3://dbwangpdf --acl public-read").Output()

	return "https://dbwangpdf.s3-us-west-1.amazonaws.com/" + rdata[0].ApptID + ".pdf"

}
func getGrayishColor() color.Color {
	return color.Color{
		Red:   248,
		Green: 248,
		Blue:  248,
	}
}
