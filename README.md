<img src="assets/logo.png" alt="logo" width="300" height="300" />

# Gorya

Schedule for EC2, RDS and EKS instances. A Golang port of [Doiintl's Zorya](https://github.com/doitintl/zorya).

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/nduyphuong/gorya/dev/LICENSE)
[![Build status](https://github.com/nduyphuong/gorya/actions/workflows/release.yml/badge.svg)](https://github.com/nduyphuong/gorya/actions)

## Setup

By default, in-mem sqlite is used but MySQL is recommended for production setup.

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
