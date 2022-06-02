package main

import (
	"io/ioutil"
	"log"
	"strings"

	pdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

const filename = "CustomerAuthorizationMaintenance"

func main() {
	data := map[string]interface{}{
		"CUSTOMER_NAME":    "ANASTASIA",
		"MAINTENANCE_TYPE": "UPGRADE LAYANAN",
		"EXPIRY_DATE":      "02-01-2006",
		"SUBJECT":          "OCBC_UPGRADE LAYANAN",
		"EMAIL":            "ADMIN@GMAIL.COM",
	}

	content := generateHTML(data)
	generatePDF(content)
}

func generateHTML(data map[string]interface{}) string {
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

	return replaced
}

func generatePDF(content string) {
	var err error

	pdfg, err := pdf.NewPDFGenerator()
	checkError(err)

	pdfg.AddPage(pdf.NewPageReader(strings.NewReader(content)))

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
