# MacLeod

## SSL terminator

#### For when **there can be only one** entry point

## Features

- Terminates SSL connections and redirects traffic internally
- No external dependencies

## Download precompiled

Download for your OS/Arch from https://github.com/reneManqueros/macleod/releases

## Compile from source:

```sh
go build .
```

## Configuration

Copy **config.sample.json** to **config.json**,
Create a node under backends for each domain that will be used

```json
 {
  "web.domain.com": {
    "destination": "10.42.0.10:8080",
    "certificate": "/etc/letsencrypt/live/web.domain.com/fullchain.pem",
    "key": "/etc/letsencrypt/live/web.domain.com/privkey.pem"
  }
}
```

| Setting | Description                                        |
| ------ |----------------------------------------------------|
| web.domain.com | Set this to your domain name                       |
| destination | IP or hostname that should receive the traffic     |
| certificate | Location of chain file                             |
| key | Location of private key file                       |

You can also use **macleod.service** as reference to install as daemon