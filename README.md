# Net-Cat
![Net-Cat Banner](https://i.pinimg.com/736x/93/33/75/933375235e70486148d5880d2a2bd617.jpg)

Net-Cat is a TCP-based chat application written in Go that replicates the behavior of the traditional `netcat` utility in a Server-Client architecture. It supports multi-client group chats with robust features to make real-time communication seamless and engaging.

---
## Table of Contents
- [Net-Cat](#net-cat)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Learning Objectives](#learning-objectives)
  - [Features](#features)
  - [Installation](#installation)
    - [Clone the Repository](#clone-the-repository)
  - [Usage](#usage)
    - [Run the Server](#run-the-server)
    - [Connecting as a Client](#connecting-as-a-client)
  - [Example](#example)
    - [Starting server](#starting-server)
    - [Connecting as a Client](#connecting-as-a-client-1)
    - [Chat Messages](#chat-messages)
    - [Leave Notifications:](#leave-notifications)
  - [Error Handling](#error-handling)
  - [Contributions](#contributions)
  - [Authors](#authors)
  - [License](#license)

## Overview

This project recreates the `NetCat` command-line tool in a server-client architecture, enabling the following:

- **Server Mode:** Starts a TCP server on a specified port and listens for incoming client connections.
- **Client Mode:** Connects to a specified port, allowing message exchange with the server and other clients in real time.

Inspired by the original `netcat` (`nc`), this project enables learning about networking concepts such as TCP connections, sockets, concurrency, and synchronization. For more information about `NetCat`, refer to the manual using `man nc`.

---
## Learning Objectives
- Understand TCP/UDP protocols and socket programming.
- Implement concurrency in Go using Go routines, channels, and mutexes.
- Learn how to build scalable server-client applications.
- Improve error handling in distributed systems.
  
## Features

1. **Group Chat Functionality**:
   - Supports multiple clients (up to 10 connections).
   - TCP communication with a 1-to-many relation between the server and clients.
   - Real-time message broadcasting between clients.
   - All messages include timestamps and the sender's username:  
     `[YYYY-MM-DD HH:MM:SS][username]:[message]`

2. **Dynamic Updates**:
   - Notifies all clients when a new client joins or leaves the chat.
   - Automatically uploads chat history to newly connected clients.

3. **Validation and Error Handling**:
   - Requires a non-empty username upon joining.
   - Recquires a unique name for each client
   - Handles errors gracefully on both server and client sides.
   - Prevents broadcasting empty messages.

4. **Default Port**: Uses port `8989` if no port is specified, with an appropriate usage message for invalid inputs.

5. **Concurrency**: Implements Go routines, channels, and mutexes to manage multiple connections.

6. **Customizable and Scalable**:
   - Logs chat history to a file for persistence.
   - Supports additional group chats as an optional bonus.

---

 ## Installation
- Install Go
   (https://golang.org/doc/install).
### Clone the Repository
Open your terminal and run the following commands:
```bash
git clone https://learn.zone01kisumu.ke/git/nymaina/net-cat
cd net-cat
```
## Usage
### Run the Server

Start the server on the default port `8989`:

```bash
go run .
```
You can start the server on a custom port (e.g., 2525):
```bash
go run . 2525
```
The server will listen for incoming connections on the specified port.

### Connecting as a Client
You can use any TCP client to connect to the server (e.g., nc or this application). For example:
 - This will connect when using a different laptop
```bash
nc <server_ip> <port>
```
- Example to connect locally:
 ```bash
nc localhost 2525
```

## Example 
  ### Starting server
  ```bash
 go run . 2525
Listening on the port :2525

```
### Connecting as a Client
```bash
nc localhost 2525
```
This is will be the  output
```bash
Welcome to TCP-Chat!
      _nnnn_
     dGGGGMMb
    @p~qp~~qMb
    M|@||@) M|
    @,----.JM|
   JS^\__/  qKL
  dZP        qKRb
 dZP          qKKb
fZP            SMMb
HZM            MMMM
FqM            MMMM
__| ".        |\dS"qML
|    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]: 
```
output after you enter your name
```bash
Welcome to TCP-Chat!
      _nnnn_
     dGGGGMMb
    @p~qp~~qMb
    M|@||@) M|
    @,----.JM|
   JS^\__/  qKL
  dZP        qKRb
 dZP          qKKb
fZP            SMMb
HZM            MMMM
FqM            MMMM
__| ".        |\dS"qML
|    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--' 
[ENTER YOUR NAME]: John
Hello, world!
[2024-12-09 14:40:10][John]: Hello, world!
```
### Chat Messages
```bash
[2024-12-09 14:40:10][John]: Hello, world!
[2024-12-09 14:30:10][Bob]: Hi Alice!
```
### Leave Notifications:
```bash
Bob has left the chat.
```
## Error Handling

 **Empty Username Input**:

The client will not be authenticated
```bash
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]:
Name cannot be empty. Disconnecting...
```

**Duplicate Username**:

The client will received a message to change name
 ```bash
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'
[ENTER YOUR NAME]:ann
Name is already taken. Please choose a different name.
[ENTER YOUR NAME]:
```
 **Maximum Connections**:

- The server supports a maximum of 10 concurrent clients.

 **Invalid Port Input**:
 ```bash
 go run . 2525 localhost
 ```
 the program will throw an error message and exit
 ```bash
[USAGE]: ./TCPChat $port

 ```
## Contributions

Contributions are welcome! If you’d like to improve the project or add new features:
- Fork this repository.
- Create a new branch: git checkout -b feature-branch-name.
- Commit your changes: git commit -m "Add a new feature".
- Push the branch: git push origin feature-branch-name.
- Submit a pull request.

## Authors

This project was developed by:
 - [Nyagooh](https://github.com/nyagooh/)
 - [Hezron](https://github.com/hezronokwach)
 - [Anxielray](https://github.com/anxielray)
## [License](/home/nymaina/Documents/net-cat/LICENSE)




