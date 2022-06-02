package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	var filename = "CustomerAuthorizationMaintenance"

	data := map[string]interface{}{
		"CUSTOMER_NAME":    "ANASTASIA",
		"MAINTENANCE_TYPE": "UPGRADE LAYANAN",
		"EXPIRY_DATE":      "02-01-2006",
		"SUBJECT":          "OCBC_UPGRADE LAYANAN",
		"EMAIL":            "ADMIN@GMAIL.COM",
	}

	generateHTML(filename, data)
	generatePDF(filename)
}

func generateHTML(filename string, data map[string]interface{}) {
	var (
		err          error
		templateByte []byte
	)

	templateByte, err = ioutil.ReadFile("./html-input/" + filename + ".html")
	checkError(err)

	change := strings.NewReplacer(
		"{{CUSTOMER_NAME}}", data["CUSTOMER_NAME"].(string),
		"{{MAINTENANCE_TYPE}}", data["MAINTENANCE_TYPE"].(string),
		"{{EXPIRY_DATE}}", data["EXPIRY_DATE"].(string),
		"{{SUBJECT}}", data["SUBJECT"].(string),
		"{{EMAIL}}", data["EMAIL"].(string),
	)

	replaced := change.Replace(string(templateByte))

	err = ioutil.WriteFile("./html-output/"+filename+".html", []byte(replaced), 0600)
	checkError(err)
}

func generatePDF(filename string) {
	var (
		err error
		f   *os.File
	)

	pdfg, err := pdf.NewPDFGenerator()
	checkError(err)

	if f, err = os.Open("./html-output/" + filename + ".html"); f != nil {
		defer f.Close()
	}
	checkError(err)

	pdfg.AddPage(pdf.NewPageReader(f))

	pdfg.Orientation.Set(pdf.OrientationPortrait)
	pdfg.Dpi.Set(400)

	err = pdfg.Create()
	checkError(err)

	err = pdfg.WriteFile("./pdf-output/" + filename + ".pdf")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
