// 26 august 2016
package main

import (
"os"
"encoding/xml"
)

// https://msdn.microsoft.com/en-us/library/office/aa563009(v=exchg.140).aspx

type CreateItem struct {
	SavedItemFolderId	SavedItemFolderId
	Items			Messages
}

type Messages struct {
	Message		[]Message
}

type SavedItemFolderId struct {
	DistinguishedFolderId	DistinguishedFolderId
}

type DistinguishedFolderId struct {
	Id		string		`xml:"Id,attr"`
}

type Message struct {
	ItemClass		string
	Subject		string
	Body			Body
	Sender		XMailbox
	ToRecipients	XMailbox
}

type Body struct {
	BodyType		string	`xml:"BodyType,attr"`
	Body			string	`xml:",chardata"`
}

type XMailbox struct {
	Mailbox		Mailbox
}

type Mailbox struct {
	EmailAddress		[]string
}

func main() {
	c := new(CreateItem)
	c.SavedItemFolderId.DistinguishedFolderId.Id = "sentitems"
	m := new(Message)
	m.ItemClass = "IPM.Note"
	m.Subject = "Daily Report"
	m.Body.BodyType = "Text"
	m.Body.Body = "(1) Handled customer issues, (2) Saved the world."
	m.Sender.Mailbox.EmailAddress = append(m.Sender.Mailbox.EmailAddress, "user1@example.com")
	m.ToRecipients.Mailbox.EmailAddress = append(m.ToRecipients.Mailbox.EmailAddress, "user2@example.com")
	m.ToRecipients.Mailbox.EmailAddress = append(m.ToRecipients.Mailbox.EmailAddress, "user3@example.com")
	c.Items.Message = append(c.Items.Message, *m)
	e := xml.NewEncoder(os.Stdout)
	e.Indent("", "  ")
	e.Encode(c)
}
