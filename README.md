#lambhack
A vulnerable serverless lambda application. This is certainly a bad idea to base any coding patterns of what you see here.

lambhack allows you to take advantage of our tried and true application security problems, namely arbitrary code execution, XSS, injection attacks aand more.

This first release only contains arbitrary code execution through the query string.  Please feel free to contribute new vulnerabilities.

## What can you do with lambhack?

See Velocity preso > http://www.slideshare.net/wickett/serverless-security-are-you-ready-for-the-future

## Example CMDEXE

You can pass OS commands in the query string args
```
$ curl “https://XXXX.execute-api.us-east-1.amazonaws.com/prod/lambhack/c?args=uname+-a;+sleep+1"
```

Lambda container reuse in action
```
$ curl “https://XXXX.execute-api.us-east-1.amazonaws.com/prod/lambhack/c?args=ls+/tmp;+sleep+1"

$ curl “https://XXXX.execute-api.us-east-1.amazonaws.com/prod/lambhack/c?args=touch+/tmp/wickettfile;+sleep+1”

$ curl “https://XXXX.execute-api.us-east-1.amazonaws.com/prod/lambhack/args=ls+/tmp;+sleep+1"
```

## Setup

```
go get github.com/wickett/lambhack
```

In case you are new to golang, this clones the prohect to `$GOPATH/src/github.com/wickett/lambhack`

Now you need to setup your AWS user and local credentials.  I recommend setting up creds in `.aws/credentials` and using a profile called sparta with limited perms. 

## License
MIT License

## Contributing
Send in PRs

## Known Problems
I started out calling this thing serverless-audit but have renamed it lambhack. None of the code reflects this yet.
