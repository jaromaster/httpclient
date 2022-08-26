# HTTPCLIENT

## simple http client written in Go

# Features
- extremely easy to use
- supports: GET, POST, DELETE, PUT
- auth
- cookies
- headers
- json or text data
- persiting response data

# Installation
clone repository<br>
`go build`<br>
`./httpclient help`
<br><br>

# Examples

## GET
### get request
`httpclient get https://someip`
<br><br>

## POST
### post requests (text and json)
`httpclient post https://someip --data='text-data'`<br>
`httpclient post https://someip --data='{"jsonkey": "jsonvalue"}'`
<br><br>

## PUT
### put request (text and json)
`httpclient put https://someip --data='text-data'`<br>
`httpclient put https://someip --data='{"jsonkey": "jsonvalue"}'`
<br><br>

## DELETE
### delete request
`httpclient delete https://someip`
<br><br>

## Save to file
### store response data in file
`httpclient get https://someip --save='./outputfile.txt'`
<br><br>

## Silent mode
### disable output of response data
`httpclient get https://someip --silent`
<br><br>

## Cookies
### set the cookies used in the request
`httpclient get https://someip --cookies='cookie1=1; cookie2=2'`
<br><br>

## Header
### set header (key-value pairs separated by \n)
`httpclient get https://someip --header='Foo: Bar\nOther: Value'`
<br><br>

## Auth
### handle http auth
`httpclient get https://someip --auth='Basic encoded_credentials'`
<br><br>

## User-Agent
### set the user agent used in the request
`httpclient get https://someip --useragent='myclient/1.0'`