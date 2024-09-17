## Runtipi CLI Go

A reimplementation of the [Runtipi CLI](https://github.com/runtipi/cli) in the Go programming language
with some imporevements and new features.

<img src="screenshots/screenshot.png" width="457" height="296" />

> Warning ⚠️: This project is built and maintained by volunteers. There is no guarantee of support or security when you use it. While the system is considered stable, it is still in active development and may contain bugs.

> Info ℹ️: Runtipi CLI Go is written in 100% Go, if you would like to help check out the issues section or our Discord server!
> 

### Why?

We are just having fun with [Nicolas](https://github.com/meienberger) learning new languages and we thing Go is a better option for our CLI.

## Features

The Runtipi Go CLI has all the official Runtipi CLI features and the following ones:

- List app backups command (`./runtipi-cli app list-backups`)
- Healthcheck command using the worker API
- System readings from the worker API
- Backup command
- List backups command (for runtipi)
- Neofetch command (easter egg)
- Automatic backup of the CLI before update
- Increased timeout on app commands to 5 and 15 minutes (for updates) to make sure it is the same as the dashboard timeout

### Installation

To install the CLI you need to follow 4 simple steps.

1. Download the latest version matching your system's arch from the [releases](https://github.com/runtipi/runtipi-cli-go/releases/) page.
2. Put the CLI in your `runtipi` folder and rename it to `runtipi-cli-go`.
3. Make it executable `chmod +x runtipi-cli-go`
4. Start using it `sudo ./runtipi-cli-go start`

### Building

To build the CLI you need to have go installed and then run these commands:

Firstly clone the repository:

```bash
git clone https://github.com/runtipi/runtipi-cli-go
cd runtipi-cli-go/
```

Install packages:

```bash
go get .
```

Build:

```bash
go build
```

> Note 🗒️: You can get the CLI down to around 8mb using the `go build -ldflags '-w -s'` command.

You should have the CLI named `runtipi-cli-go`.

### License

The license is the same as the official Runtipi CLI, so the project is licensed under the GNU General Public License v3.0. TL;DR — You may copy, distribute and modify the software as long as you track changes/dates in source files. Any modifications to or software including (via compiler) GPL-licensed code must also be made available under the GPL along with build & install instructions.

### Contributing

If you like you can contribute to this project by creating a pull request. Any help is appreciated.

### Acknowledgements

Thank's a lot to:

- [Carbon](https://carbon.sh) for the cool CLI screenshot