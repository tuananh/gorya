<img src="assets/logo.png" alt="logo" width="300" height="300" />

# Gorya

Schedule for EC2, RDS and EKS instances. A Golang port of [Doiintl's Zorya](https://github.com/doitintl/zorya).

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/nduyphuong/gorya/dev/LICENSE)
[![Build status](https://github.com/nduyphuong/gorya/actions/workflows/release.yml/badge.svg)](https://github.com/nduyphuong/gorya/actions)

## Building Gorya

### Software requirements

-   [go 1.20+]
-   [git]

## Setup your environments

By default, in-mem sqlite is used but MySQL is recommended for production setup.

#### Option 1: Set up with docker-compose
1. Create a new directory for project if not exists.
```bash
mkdir -p ~/go/src/github.com/nduyphuong/gorya
```
2. Clone the source code
```bash
cd ~/go/src/github.com/nduyphuong/gorya
git clone https://github.com/nduyphuong/gorya
```
3. Set up the stack with docker
```bash
cd ~/go/src/github.com/nduyphuong/gorya
docker-compose up -d
```
#### Option 2: Set up with helm

TBD

## How it works

```mermaid
sequenceDiagram
autonumber
actor U as User
participant G as Gorya

participant P as Gorya Processor
participant D as Gorya Dispatcher

participant Q as GoryaQueue
participant C as Cloud Provider APIs

loop Every 60 Minutes
U->>G: Create off time schedule
D->>G: Evaluate
end
D->>Q: Dispatch task
Q->>P: Process next item
P->>C: Change resource status

```

[go 1.20+]: https://go.dev/doc/install
[git]: https://docs.github.com/en/get-started/quickstart/set-up-git
