# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
All dates in this file are given in the [UTC time zone](https://en.wikipedia.org/wiki/Coordinated_Universal_Time).

## Unreleased

### Changed

- All links now open in new tabs (see [PlanktoScope#231](https://github.com/PlanktoScope/PlanktoScope/pull/231)).
- The landing page now shows a deprecation notice to the user if it is accessed using the `planktoscope.local` hostname.

## 0.1.7 - 2023-09-06

### Changed

- Shortened machine names and updated information about machine-specific domain names and URLs to use the new `pkscope` abbreviation instead of `planktoscope` (see [PlanktoScope#214](https://github.com/PlanktoScope/PlanktoScope/pull/214)).

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
