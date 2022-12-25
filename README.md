otgkey
=========

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/akerl/otgkey/build.yml?branch=main)](https://github.com/akerl/otgkey/actions)
[![GitHub release](https://img.shields.io/github/release/akerl/otgkey.svg)](https://github.com/akerl/otgkey/releases)
[![License](https://img.shields.io/github/license/akerl/otgkey)](https://github.com/akerl/otgkey/blob/master/LICENSE)

otgkey sends keycodes to /dev/hidg0 to emulate a keyboard on devices that support the USB Gadget protocol.

## Usage

Send codes as sequences of "modifier:key". Multiple modifiers or keys can be pressed at once using comma-delimited lists, like "shift,meta:i"

For example, to type "Hello":

```
otgkey send shift:h e l l o
```

## Installation

```
go install github.com/akerl/otgkey@latest
```

## License

otgkey is released under the MIT License. See the bundled LICENSE file for details.
