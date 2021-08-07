# IP Checker
This tool can be used to check for changes in your server's public IP address. It can be useful if your server doesn't have a static public IP assigned by your ISP or if it's IP changes frequently.

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
[![License](https://img.shields.io/github/license/Loupeznik/ServerStatusChecker?style=for-the-badge)](./LICENSE)

The script checks the currently known IP address of the server and compares it to an actual IP address fetched from the Internet. If the address changed, it alerts you by email (using [SendGrid](https://sendgrid.com/)) or via Slack. It is best to run this script as a CRON job.

## Prerequisities
For sending alerts via Email
- A verified [SendGrid](https://sendgrid.com/) account

For sending alerts via Slack
- A Slack workspace with an authorized Slackbot
- A room to post alerts to within the Slack workspace

## Installation
To manually clone and run the script:

```bash
git clone https://github.com/Loupeznik/ip-checker.git
cd ip-checker
cp .env.example .env
go get
```

Or you can download a pre-built binary from the [Releases](https://github.com/Loupeznik/ip-checker/releases) tab.

## Usage
```bash
go run . --email
go run . --slack
```

or

```bash
ip-checker --email
ip-checker --slack
```

## License
This project is [MIT](https://github.com/Loupeznik/ip-checker/blob/master/LICENSE) licensed.
