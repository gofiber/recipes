package mail

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"mime"
	"net"
	"net/mail"
	"net/textproto"
	"path/filepath"
	"time"
)

// Email represents an email message.
type Email struct {
	from        string
	sender      string
	replyTo     string
	returnPath  string
	recipients  []string
	headers     textproto.MIMEHeader
	parts       []part
	attachments []*file
	inlines     []*file
	Charset     string
	Encoding    encoding
	Error       error
	SMTPServer  *smtpClient
}

/*
SMTPServer represents a SMTP Server
If authentication is CRAM-MD5 then the Password is the Secret
*/
type SMTPServer struct {
	Authentication authType
	Encryption     encryption
	Username       string
	Password       string
	ConnectTimeout time.Duration
	SendTimeout    time.Duration
	Host           string
	Port           int
	KeepAlive      bool
}

//SMTPClient represents a SMTP Client for send email
type SMTPClient struct {
	Client      *smtpClient
	KeepAlive   bool
	SendTimeout time.Duration
}

// part represents the different content parts of an email body.
type part struct {
	contentType string
	body        *bytes.Buffer
}

// file represents the files that can be added to the email message.
type file struct {
	filename string
	mimeType string
	data     []byte
}

type encryption int

const (
	// EncryptionNone uses no encryption when sending email
	EncryptionNone encryption = iota
	// EncryptionSSL sets encryption type to SSL when sending email
	EncryptionSSL
	// EncryptionTLS sets encryption type to TLS when sending email
	EncryptionTLS
)

var encryptionTypes = [...]string{"None", "SSL", "TLS"}

func (encryption encryption) string() string {
	return encryptionTypes[encryption]
}

type encoding int

const (
	// EncodingNone turns off encoding on the message body
	EncodingNone encoding = iota
	// EncodingBase64 sets the message body encoding to base64
	EncodingBase64
	// EncodingQuotedPrintable sets the message body encoding to quoted-printable
	EncodingQuotedPrintable
)

var encodingTypes = [...]string{"binary", "base64", "quoted-printable"}

func (encoding encoding) string() string {
	return encodingTypes[encoding]
}

type contentType int

const (
	// TextPlain sets body type to text/plain in message body
	TextPlain contentType = iota
	// TextHTML sets body type to text/html in message body
	TextHTML
)

var contentTypes = [...]string{"text/plain", "text/html"}

func (contentType contentType) string() string {
	return contentTypes[contentType]
}

type authType int

const (
	// AuthPlain implements the PLAIN authentication
	AuthPlain authType = iota
	// AuthLogin implements the LOGIN authentication
	AuthLogin
	// AuthCRAMMD5 implements the CRAM-MD5 authentication
	AuthCRAMMD5
)

// NewMSG creates a new email. It uses UTF-8 by default. All charsets: http://webcheatsheet.com/HTML/character_sets_list.php
func NewMSG() *Email {
	email := &Email{
		headers:  make(textproto.MIMEHeader),
		Charset:  "UTF-8",
		Encoding: EncodingQuotedPrintable,
	}

	email.AddHeader("MIME-Version", "1.0")

	return email
}

//NewSMTPClient returns the client for send email
func NewSMTPClient() *SMTPServer {
	server := &SMTPServer{
		Authentication: AuthPlain,
		Encryption:     EncryptionNone,
		ConnectTimeout: 10 * time.Second,
		SendTimeout:    10 * time.Second,
	}
	return server
}

// GetError returns the first email error encountered
func (email *Email) GetError() error {
	return email.Error
}

// SetFrom sets the From address.
func (email *Email) SetFrom(address string) *Email {
	if email.Error != nil {
		return email
	}

	email.AddAddresses("From", address)

	return email
}

// SetSender sets the Sender address.
func (email *Email) SetSender(address string) *Email {
	if email.Error != nil {
		return email
	}

	email.AddAddresses("Sender", address)

	return email
}

// SetReplyTo sets the Reply-To address.
func (email *Email) SetReplyTo(address string) *Email {
	if email.Error != nil {
		return email
	}

	email.AddAddresses("Reply-To", address)

	return email
}

// SetReturnPath sets the Return-Path address. This is most often used
// to send bounced emails to a different email address.
func (email *Email) SetReturnPath(address string) *Email {
	if email.Error != nil {
		return email
	}

	email.AddAddresses("Return-Path", address)

	return email
}

