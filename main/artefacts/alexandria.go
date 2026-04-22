package alexandria

import (
	"fmt"
	"net/mail"
	"os"
	"strings"
)

type EmailHeaders struct {
	From                 string
	ReplyTo              string
	ReturnPath           string
	MessageID            string
	DKIMSignature        []string
	AuthUser             string
	ForwardingLoop       string
	HELO                 string
	EnvelopeFrom         string
	Received             []string
	AuthResults          []string
	SPF                  string
	DMARC                string
	XAuthenticatedSender string
}

func ReadFile(path string) (*EmailHeaders, error) {

	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer f.Close()

	msg, err := mail.ReadMessage(f)
	if err != nil {
		fmt.Println("Error parsing email:", err)
		return nil, err
	}

	h := &EmailHeaders{}

	h.From = msg.Header.Get("From")
	h.ReplyTo = msg.Header.Get("Reply-To")
	h.ReturnPath = msg.Header.Get("Return-Path")
	h.MessageID = msg.Header.Get("Message-ID")
	h.ForwardingLoop = msg.Header.Get("X-MS-Exchange-ForwardingLoop")
	h.AuthUser = msg.Header.Get("X-AuthUser")
	h.XAuthenticatedSender = msg.Header.Get("X-Authenticated-Sender")
	//X-Authenticated-Sender

	// Multi-value headers
	h.Received = msg.Header["Received"]
	h.DKIMSignature = msg.Header["Dkim-Signature"]
	h.AuthResults = msg.Header["Authentication-Results"]

	// Extract SPF and DMARC from Authentication-Results
	for _, r := range h.AuthResults {
		lower := strings.ToLower(r)
		if strings.Contains(lower, "spf=") {
			h.SPF = extractValue(r, "spf=")
		}
		if strings.Contains(lower, "dmarc=") {
			h.DMARC = extractValue(r, "dmarc=")
		}
	}

	// Extract HELO and EnvelopeFrom from Received headers
	for _, r := range h.Received {
		lower := strings.ToLower(r)
		if strings.Contains(lower, "helo=") {
			h.HELO = extractValue(r, "helo=")
		}
		if strings.Contains(lower, "envelope-from=") {
			h.EnvelopeFrom = extractValue(r, "envelope-from=")
		}
	}

	return h, nil
}

func extractValue(s, key string) string {
	lower := strings.ToLower(s)
	idx := strings.Index(lower, strings.ToLower(key))
	if idx == -1 {
		return ""
	}

	rest := s[idx+len(key):]
	rest = strings.Trim(rest, "\"")
	end := strings.IndexAny(rest, " ;\r\n")
	if end == -1 {
		return rest
	}

	return rest[:end]
}

func PrintHeaders(h *EmailHeaders) {
	fmt.Print("\n---------------------------")
	fmt.Println("\n\n|FROM:", h.From)
	fmt.Println("|REPLY-TO:", h.ReplyTo)
	fmt.Println("|RETURN-PATH:", h.ReturnPath)
	fmt.Println("|MESSAGE-ID:", h.MessageID)
	fmt.Println("|AUTH-USER:", h.AuthUser)
	fmt.Println("|FORWARDING-LOOP:", h.ForwardingLoop)
	fmt.Println("|HELO:", h.HELO)
	fmt.Println("|ENVELOPE-FROM:", h.EnvelopeFrom)
	fmt.Println("|SPF:", h.SPF)
	fmt.Println("|DMARC:", h.DMARC)
	fmt.Println("|X-AUTHENTICATED-SENDER:", h.XAuthenticatedSender)

	fmt.Println("\n|DKIM:")
	for i, d := range h.DKIMSignature {
		fmt.Printf("  [%d] %s\n", i+1, d[:min(80, len(d))])
	}

	fmt.Println("\n|RECEIVED CHAIN:\n")
	for i, r := range h.Received {
		fmt.Printf("  [%d] %s\n", i+1, r[:min(100, len(r))])
	}

	fmt.Println("\n|AUTH RESULTS:\n\n")
	for i, a := range h.AuthResults {
		fmt.Printf("  [%d] %s\n", i+1, a[:min(100, len(a))])
	}
	fmt.Println("\n\n---------------------------")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
