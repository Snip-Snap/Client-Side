package pdf

import (
	"api/model"
	"fmt"
	"os"
	"os/exec"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func Receipt([]*model.ReceiptData) {

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)

	err := m.OutputFileAndClose("./billing.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}
	out, err := exec.Command("sh", "-c", "aws s3 cp ./billing.pdf s3://dbwangpdf --acl public-read").Output()
	fmt.Println(out)
	fmt.Println("here")
	fmt.Println(err)
}
