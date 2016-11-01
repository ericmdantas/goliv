[![Build Status](https://travis-ci.org/ericmdantas/goliv.svg?branch=master)](https://travis-ci.org/ericmdantas/goliv)

## Disclaimer

This is a work in progress.

#### Todo:

As of right now, there's no implementation of either `Reverse-Proxy` or `HTTP/2`.

The first is all about either using the `httputil` from the `stdlib` or using a 3rd party lib. 
The later, even though is easier to implement, I still have to find a way to use the certs in the user machine. 
Right now, when Go installs `goliv` it won't bring the certs left and right, which makes the `-secure` option fail - 
because it needs a physical path to those files, both the `server.crt` and the `server.key`.


#### How to help

If you want to help the development, you can either clone or fork the repo and run:

```shell
$ go run main.go -root=_fixture
```

Navigate to `http://127.0.0.1:1309` (you can open in more than one browser), it'll show the html, 
which is being served from the `_fixture` folder. Everytime you edit the `index.html`, `app.js` or `style.css`, 
the browsers will be refreshed - showing the new content.

Please, check if things are working correctly and try the other params to see if they're working. Report any bugs you find.



## What?

Light, fast and powerful one liner live-reloading Golang server.

From the simplest live-reloading server to complex apps that need `compression`, `proxies` and `https`/`http2` - `goliv` got you covered.


## Install

```shell
$ go get -u github.com/ericmdantas/goliv
```


## Why?

`goliv` simplifies a lot of headache we have when developing complex web apps. 

- Refresh all your browsers with each file change;
- Proxy request/responses;
- Automagically gzip the response of your server so your page loads faster;
- Use HTTP/2 by simply setting `secure` to `true`;
- Use less memory/CPU possible.


## How?

You can choose the way to work with go: `CLI` (terminal), `.golivrc` (config file) or a `local golang package`.

Go to the folder that contains the `index.html` file and run:

```shell
$ goliv
```

There you go, all running!

Oh, do you want some specific stuff? Checkout the available <a href="#config">config</a>.


## Config

#### CLI


```
-port                      change port
-host                      the host name, instead of showing localhost/127.0.0.1
-secure                    use https/wss
-quiet                     no logging whatsoever
-noBrowser                 won't open the browser automagically
-only                      will only watch for changes in the given - slice
-ignore                    won't watch for changes in the given paths - slice
-pathIndex                 change the path to your index.html
-proxy                     uses proxy
-proxyTarget               the http/https server where the proxy will "redirect"
-proxyWhen                 when the proxy should be activated; like --pxw /api/*
-root                      set the root to a different folder, like "./src/my/deep/folder/"
-watch                     choose to watch for files change or not
-static                    choose what paths are going to be served
```


#### .golivrc

All the <a href="#config">config</a> being used on the `CLI` can be added to the `.golivrc` file, like this:

```json
{
  "port": ":9999",
  "quiet": true,
  "pathIndex": "src/",
  "only": [
      "src"
  ],
  "proxy": true,
  "proxyTarget": "http://my-other-server.com:1234",
  "proxyWhen": "/api/*"
}
```

By doing that, when running `$ goliv`, it'll get all that's inside the config in `.golivrc` and use it.

But keep in mind that if you have such file and still use something like `$ goliv -port :1234`, **the cli will have priority** over the file.


#### Using in a Go package

```go
import (
    "github.com/ericmdantas/goliv"
)

func main() {
     cfg := goliv.NewConfig()

     cfg.Quiet = true
     cfg.Root = "client/dev"

    // yes, that easy - now your browser will open 
    // and it'll be refreshed every time a file changes
     if err := goliv.Start(cfg); err != nil {
         panic(err)
     }
}
```

#### Default values

```
-port          is :1308
-host          is 127.0.0.1
-secure        is false
-quiet         is false
-only          is []string{"."}, which means it'll watch everything
-ignore        is []string{}, which means it won't ignore anything
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
