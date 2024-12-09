# TCP Client-Server Program in Go

## Overview

- This repository contains a simple TCP client-server application written in Go. The server accepts incoming connections from clients, reads messages sent by them, and sends responses back.
- This program serves as a basic demonstration of network programming in Go using TCP.

## Features

- TCP server that listens for incoming client connections.
- Handling of multiple client sessions concurrently.
- The server can read messages sent by clients and respond accordingly.
- Graceful handling of network errors and connection closures.

## Getting Started

To run the program, ensure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

### Prerequisites

- Go 1.14 or later
- Basic understanding of Go programming and TCP networking concepts

### Installation

1. Clone the repository to your local machine:

```sh
   git clone https://learn.zone01kisumu.ke/git/nymaina/net-cat.git
   cd net-cat
```

# Navigate to the directory containing the server code.

## Running the Server

1. Open your terminal and navigate to the server code directory.

2. Compile and run the server:ou could setup the desired server port number in linux-based environments by

```sh
export SERVER_PORT=<port_number>
```
- However the default port number is set to 8080

- The server will start listening on the specified port/8080(if none is specified)
- You will see output indicating that the server is running and waiting for connections.

## Connecting a Client
========????????????????

## Error Handling

The server is designed to handle errors gracefully. If an error occurs while accepting a connection or reading from the connection, it will print an error message and continue running instead of crashing.

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, feel free to create an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 