// AddTo adds a To address. You can provide multiple
// addresses at the same time.
func (email *Email) AddTo(addresses ...string) *Email {
	if email.Error != nil {
		return email
	}

	email.AddAddresses("To", addresses...)

	return email
}

// AddCc adds a Cc address. You can provide multiple
// addresses at the same time.
func (email *Email) AddCc(addresses ...string) *Email {
	if email.Error != nil {
		return email
	}

	email.AddAddresses("Cc", addresses...)

	return email
}

// AddBcc adds a Bcc address. You can provide multiple
// addresses at the same time.
func (email *Email) AddBcc(addresses ...string) *Email {
	if email.Error != nil {
		return email
	}

	email.AddAddresses("Bcc", addresses...)

	return email
}

// AddAddresses allows you to add addresses to the specified address header.
func (email *Email) AddAddresses(header string, addresses ...string) *Email {
	if email.Error != nil {
		return email
	}

	found := false

	// check for a valid address header
	for _, h := range []string{"To", "Cc", "Bcc", "From", "Sender", "Reply-To", "Return-Path"} {
		if header == h {
			found = true
		}
	}

	if !found {
		email.Error = errors.New("Mail Error: Invalid address header; Header: [" + header + "]")
		return email
	}

	// check to see if the addresses are valid
	for i := range addresses {
		var address = new(mail.Address)
		var err error

		// ignore parse the address if empty
		if len(addresses[i]) > 0 {
			address, err = mail.ParseAddress(addresses[i])
			if err != nil {
				email.Error = errors.New("Mail Error: " + err.Error() + "; Header: [" + header + "] Address: [" + addresses[i] + "]")
				return email
			}
		} else {
			continue
		}

		// check for more than one address
		switch {
		case header == "From" && len(email.from) > 0:
			fallthrough
		case header == "Sender" && len(email.sender) > 0:
			fallthrough
		case header == "Reply-To" && len(email.replyTo) > 0:
			fallthrough
		case header == "Return-Path" && len(email.returnPath) > 0:
			email.Error = errors.New("Mail Error: There can only be one \"" + header + "\" address; Header: [" + header + "] Address: [" + addresses[i] + "]")
			return email
		default:
			// other address types can have more than one address
		}

		// save the address
		switch header {
		case "From":
			email.from = address.Address
		case "Sender":
			email.sender = address.Address
		case "Reply-To":
			email.replyTo = address.Address
		case "Return-Path":
			email.returnPath = address.Address
		default:
			// check that the address was added to the recipients list
			email.recipients, err = addAddress(email.recipients, address.Address)
			if err != nil {
				email.Error = errors.New("Mail Error: " + err.Error() + "; Header: [" + header + "] Address: [" + addresses[i] + "]")
				return email
			}
		}

		// make sure the from and sender addresses are different
		if email.from != "" && email.sender != "" && email.from == email.sender {
			email.sender = ""
			email.headers.Del("Sender")
			email.Error = errors.New("Mail Error: From and Sender should not be set to the same address")
			return email
		}

		// add all addresses to the headers except for Bcc and Return-Path
		if header != "Bcc" && header != "Return-Path" {
			// add the address to the headers
			email.headers.Add(header, address.String())
		}
	}

	return email
}

// addAddress adds an address to the address list if it hasn't already been added
func addAddress(addressList []string, address string) ([]string, error) {
	// loop through the address list to check for dups
	for _, a := range addressList {
		if address == a {
			return addressList, errors.New("Mail Error: Address: [" + address + "] has already been added")
		}
	}

	return append(addressList, address), nil
}

type priority int

const (
	// PriorityLow sets the email priority to Low
	PriorityLow priority = iota
	// PriorityHigh sets the email priority to High
	PriorityHigh
)

// SetPriority sets the email message priority. Use with
// either "High" or "Low".
func (email *Email) SetPriority(priority priority) *Email {
	if email.Error != nil {
		return email
	}

	switch priority {
	case PriorityLow:
		email.AddHeaders(textproto.MIMEHeader{
			"X-Priority":        {"5 (Lowest)"},
			"X-MSMail-Priority": {"Low"},
			"Importance":        {"Low"},
		})
	case PriorityHigh:
		email.AddHeaders(textproto.MIMEHeader{
			"X-Priority":        {"1 (Highest)"},
			"X-MSMail-Priority": {"High"},
			"Importance":        {"High"},
		})
	default:
	}

	return email
}

