# Net-Cat
![Net-Cat Banner](https://i.pinimg.com/736x/93/33/75/933375235e70486148d5880d2a2bd617.jpg)
Net-Cat is a TCP-based chat application written in Go that replicates the behavior of the traditional `netcat` utility in a Server-Client architecture. It supports multi-client group chats with robust features to make real-time communication seamless and engaging.

---

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

### Usage
 ## Steps to Install
- Install Go (https://golang.org/doc/install).
### Clone the Repository
Open your terminal and run the following commands:

for github:
```bash
git clone https://github.com/nyagooh/net-cat.git
```
for gitea:
```bash
git clone https://learn.zone01kisumu.ke/git/nymaina/net-cat
cd net-cat
```
## Run the Server

Start the server on the default port `8989`:

```bash
go run .
```
You can start the server on a custom port (e.g., 2525):
```bash
go run . 2525
```
The server will listen for incoming connections on the specified port.

## Connecting as a Client
You can use any TCP client to connect to the server (e.g., nc or this application). For example:
 - This will connect when using a different laptop
```bash
nc <server_ip> <port>
```
- Example to connect locally:
  ```bash
  nc localhost 2525
  ```

## Example Interaction
  #### server
  ```bash
  Listening on the port :2525
```
#### client interaction
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
[ENTER YOUR NAME]: Alice
```
#### Chat Messages
```bash
[2024-12-09 14:30:01][Alice]: Hello, everyone!
[2024-12-09 14:30:10][Bob]: Hi Alice!
```
#### Leave Notifications:
```bash
Bob has left the chat.
```
### Contributions

Contributions are welcome! If you’d like to improve the project or add new features:

- Fork this repository.
- Create a new branch: git checkout -b feature-branch-name.
- Commit your changes: git commit -m "Add a new feature".
- Push the branch: git push origin feature-branch-name.
- Submit a pull request.

### Authors

This project was developed by:
 - nyagooh
 - okwach
 - Anxielray
- 
### [License](/home/nymaina/Documents/net-cat/LICENSE)




