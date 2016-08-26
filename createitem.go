// 26 august 2016
package ews

import (
	"encoding/xml"
)

// https://msdn.microsoft.com/en-us/library/office/aa563009(v=exchg.140).aspx

type CreateItem struct {
	XMLName		struct{}	`xml:"m:CreateItem"`
	MessageDisposition		string	`xml:"MessageDisposition,attr"`
	SavedItemFolderId	SavedItemFolderId	`xml:"m:SavedItemFolderId"`
	Items			Messages		`xml:"m:Items"`
}

type Messages struct {
	Message		[]Message		`xml:"t:Message"`
}

type SavedItemFolderId struct {
	DistinguishedFolderId	DistinguishedFolderId	`xml:"t:DistinguishedFolderId"`
}

type DistinguishedFolderId struct {
	Id		string		`xml:"Id,attr"`
}

type Message struct {
	ItemClass		string		`xml:"t:ItemClass"`
	Subject		string		`xml:"t:Subject"`
	Body			Body			`xml:"t:Body"`
	Sender		OneMailbox	`xml:"t:Sender"`
	ToRecipients	XMailbox		`xml:"t:ToRecipients"`
}

type Body struct {
	BodyType		string	`xml:"BodyType,attr"`
	Body			[]byte	`xml:",chardata"`
}

type OneMailbox struct {
	Mailbox		Mailbox			`xml:"t:Mailbox"`
}

type XMailbox struct {
	Mailbox		[]Mailbox			`xml:"t:Mailbox"`
}

type Mailbox struct {
	EmailAddress		string		`xml:"t:EmailAddress"`
}

func BuildTextEmail(from string, to []string, subject string, body []byte) ([]byte, error) {
	c := new(CreateItem)
	c.MessageDisposition = "SendAndSaveCopy"
	c.SavedItemFolderId.DistinguishedFolderId.Id = "sentitems"
	m := new(Message)
	m.ItemClass = "IPM.Note"
	m.Subject = subject
	m.Body.BodyType = "Text"
	m.Body.Body = body
	m.Sender.Mailbox.EmailAddress = from
	mb := make([]Mailbox, len(to))
	for i, addr := range to {
		mb[i].EmailAddress = addr
	}
	m.ToRecipients.Mailbox = append(m.ToRecipients.Mailbox, mb...)
	c.Items.Message = append(c.Items.Message, *m)
	return xml.MarshalIndent(c, "", "  ")
}