// SetDate sets the date header to the provided date/time.
// The format of the string should be YYYY-MM-DD HH:MM:SS Time Zone.
//
// Example: SetDate("2015-04-28 10:32:00 CDT")
func (email *Email) SetDate(dateTime string) *Email {
	if email.Error != nil {
		return email
	}

	const dateFormat = "2006-01-02 15:04:05 MST"

	// Try to parse the provided date/time
	dt, err := time.Parse(dateFormat, dateTime)
	if err != nil {
		email.Error = errors.New("Mail Error: Setting date failed with: " + err.Error())
		return email
	}

	email.headers.Set("Date", dt.Format(time.RFC1123Z))

	return email
}

// SetSubject sets the subject of the email message.
func (email *Email) SetSubject(subject string) *Email {
	if email.Error != nil {
		return email
	}

	email.AddHeader("Subject", subject)

	return email
}

// SetBody sets the body of the email message.
func (email *Email) SetBody(contentType contentType, body string) *Email {
	if email.Error != nil {
		return email
	}

	email.parts = []part{
		{
			contentType: contentType.string(),
			body:        bytes.NewBufferString(body),
		},
	}

	return email
}

// AddHeader adds the given "header" with the passed "value".
func (email *Email) AddHeader(header string, values ...string) *Email {
	if email.Error != nil {
		return email
	}

	// check that there is actually a value
	if len(values) < 1 {
		email.Error = errors.New("Mail Error: no value provided; Header: [" + header + "]")
		return email
	}

	switch header {
	case "Sender":
		fallthrough
	case "From":
		fallthrough
	case "To":
		fallthrough
	case "Bcc":
		fallthrough
	case "Cc":
		fallthrough
	case "Reply-To":
		fallthrough
	case "Return-Path":
		email.AddAddresses(header, values...)
	case "Date":
		if len(values) > 1 {
			email.Error = errors.New("Mail Error: To many dates provided")
			return email
		}
		email.SetDate(values[0])
	default:
		email.headers[header] = values
	}

	return email
}

// AddHeaders is used to add multiple headers at once
func (email *Email) AddHeaders(headers textproto.MIMEHeader) *Email {
	if email.Error != nil {
		return email
	}

	for header, values := range headers {
		email.AddHeader(header, values...)
	}

	return email
}

// AddAlternative allows you to add alternative parts to the body
// of the email message. This is most commonly used to add an
// html version in addition to a plain text version that was
// already added with SetBody.
func (email *Email) AddAlternative(contentType contentType, body string) *Email {
	if email.Error != nil {
		return email
	}

	email.parts = append(email.parts,
		part{
			contentType: contentType.string(),
			body:        bytes.NewBufferString(body),
		},
	)

	return email
}

// AddAttachment allows you to add an attachment to the email message.
// You can optionally provide a different name for the file.
func (email *Email) AddAttachment(file string, name ...string) *Email {
	if email.Error != nil {
		return email
	}

	if len(name) > 1 {
		email.Error = errors.New("Mail Error: Attach can only have a file and an optional name")
		return email
	}

	email.Error = email.attach(file, false, name...)

	return email
}

// AddAttachmentBase64 allows you to add an attachment in base64 to the email message.
// You need provide a name for the file.
func (email *Email) AddAttachmentBase64(b64File string, name string) *Email {
	if email.Error != nil {
		return email
	}

	if len(name) < 1 || len(b64File) < 1 {
		email.Error = errors.New("Mail Error: Attach Base64 need have a base64 string and name")
		return email
	}

	email.Error = email.attachB64(b64File, name)

	return email
}

// AddInline allows you to add an inline attachment to the email message.
// You can optionally provide a different name for the file.
func (email *Email) AddInline(file string, name ...string) *Email {
	if email.Error != nil {
		return email
	}

	if len(name) > 1 {
		email.Error = errors.New("Mail Error: Inline can only have a file and an optional name")
		return email
	}

	email.Error = email.attach(file, true, name...)

	return email
}

