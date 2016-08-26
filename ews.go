// 26 august 2016
package ews

import (
	"bytes"
	"net/http"
)

// https://msdn.microsoft.com/en-us/library/office/dd877045(v=exchg.140).aspx
// https://arvinddangra.wordpress.com/2011/09/29/send-email-using-exchange-smtp-and-ews-exchange-web-service/
// https://msdn.microsoft.com/en-us/library/office/dn789003(v=exchg.150).aspx

var soapheader = `<?xml version="1.0" encoding="utf-8" ?>
<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:m="http://schemas.microsoft.com/exchange/services/2006/messages" xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Header>
    <t:RequestServerVersion Version="Exchange2007_SP1" />
  </soap:Header>
  <soap:Body>
`

func Issue(ewsAddr string, username string, password string, body []byte) (*http.Response, error) {
	b2 := []byte(soapheader)
	b2 = append(b2, body...)
	b2 = append(b2, "\n  </soap:Body>\n</soap:Envelope>"...)
	req, err := http.NewRequest("POST", ewsAddr, bytes.NewReader(b2))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "text/xml")
	client := new(http.Client)
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }
	return client.Do(req)
}
