## Structure
```
.
├── README.md
├── handlers.go
├── internal
│   ├── messages
│   │   ├── mailgun.go
│   │   ├── mandrill.go
│   │   ├── messages.go
│   │   └── messages_test.go
│   └── testclient
│       └── testclient.go
├── main.go
├── models
│   └── models.go
└── settings.json
```

## Language Chosen
I chose to write this in Go because of it's readbility and the fact that it's what I've been using most often the last couple of years.

## Tradeoffs
There are many things I would do differently given more time. First, I didn't really have time to write good tests. I would add in unit tests for each function. I would take advantage of the Message interface to mock out responses from the email providers. I included one very basic example of how I might structure a test, but it's not complete in that it doesn't test all functionality.

I would also:
* add auth to the endpoint, as well as encrypt the api_key vs having it in plain text in the settings file.

* use the from and to names

* add more validation to the html being passed in.

* implement a second email provider (it looks like Mandrill is no longer free; Mailchip would not allow me to create an api key with my free account).

* genericize the NewRequest function.

## How to Install & Run
 1. Install go (this may vary by package manager)
 ex. `brew install go`

2. Create go workspace
* `mkdir ~/go`
* `cd ~/go && mkdir src`
* `cd ~/go/src && mkdir github.com`

3. Clone into go path
`cd ~/go/src/github.com && git clone git@github.com:hlhendy/emails.git`

4. Add API Key to settings.json file

5. Start server
`cd ~/go/src/github.com/email && go build && ./emails`

6. Curl
```
curl -XPOST localhost:8080/email -d '{"to":"<to>", "from":"<from>", "to_name":"<to-name>", "from_name":"<from-name>", "subject":"It's a test subject!", "body":"<body>"}' -H "Content-Type: application/json"
```

7. Run Tests
`cd ~/go/emails && go test ./...`
