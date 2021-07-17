# IP Checker
This tool can be used to check for changes in your server's public IP address. It can be useful if your server doesn't have a static public IP assigned by your ISP or if it's IP changes frequently.

The script checks the currently known IP address of the server and compares it to an actual IP address fetched from the Internet. If the address changed, it alerts you by email (using [SendGrid](https://sendgrid.com/)). It is best to run this script as a CRON job.

## Prerequisities
- A verified [SendGrid](https://sendgrid.com/) account

## Installation
To manually clone and run the script:

```bash
git clone https://github.com/Loupeznik/ip-checker.git
cd ip-checker
cp .env.example .env
go get
go run main.go
```

## License
This project is [MIT](https://github.com/Loupeznik/ip-checker/blob/master/LICENSE) licensed.
