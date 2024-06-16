<h1 align="center">
  <br>
  <a href="http://www.amitmerchant.com/electron-markdownify"><img src="https://img001.prntscr.com/file/img001/GNcP_Rd9S8We8vyzdKWmow.png" alt="Markdownify" width="200"></a>
  <br>
  Cyclists
  <br>
</h1>

<h4 align="center">RESTful API made with <a href="https://go.dev/" target="_blank">Golang</a>.</h4>

<p align="center">
  <a>
    <img src="https://img.shields.io/github/go-mod/go-version/matheusgb/cyclists" alt="go version">
  </a>
</p>

<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#documentation">Documentation</a>
</p>

## Key Features

* Login with JWT
* Users CRUD
* Bike Events CRUD
* Possibility for users to register for bike events
* SendGrid to send password redefinition emails
* Paginations
* Admin privileges

## How To Use

Fill in `config.json` file with your PostgreSQL database, JWT, and SendGrid credentials.

There is no necessity to change `enviroment` key.

If you have [air](https://github.com/air-verse/air) installed, you can run the project with the command in root folder:

```
air
```

Also, you can run with:

```
go run main.go
```

You can run all tests with:

```
go test -v  ./src/tests/... 
```

## Documentation

You can find the documentation here:
[Postman](https://documenter.getpostman.com/view/23223146/2sA3XQiNMY)