// attach does the low level attaching of the files
func (email *Email) attach(f string, inline bool, name ...string) error {
	// Get the file data
	data, err := ioutil.ReadFile(f)
	if err != nil {
		return errors.New("Mail Error: Failed to add file with following error: " + err.Error())
	}

	// get the file mime type
	mimeType := mime.TypeByExtension(filepath.Ext(f))
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	// get the filename
	_, filename := filepath.Split(f)

	// if an alternative filename was provided, use that instead
	if len(name) == 1 {
		filename = name[0]
	}

	if inline {
		email.inlines = append(email.inlines, &file{
			filename: filename,
			mimeType: mimeType,
			data:     data,
		})
	} else {
		email.attachments = append(email.attachments, &file{
			filename: filename,
			mimeType: mimeType,
			data:     data,
		})
	}

	return nil
}

// attachB64 does the low level attaching of the files but decoding base64 instead have a filepath
func (email *Email) attachB64(b64File string, name string) error {

	// decode the string
	dec, err := base64.StdEncoding.DecodeString(b64File)
	if err != nil {
		return errors.New("Mail Error: Failed to decode base64 attachment with following error: " + err.Error())
	}

	// get the file mime type
	mimeType := mime.TypeByExtension(name)
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	email.attachments = append(email.attachments, &file{
		filename: name,
		mimeType: mimeType,
		data:     dec,
	})

	return nil
}

// getFrom returns the sender of the email, if any
func (email *Email) getFrom() string {
	from := email.returnPath
	if from == "" {
		from = email.sender
		if from == "" {
			from = email.from
			if from == "" {
				from = email.replyTo
			}
		}
	}

	return from
}

func (email *Email) hasMixedPart() bool {
	return (len(email.parts) > 0 && len(email.attachments) > 0) || len(email.attachments) > 1
}

func (email *Email) hasRelatedPart() bool {
	return (len(email.parts) > 0 && len(email.inlines) > 0) || len(email.inlines) > 1
}

func (email *Email) hasAlternativePart() bool {
	return len(email.parts) > 1
}

// GetMessage builds and returns the email message
func (email *Email) GetMessage() string {
	msg := newMessage(email)

	if email.hasMixedPart() {
		msg.openMultipart("mixed")
	}

	if email.hasRelatedPart() {
		msg.openMultipart("related")
	}

	if email.hasAlternativePart() {
		msg.openMultipart("alternative")
	}

	for _, part := range email.parts {
		msg.addBody(part.contentType, part.body.Bytes())
	}

	if email.hasAlternativePart() {
		msg.closeMultipart()
	}

	msg.addFiles(email.inlines, true)
	if email.hasRelatedPart() {
		msg.closeMultipart()
	}

	msg.addFiles(email.attachments, false)
	if email.hasMixedPart() {
		msg.closeMultipart()
	}

	return msg.getHeaders() + msg.body.String()
}

// Send sends the composed email
func (email *Email) Send(client *SMTPClient) error {

	if email.Error != nil {
		return email.Error
	}

	if len(email.recipients) < 1 {
		return errors.New("Mail Error: No recipient specified")
	}

	msg := email.GetMessage()

	return send(email.from, email.recipients, msg, client)

}

// dial connects to the smtp server with the request encryption type
func dial(host string, port string, encryption encryption, config *tls.Config) (*smtpClient, error) {
	var conn net.Conn
	var err error

	address := host + ":" + port

	// do the actual dial
	switch encryption {
	case EncryptionSSL:
		conn, err = tls.Dial("tcp", address, config)
	default:
		conn, err = net.Dial("tcp", address)
	}

	if err != nil {
		return nil, errors.New("Mail Error on dailing with encryption type " + encryption.string() + ": " + err.Error())
	}

	c, err := newClient(conn, host)

	if err != nil {
		return nil, errors.New("Mail Error on smtp dial: " + err.Error())
	}

	return c, err
}

// smtpConnect connects to the smtp server and starts TLS and passes auth
// if necessary
func smtpConnect(host string, port string, a auth, encryption encryption, config *tls.Config) (*smtpClient, error) {
	// connect to the mail server
	c, err := dial(host, port, encryption, config)

	if err != nil {
		return nil, err
	}

	// send Hello
	if err = c.hi("localhost"); err != nil {
		c.close()
		return nil, errors.New("Mail Error on Hello: " + err.Error())
	}

	// start TLS if necessary
	if encryption == EncryptionTLS {
		if ok, _ := c.extension("STARTTLS"); ok {
			if config.ServerName == "" {
				config = &tls.Config{ServerName: host}
			}

			if err = c.startTLS(config); err != nil {
				c.close()
				return nil, errors.New("Mail Error on Start TLS: " + err.Error())
			}
		}
	}

	// pass the authentication if necessary
	if a != nil {
		if ok, _ := c.extension("AUTH"); ok {
			if err = c.authenticate(a); err != nil {
				c.close()
				return nil, errors.New("Mail Error on Auth: " + err.Error())
			}
		}
	}

	return c, nil
}

