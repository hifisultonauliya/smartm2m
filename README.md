<h1 align="center">
<!--   <br>
  <a href="http://www.amitmerchant.com/electron-markdownify"><img src="https://raw.githubusercontent.com/amitmerchant1990/electron-markdownify/master/app/img/markdownify.png" alt="Markdownify" width="200"></a>
  <br> -->
  SmartM2M
  <br>
</h1>

<h4 align="center">A minimal API for testing SmartM2M use case for itemNFT write using <a href="https://go.dev/" target="_blank">Golang</a>.</h4>

<!--
<p align="center">
  <a href="https://badge.fury.io/js/electron-markdownify">
    <img src="https://badge.fury.io/js/electron-markdownify.svg"
         alt="Gitter">
  </a>
  <a href="https://gitter.im/amitmerchant1990/electron-markdownify"><img src="https://badges.gitter.im/amitmerchant1990/electron-markdownify.svg"></a>
  <a href="https://saythanks.io/to/bullredeyes@gmail.com">
      <img src="https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg">
  </a>
  <a href="https://www.paypal.me/AmitMerchant">
    <img src="https://img.shields.io/badge/$-donate-ff69b4.svg?maxAge=2592000&amp;style=flat">
  </a>
</p>
-->

<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#credits">Credits</a> •
  <a href="#license">License</a>
</p>

<!--
![screenshot](https://raw.githubusercontent.com/amitmerchant1990/electron-markdownify/master/app/img/markdownify.gif)
-->

## Key Features

- Using golang and libraries gin for the api
- Using jwt for simple login and register
  - to accessing itemNFT API must login to get auth token.
- Using mongodb for databases
- Implement some API that required in modules
- Including Testing schenario
- Using docker container

## How To Use

To clone and run this application, you'll need [Git](https://git-scm.com) and [Docker](https://www.docker.com/) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/hifisultonauliya/smartm2m.git

# Go into the repository
$ cd smartm2m

# Run docker container
$ docker-compose up --build

# Checking API using Postman and enjoy :)
```

To run unit test, need to run this following steps.

```bash
# Go into the repository
$ cd smartm2m

# Run go mod tidy to get dependencies
$ go mod tidy

# Run Unit testing for models dependencies
$ go test -v ./src/models
```

<!--
> **Note**
> If you're using Linux Bash for Windows, [see this guide](https://www.howtogeek.com/261575/how-to-run-graphical-linux-desktop-applications-from-windows-10s-bash-shell/) or use `node` from the command prompt.
-->

## Credits

This software uses the following open source packages:

- [Golang](https://go.dev/)
- [Docker](https://www.docker.com/)
- [JWT Libraries](github.com/dgrijalva/jwt-go)
- [Go Gin Framework](github.com/gin-gonic/gin)

## License

MIT

---

> GitHub [@hifisultonauliya](https://github.com/hifisultonauliya)
