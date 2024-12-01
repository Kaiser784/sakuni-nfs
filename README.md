# Sakuni NFS

Encrypted File System
## Overview

Sakuni NFS is an encrypted file system implemented in Go. It provides secure storage and access to files with encryption to protect sensitive data.
## Features
### Current Features

**File and Directory Management:**
- List files and directories (ls)
- Change directory (cd)
- Create directories (mkdir)
- Create files (touch)
- Print working directory (pwd)
- Open files (open [filename])
- Display file contents (cat)
- Clear screen (clear)
- Exit the filesystem (exit)

**Shell Interface:**
- Interactive shell with command execution
- User authentication and session management
- Command auto-completion

**Filesystem Operations:**
- Save the current state of the filesystem
- Reset filesystem

### Upcoming Features

Based on the current issues and pull requests, the following features are planned for future releases:=
- Enhanced encryption mechanisms for added security.
- Improved user management with access control lists.
- Support for decentralized file storage.
- Implementation of more file manipulation commands.
- Integration with cloud storage solutions.
- Performance optimizations and bug fixes.

## Installation

To install Sakuni NFS, ensure you have Go installed on your system and run the following commands:
```bash
git clone https://github.com/Kaiser784/sakuni-nfs.git
cd sakuni-nfs
go build
```
## Usage

Once built, you can use Sakuni NFS with the following command:
```bash
./sakuni-nfs
```
Ensure to follow the usage guidelines and configuration options provided in the documentation.
