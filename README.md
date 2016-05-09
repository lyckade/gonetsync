# gonetsync
A simple sync solution in go


## Concept

There is a client and a server communicating together. the gonetsync instance works as client and server. When started it listens

## API

Just a short description how this very simple API works between the client and the server.

* Client sends a GET to the server with the filepath
  * Path: /server/file/
* Server answers if file exists. If the file already exists on the server the fileInfo is sent, too.
  * Path: /client/file/
* Client decides to send or not to send the file via PUT
  * Path /client/file/
