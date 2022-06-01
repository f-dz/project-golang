package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data := map[string]interface{}{
		"CUSTOMER_NAME":    "ANASTASIA",
		"MAINTENANCE_TYPE": "UPGRADE LAYANAN",
		"EXPIRY_DATE":      "02-01-2006",
		"SUBJECT":          "OCBC_UPGRADE LAYANAN",
		"EMAIL":            "ADMIN@GMAIL.COM",
	}
	var (
		err          error
		templateByte []byte
	)

	if templateByte, err = ioutil.ReadFile("./html-input/CustomerAuthorizationMaintenance.html"); err != nil {
		log.Fatal(err)
	}

	change := strings.NewReplacer(
		"{{CUSTOMER_NAME}}", data["CUSTOMER_NAME"].(string),
		"{{MAINTENANCE_TYPE}}", data["MAINTENANCE_TYPE"].(string),
		"{{EXPIRY_DATE}}", data["EXPIRY_DATE"].(string),
		"{{SUBJECT}}", data["SUBJECT"].(string),
		"{{EMAIL}}", data["EMAIL"].(string),
	)

	replaced := change.Replace(string(templateByte))

	if err = ioutil.WriteFile("./html-output/CustomerAuthorizationMaintenance.html", []byte(replaced), 0600); err != nil {
		log.Fatal(err)
	}

}
