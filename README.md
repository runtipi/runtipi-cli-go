## Runtipi CLI GO

A reimplementation of the [Runtipi CLI](https://github.com/runtipi/cli) in the GO programming language
with some imporevements and new features.

<img src="screenshots/screenshot.png" width="457" height="296" />

> Warning ⚠️: This is in early stages of development, I am trying to make this better than the official CLI and maybe just maybe make Nicolas use this versison

### Why?

Why am I building this? Two reasons, firstly I don't like rust and I think its a bad choice for the CLI so I want to build
something better. Secondly, I just want to learn GO.

### Building

Since this is still in early stages of development the only reason to run this is to build it. To build it you need to have go installed and then run these commands:

Firstly clone the repository:

```bash
git clone https://github.com/steveiliop56/runtipi-cli-go
cd runtipi-cli-go/
```

Install packages:

```bash
go mod tidy
```

Build:

```bash
go build
```

You should have the CLI named `runtipi-cli-go`.

### License

The license is the same as the official Runtipi CLI, so the project is licensed under the GNU General Public License v3.0. TL;DR — You may copy, distribute and modify the software as long as you track changes/dates in source files. Any modifications to or software including (via compiler) GPL-licensed code must also be made available under the GPL along with build & install instructions.

### Contributing

If you like you can contribute to this project by creating a pull request. Any help is appreciated.
