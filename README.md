# device-portal
A browser landing page for user access to the PlanktoScope

## Introduction

Users connect to the PlanktoScope from the web browser by opening a URL like http://home.planktoscope . This tool provides a web page with a list of links to the various network services running on the PlanktoScope (e.g. the Node-RED frontend dashboard, the Cockpit system administration panel, or the system file manager) for easy access. It is meant to be served from a reverse-proxy on port 80 along with all other network services, configured as in https://github.com/PlanktoScope/pallets - otherwise, the links will be incorrect.

In the future, this tool will hopefully give the user (or otherwise direct the user to) a setup wizard for configuring localization settings (e.g. for languages and wifi networks) upon the first boot of the PlanktoScope.

## Usage

### Deployment

First, you will need to download device-portal, which is available as a single self-contained executable file. You should visit this repository's [releases page](https://github.com/PlanktoScope/device-portal/releases/latest) and download an archive file for your platform and CPU architecture; for example, on a Raspberry Pi 4, you should download the archive named `device-portal_{version number}_linux_arm.tar.gz` (where the version number should be substituted). You can extract the device-portal binary from the archive using a command like:
```
tar -xzf device-portal_{version number}_{os}_{cpu architecture}.tar.gz device-portal
```

Then you may need to move the device-portal binary into a directory in your system path, or you can just run the device-portal binary in your current directory (in which case you should replace `device-portal` with `./device-portal` in the commands listed below).

Once you have device-portal, you can run it as follows on a Raspberry Pi:
```
./device-portal
```

Then you can view the landing page at http://localhost:3000 . Note that if you are running it on a computer other than the Raspberry Pi with the standard PlanktoScope OS, then you will need to set some environment variables (see below) to non-default values.

### Development

To install various backend development tools, run `make install`. You will need to have installed Go first.

Before you start the server for the first time, you'll need to generate the webapp build artifacts by running `make buildweb` (which requires you to have first installed [Node.js v16](https://nodejs.org/en/) (lts/gallium) and [Yarn Classic](https://classic.yarnpkg.com/lang/en/)). Then you can start the server by running `make run` with the appropriate environment variables (see below). You will need to have installed golang first. Any time you modify the webapp files (in the web/app directory), you'll need to run `make buildweb` again to rebuild the bundled CSS and JS. Whenever you use a CSS selector in a template file (in the web/templates directory), you should *also* run `make buildweb`, because the build process for the bundled CSS omits any selectors not used by the templates.

### Building

Because the build pipeline builds Docker images, you will need to either have Docker Desktop or (on Ubuntu) to have installed QEMU (either with qemu-user-static from apt or by running [tonistiigi/binfmt](https://hub.docker.com/r/tonistiigi/binfmt)). You will need a version of Docker with buildx support.

To execute the full build pipeline, run `make`; to build the docker images, run `make build` (make sure you've already run `make install`). Note that `make build` will also automatically regenerate the webapp build artifacts, which means you also need to have first installed Node.js as described in the "Development" section. The resulting built binaries can be found in directories within the dist directory corresponding to OS and CPU architecture (e.g. `./dist/device-portal_window_amd64/device-portal.exe` or `./dist/device-portal_linux_amd64/device-portal`)

### Environment Variables

If you are running device-portal on a computer which is not a Raspberry Pi with the standard PlanktoScope OS, then you'll need to set some environment variables. Specifically, you'll need to set:

- Either `MACHINENAME_NAME`, which should be a string representing the name of the machine to be displayed on the landing page, or `MACHINENAME_NAMEFILE`, which should be the path to a file containing the name of the machine to be displayed on the landing page.

For example, you could run device-portal with the machine name `metal-slope-23501` with one of the following commands:
```
# If you downloaded a device-portal binary:
MACHINENAME_NAME=metal-slope-23501 ./device-portal
# If you are developing the project:
MACHINENAME_NAME=metal-slope-23501 make run
```

## Licensing

Except where otherwise indicated, source code provided here is covered by the following information:

Copyright Ethan Li and PlanktoScope project contributors

SPDX-License-Identifier: Apache-2.0 OR BlueOak-1.0.0

You can use the source code provided here either under the [Apache 2.0 License](https://www.apache.org/licenses/LICENSE-2.0) or under the [Blue Oak Model License 1.0.0](https://blueoakcouncil.org/license/1.0.0); you get to decide. We are making the software available under the Apache license because it's [OSI-approved](https://writing.kemitchell.com/2019/05/05/Rely-on-OSI.html), but we like the Blue Oak Model License more because it's easier to read and understand.
