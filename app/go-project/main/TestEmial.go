package main

import (
	"log"
	"github.com/emersion/go-imap/client"
	"github.com/emersion/go-imap"
)

//func get1st(a, b interface{}) interface{} {
//	return a
//}

func main() {
	//conn, _ := net.Dial("tcp", "imap.exmail.qq.com:993")
	//client, _ := imap.NewClient(conn, "imap.exmail.qq.com")
	//defer func() {
	//	client.Close()
	//}()
	//
	//err := client.Login("xinwu-yang@cxria.com", "Chengxun@1806")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//client.Select(imap.Inbox)
	//ids, err := client.Search("unseen")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for _, id := range ids {
	//	msg, err := client.GetMessage(id)
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println("Subject:", msg.Header["Subject"][0])
	//	body, _ := ioutil.ReadAll(msg.Body)
	//	r, _ := iconv.ConvertString(string(body), "GB2312", "UTF-8")
	//	fmt.Println("body:\n", r)
	//}
	//client.Logout()

	log.Println("Connecting to server...")
	c, err := client.DialTLS("imap.exmail.qq.com:993", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	// Don't forget to logout
	defer c.Logout()

	// Login
	if err := c.Login("xinwu-yang@cxria.com", "Chengxun@1806"); err != nil {
		log.Fatal(err)
	}
	log.Println("Logged in")
	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}

	if err := <-done; err != nil {
		log.Fatal(err)
	}

	// Select INBOX
	inBox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Flags for INBOX:", inBox.Flags)
	log.Println("mbox.Messages", inBox.Messages)
	seqSet := new(imap.SeqSet)
	seqSet.AddRange(1, inBox.Messages)
	messages := make(chan *imap.Message, 10)
	c.Fetch(seqSet,[]string{imap.EnvelopeMsgAttr},messages)
	for msg := range messages {
		log.Println("* " + msg.Envelope.Subject)
	}
}
