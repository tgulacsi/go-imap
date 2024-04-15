package imapclient_test

import (
	"testing"

	"github.com/emersion/go-imap/v2"
)

func TestCreate(t *testing.T) {
	client, server := newClientServerPair(t, imap.ConnStateAuthenticated)
	defer client.Close()
	defer server.Close()

	name := "Test mailbox"
	if err := client.Create(name, nil).Wait(); err != nil {
		t.Fatalf("Create() = %v", err)
	}

	listCmd := client.List("", name, nil)
	mailboxes, err := listCmd.Collect()
	if err != nil {
		t.Errorf("List() = %v", err)
	} else if len(mailboxes) != 1 || mailboxes[0].Mailbox != name {
		t.Errorf("List() = %v, want exactly one entry with correct name", mailboxes)
	}
}
