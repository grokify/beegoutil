# Beego OAuth 2.0 Demo

[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

OAuth 2.0 demo app using the following:

* Core
  * [Beego](https://beego.me/)
  * [OAuth2](https://github.com/golang/oauth2)
  * [OAuth2 Util](https://github.com/grokify/oauth2-util-go)
* Rendering
  * [Bootstrap](http://getbootstrap.com/)
  * [Font Awesome](http://fontawesome.io/)
  * [Social Buttons](https://lipis.github.io/bootstrap-social/)
  * [Quicktemplate](https://github.com/valyala/quicktemplate)

## Usage

### App Hostname / OAuth 2.0 Redirect URI

Decide on the hostname for your app which will be used in the OAuth 2.0 URI.

The redirect URI is `oauth2callback`. For example, if your app URL is `https://example.com`, your redirect URI will be `https://example.com/oauth2callback`.

If you do not have an Internet accessible hostname, you can use ngrok for testing.

### Configure Google and Facebook

Login to the Google and Facebbook developer consoles to configure your apps with your OAuth 2.0 rediect URIs. During this process, you will get your OAuth 2.0 client ID and client secrets for Google and Facebook.

### Install & Configure App

1. Clone the repo
2. Install the dependencies
3. Set up your `app.conf` file.
4. Run the app

You can do the above with the following steps:

```bash
$ git clone https://github.com/grokify/beego-oauth2-demo
$ go get ./...
$ cd beego-oauth2-demo/conf
$ cp app.conf.sample app.conf
$ vim app.conf
$ bee run
```

 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/beego-oauth2-demo
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/beego-oauth2-demo
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/beego-oauth2-demo
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/beego-oauth2-demo/blob/master/LICENSE.md
