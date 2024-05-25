# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
All dates in this file are given in the [UTC time zone](https://en.wikipedia.org/wiki/Coordinated_Universal_Time).

## Unreleased

## 0.2.2 - 2024-05-25

### Removed

- The links to Portainer have been removed, as part of the deprecation of Portainer in the default configuration of v2024.0.0 of the PlanktoScope OS.

### Fixed

- The links to the offline copies of the operation protocols are no longer broken.

## 0.2.1 - 2024-04-23

### Fixed

- The stylesheets were not generated correctly in 0.2.0; they should now be generated correctly again.

## 0.2.0 - 2024-04-22

### Changed

- (Breaking change) Instead of generating a machine name from a serial number which is either specified as the `MACHINENAME_SN` environment variable or loaded from a file specified by the `MACHINENAME_SNFILE` environment variable, now the device portal just tries to load the machine name from the `MACHINENAME_NAME` environment variable or else from a file specified by the `MACHINENAME_NAMEFILE` environment variable (which defaults to `/run/machine-name`), and it falls back to a name of "unknown" if no machine name is found.

## 0.1.15 - 2024-01-11

### Fixed

- Fixed malformed HTML structure on the landing page.

## 0.1.14 - 2024-01-11

### Added

- Added a link to the Grafana dashboard.
- Added an API entry for the node-exporter host metrics.
- Added an infrastructure entry for the Prometheus server.

## 0.1.13 - 2024-01-10

### Added

- Added a link to the Dozzle log viewer.

## 0.1.12 - 2023-12-01

### Changed

- All links now open in new tabs (see [PlanktoScope#231](https://github.com/PlanktoScope/PlanktoScope/pull/231)).
- The landing page now shows a deprecation notice to the user if it is accessed using the `planktoscope.local` hostname.
- The URL for the online PlanktoScope project documentation has been changed from the old site (<https://planktoscope.readthedocs.io>) to the new site (<https://docs.planktoscope.community>).

## 0.1.11 - 2023-09-06

### Changed

- Shortened machine names and updated information about machine-specific domain names and URLs to use the new `pkscope` abbreviation instead of `planktoscope` (see [PlanktoScope#214](https://github.com/PlanktoScope/PlanktoScope/pull/214)).

### Fixed

- Fixed mistakes in machine-specific URLs.

## 0.1.10 - 2023-08-31

### Added

- Added more information to the "Wrong PlanktoScope?" section to help users troubleshoot situations where they're connected to the wrong PlanktoScope.

## 0.1.9 - 2023-08-31

### Fixed

- Fixed the ability to scroll the landing page with arrow keys upon initial load (rather than having to click on the landing page first).

## 0.1.8 - 2023-08-29

### Added

- Added link to a file browser for the device-backend's logs

## 0.1.7 - 2023-08-22

### Added

- Added links to offline copies of the PlanktoScope quantitative protocols from protocols.io.

## 0.1.6 - 2023-08-10

### Changed

- Moved links for Cockpit, the system file manager, Portainer, and the Node-RED dashboard editor into the "For advanced users" section of the home page.

## 0.1.5 - 2023-05-26

### Added

- Added a link to the offline PlanktoScope documentation, served from the PlanktoScope itself.

## 0.1.4 - 2023-05-24

### Fixed

- Fixed a few incorrect URLs to other services from the landing page.

## 0.1.3 - 2023-05-24

### Fixed

- When trying to determine the machine's serial number from a file (e.g. from the Raspberry Pi's firmware device tree), only the 32 least-significant bits (i.e. the 8 rightmost hex characters) are used. This is needed because the Raspberry Pi 4's serial number is prefixed with `10000000`.

## 0.1.2 - 2023-05-23

### Fixed

- When trying to determine the machine's serial number from a file (e.g. from the Raspberry Pi's firmware device tree), any trailing null terminators in the file are ignored.

## 0.1.1 - 2023-05-23

### Changed

- The machine name is now lazily loaded upon the first time the landing page is loaded and, unless an error occurs (e.g. from trying to determine the machine's serial number), cached for future use.
- (Breaking change) The optional `SERIAL_NUMBER` environment variable has been renamed to `MACHINENAME_SN`, and the optional `SERIAL_NUMBER_FILE` environment variable has been renamed to `MACHINENAME_SNFILE`.

## 0.1.0 - 2023-05-20

### Added

- A minimalist static web page with a list of useful links.