//Connect returns the smtp client
func (server *SMTPServer) Connect() (*SMTPClient, error) {

	var a auth

	switch server.Authentication {
	case AuthPlain:
		if server.Username != "" || server.Password != "" {
			a = plainAuthfn("", server.Username, server.Password, server.Host)
		}
	case AuthLogin:
		if server.Username != "" || server.Password != "" {
			a = loginAuthfn("", server.Username, server.Password, server.Host)
		}
	case AuthCRAMMD5:
		if server.Username != "" || server.Password != "" {
			a = cramMD5Authfn(server.Username, server.Password)
		}
	}

	var smtpConnectChannel chan error
	var c *smtpClient
	var err error

	// if there is a ConnectTimeout, setup the channel and do the connect under a goroutine
	if server.ConnectTimeout != 0 {
		smtpConnectChannel = make(chan error, 2)
		go func() {
			c, err = smtpConnect(server.Host, fmt.Sprintf("%d", server.Port), a, server.Encryption, new(tls.Config))
			// send the result
			smtpConnectChannel <- err
		}()
		// get the connect result or timeout result, which ever happens first
		select {
		case err = <-smtpConnectChannel:
			if err != nil {
				return nil, err
			}
		case <-time.After(server.ConnectTimeout):
			return nil, errors.New("Mail Error: SMTP Connection timed out")
		}
	} else {
		// no ConnectTimeout, just fire the connect
		c, err = smtpConnect(server.Host, fmt.Sprintf("%d", server.Port), a, server.Encryption, new(tls.Config))
		if err != nil {
			return nil, err
		}
	}

	return &SMTPClient{
		Client:      c,
		KeepAlive:   server.KeepAlive,
		SendTimeout: server.SendTimeout,
	}, nil
}

// Reset send RSET command to smtp client
func (smtpClient *SMTPClient) Reset() error {
	return smtpClient.Client.reset()
}

// Noop send NOOP command to smtp client
func (smtpClient *SMTPClient) Noop() error {
	return smtpClient.Client.noop()
}

// Quit send QUIT command to smtp client
func (smtpClient *SMTPClient) Quit() error {
	return smtpClient.Client.quit()
}

// Close closes the connection
func (smtpClient *SMTPClient) Close() error {
	return smtpClient.Client.close()
}

// send does the low level sending of the email
func send(from string, to []string, msg string, client *SMTPClient) error {
	//Check if client struct is not nil
	if client != nil {

		//Check if client is not nil
		if client.Client != nil {
			var smtpSendChannel chan error

			// if there is a SendTimeout, setup the channel and do the send under a goroutine
			if client.SendTimeout != 0 {
				smtpSendChannel = make(chan error, 1)

				go func(from string, to []string, msg string, c *smtpClient) {
					smtpSendChannel <- sendMailProcess(from, to, msg, c)
				}(from, to, msg, client.Client)
			}

			if client.SendTimeout == 0 {
				// no SendTimeout, just fire the sendMailProcess
				return sendMailProcess(from, to, msg, client.Client)
			}

			// get the send result or timeout result, which ever happens first
			select {
			case sendError := <-smtpSendChannel:
				checkKeepAlive(client)
				return sendError
			case <-time.After(client.SendTimeout):
				checkKeepAlive(client)
				return errors.New("Mail Error: SMTP Send timed out")
			}

		}
	}

	return errors.New("Mail Error: No SMTP Client Provided")
}

func sendMailProcess(from string, to []string, msg string, c *smtpClient) error {
	// Set the sender
	if err := c.mail(from); err != nil {
		return err
	}

	// Set the recipients
	for _, address := range to {
		if err := c.rcpt(address); err != nil {
			return err
		}
	}

	// Send the data command
	w, err := c.data()
	if err != nil {
		return err
	}

	// write the message
	_, err = fmt.Fprint(w, msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}

//check if keepAlive for close or reset
func checkKeepAlive(client *SMTPClient) {
	if client.KeepAlive {
		client.Client.reset()
	} else {
		client.Client.quit()
		client.Client.close()
	}
}
