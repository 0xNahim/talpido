## talpido

<h1 align="center">
  talpido
  <br>
</h1>

<p align="center">
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-green.svg"></a>
  <a href="https://twitter.com/0xnahim"><img src="https://img.shields.io/twitter/follow/0xnahim.svg?logo=twitter"></a>
</p>

<p align="center">
  <a href="#whats-it">What's it?</a> •
  <a href="#feautures">Features</a> •
  <a href="#background">Background</a> •
  <a href="#installation-instructions">Installation</a> •
  <a href="#usage">Usage</a> •
  <a href="#to-do">To Do</a>
</p>
<p height="300" align="center">
  <img src="./media/talpidofeliz.jpg">
</p>

# What's it?
Talpido is a tool designed for rapid collection and exfiltration of sensitive information from Linux systems.
> [!IMPORTANT]  
> Crucial information necessary for users to succeed.
> For using Talpido, it is necessary to have the [cloudflared](https://github.com/cloudflare/cloudflared.git) program installed on the attacking machine.

# Feautures
### Features
- Volatile Memory
  - Does not store information on disk; everything is kept in memory, reducing detection and enhancing stealth.

- Automatic implementation
  - Implements a server that automatically receives compressed (ZIP) files and exposes the server via a Cloudflare tunnel, hiding the underlying infrastructure.

- Payload Generation
  - Automatically generates the payload, facilitating quick execution and distribution.

# Background
## What the f* is a talpido?
This tool creates a tunnel and creates a payload that exfiltrates information through a tunnel, so I thought it was appropriate to name it after the mole family.

# Installation Instructions

Talpido requires **go1.18** to install successfully. Run the following command to install.

```
git clone https://github.com/0xNahim/talpido.git && cd ./talpido/cmd/talpido

go build -o talpido main.go
```

> [!NOTE]
> It's not possible to use go install, because I have two main functions. If anyone knows how to fix this, I'd be happy to accept your PR.

# Usage

Using this tool is as simple as typing talpido in the console. The server will automatically start and the payload will be created.
```
./talpido
```

## To Do
- Persistence Module: Implement a persistence mechanism to maintain access to compromised systems even after reboots or other system changes.

- Web Panel for Monitoring "Moles": Develop a web-based control panel to monitor the activities of the "moles" (compromised agents) in real-time.

- Privilege Escalation Implementation: Integrate a module for privilege escalation to increase access levels on compromised systems, allowing deeper exploitation.
