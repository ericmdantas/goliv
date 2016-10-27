[![Build Status](https://travis-ci.org/ericmdantas/goliv.svg?branch=master)](https://travis-ci.org/ericmdantas/goliv)

Like [aliv](https://github.com/ericmdantas/aliv), but written in Go.

---

## Disclaimer

This is a work in progress. Most of the features are still missing, so, this is still not complete port of [aliv](https://github.com/ericmdantas/aliv) yet.

If you want to help the development, you can clone the repo and run:

```shell
$ go run main.go -root=_fixture
```

Go to `http://127.0.0.1:1309` (you can open in more than one browser), it'll show a html, which is being served from the `_fixture` folder. Everytime you edit the `index.html`, `app.js` or `style.css`, the browsers will be refreshed - showing the new content.

Please, check if things are working correctly. Also, try the other params and report some bugs. Thanks!


Todo:
- Proxy;
- HTTP/2.


## What?

Light, fast and powerful one liner live-reloading Golang server.

From the simplest live-reloading server to complex apps that need compression, proxies and https - `goliv` got you covered.


## Install

```shell
$ go get -u github.com/ericmdantas/goliv
```


## Why?

`goliv` simplifies a lot of headache we have when developing complex web apps. 

- Refresh all your browsers with each file change;
- Proxy request/responses;
- Automagically gzip the response of your server;
- Use HTTP/2 by simply setting `secure` to `true`;
- Use less memory/CPU possible.


## How?

You can choose the way to work with go: `CLI` (terminal), `.golivrc` (config file) or a `local node module`.

Go to the folder that contains the `index.html` file and run:

```shell
$ goliv
```

There you go, all running!

Oh, do you want some specific stuff? Checkout the available <a href="#options">options</a>.


## Options

#### CLI


```
-port                      change port
-host
-secure                    use https/wss
-quiet                     no logging whatsoever
-noBrowser                 won't open the browser automagically
-only                      will only watch for changes in the given path/glob/regex/array
-ignore                    won't watch for changes in the given path (regex)
-pathIndex                 change the path to your index.html
-proxy                     uses proxy
-proxyTarget               the http/https server where the proxy will "redirect"
-proxyWhen                 when the proxy should be activated; like --pxw /api/*
-root                      set the root to a different folder, like "./src/my/deep/folder/"
-watch                     choose to watch for files change or not
-static                    choose what paths are going to be served
```


#### .golivrc

All the <a href="#options">options</a> being used on the `CLI` can be added to the `.golivrc` file, like this:

```json
{
  "port": 9999,
  "quiet": true,
  "pathIndex": "src/",
  "only": ["src/**/*"],
  "proxy": true,
  "proxyTarget": "http://my-other-server.com:1234",
  "proxyWhen": "/api/*"
}
```

By doing that, when running `$ goliv`, it'll get all the options in `.golivrc` and use it.

But, if you have such file and still use something like `$ goliv -port 9999`, **the cli will have priority** over the file.


#### Go module

```go
import (
    "github.com/ericmdantas/goliv"
)

func main() {
     options := goliv.NewOptions()

    // yes, that easy - now your browser will open 
    // and it'll be refreshed every time a file change
     if err := goliv.Start(options); err != nil {
         panic(err)
     }
}
```

#### Default values

```
-port          is 1307
-host          is 127.0.0.1
-secure        is false
-quiet         is false
-only          is ".", which means it'll watch everything
-ignore        is ^(node_modules|bower_components|jspm_packages|test|typings|coverage|unit_coverage)
-noBrowser     is false, which means it'll always open the browser on start
-pathIndex     is "", which means it'll look for the index.html in the root
-proxy         is false, which means it'll not look for another server to answer for the /api/, for example
-proxyTarget   is "", no server to be target
-proxyWhen     is "", and it's supposed to be set with something like /api/*
-root          is ""
-watch         is true
-static        is [root, root + "/path/to/your/index"]
```


## Wiki

Check the [wiki](https://github.com/ericmdantas/goliv/wiki) for examples, FAQ, troubleshooting and more.

## Contributing

#### I've got an idea!

Great, [let's talk](https://github.com/ericmdantas/goliv/issues/new)!

#### I want to contribute

Awesome!

First, I'd suggest you open an issue so we can talk about the changes to be made and suchs and then you can do whatever you want :smile:

## License

MIT
