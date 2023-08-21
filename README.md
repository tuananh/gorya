<img src="assets/logo.png" alt="logo" width="500" height="400" />

# Gorya

Schedule for EC2, RDS and EKS instances

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/nduyphuong/reverse-registry/dev/LICENSE)
[![Build status](https://github.com/nduyphuong/poorman-registry/actions/workflows/release.yml/badge.svg)](https://github.com/nduyphuong/reverse-registry/actions)


[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run/?git_repo=https://github.com/nduyphuong/poorman-registry.git)

## Setup

By default, in-mem sqlite is used but MySQL is recommended for production setup.

## How it works

```mermaid
sequenceDiagram
autonumber
actor U as User
participant RR as Reverse Registry
participant DB as Local Digest Database
participant CG as Chainguard Images
U->>+RR: Pull command `nginx:1.0.0`
RR->>+DB: Check if digest existed for `nginx:1.0.0`
DB-->>-RR: Found digest for `nginx:1.0.0`
RR-->>-U: Return digest if found

loop Every x Minutes
RR->>CG: Periodically checking `latest` tag for digest change
RR->>DB: Save digest for this tag to local db
end


RR->>CG: Proxied every other APIs

```
