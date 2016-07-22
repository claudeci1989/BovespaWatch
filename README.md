# Bovespa Watch

This is a simple Go application that will watch changes in the Brazilian stock market - [Bovespa] - in order to notify the user (via the SMS protocol, with the help of [Twilio]) when pre-defined events happen (if they happen).

## Libraries

This project uses several open source softwares to work properly, they are:

* [subosito/twilio]: https://github.com/subosito/twilio
* [lib/pq]: https://github.com/lib/pq

## Usage

```go
$ go run BovespaWatch.go
```

## Copyright

© 2016 Rodrigo Alves Vieira. All Rights Reserved.

[Bovespa]: http://www.bmfbovespa.com.br/en_us/
[Twilio]: https://www.twilio.com/
