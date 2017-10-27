# burl

A broken-URL checker.

Work in progress. Not even alpha quality.

## Install

```
▶ go get github.com/tomnomnom/burl
```

## Usage

Given some URLs in a file:

```
▶ cat urls
http://example.com/
http://example.com/notafile.js
http://pleasedontregisterthisdomain.com/js/main.js
An invalid URL
https://wat/foo.js
https://example.net
https://notarealsubdomain.example.com/
```

Either feed them into `burl` on stdin:

```
▶ cat urls | burl
non-200 response code: http://example.com/notafile.js (404 Not Found)
does not resolve: http://pleasedontregisterthisdomain.com/js/main.js
invalid url: An invalid URL
does not resolve: https://wat/foo.js
does not resolve: https://notarealsubdomain.example.com/
```

Or pass the filename as the first argument:

```
▶ burl urls
non-200 response code: http://example.com/notafile.js (404 Not Found)
does not resolve: http://pleasedontregisterthisdomain.com/js/main.js
invalid url: An invalid URL
does not resolve: https://wat/foo.js
does not resolve: https://notarealsubdomain.example.com/
```
