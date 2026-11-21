# Changelog

All notable changes to this project will be documented in this file.

The format is based on Keep a Changelog,
and this project adheres to Semantic Versioning.

---

## [1.9.0] - 2026-11-21

### Changed

- Moved `boolean` from `common/extension/boolean` to `common/boolean`.
- Renamed `ParseFrom()` to `From()` and `BoolToString()` to `ToString()` in `boolean`.
- Updated related package documentation, examples, and tests.
- Added `common/boolean` to the main README.

### Deprecated

- Deprecated `common/extension/boolean`; `ParseFrom()` and `BoolToString()` forward to the new package and will be removed in v2.0.0.

---

## [1.8.0] - 2026-11-21

### Added

- Added the new `testkit` package with dependency-free test helpers for Go 1.18+.
- Added value and error assertions: `Equal()`, `DeepEqual()`, `Match()`, `Pointer()`, `ErrorIs()`, `ErrorAs()`, `MustValue()`, and more.
- Added JSON and HTTP helpers: `JSONEqual()`, `DecodeJSON()`, `Serve()`, `Server()`, `RoundTripFunc`, and `Response()`.
- Added async, file, and time helpers: `Eventually()`, `Receive()`, `Context()`, `WriteFiles()`, `EqualFile()`, and `FakeClock`.
- Added `Passes()`, `FailsSoft()`, `FailsHard()`, and `Record()` for testing custom test helpers.
- Added package documentation, README, examples, and unit tests with 100% coverage.

### Changed

- Updated the main README with the `testkit` package, working package links, and the Go 1.18+ badge.
- Renamed project attribution from Entiqon Project to Entiqon Labs.

---

## [1.7.0] - 2026-11-21

### Added

- Added a release workflow that creates a draft pre-release when a version tag is pushed.
- Added contribution guidelines for roles, reviews, merging, and releases.
- Added issue templates, a pull request template, and `CODEOWNERS`.

### Changed

- Reworked `ptr.EmptyString()`, `ptr.NilIfEmpty()`, and `ptr.NilIfBlank()` to build on `ptr.Of()`.
- Updated `ptr.Of()` to avoid the Go 1.26 `newexpr` hint while staying compatible with Go 1.18+.
- Updated `ptr` tests to run without warnings on every supported Go version.
- Updated the CI workflow to test Go 1.18, `oldstable`, and `stable`.
- Pinned CI jobs to their matrix Go version with `GOTOOLCHAIN=local`.
- Skipped the Codecov upload for pull requests from forks.
- Improved the formatting check to list unformatted files.

### Deprecated

- `ptr.ToPtr()` and `ptr.FromPtr()` remain deprecated and will be removed in v2.0.0.

---

## [1.6.0] - 2026-09-12

### Added

- Added year-first, day-first, and month-first date parsing.
- Added extended and compact date-time formats with timezone support.

### Changed

- Moved `date` from `common/extension/date` to `common/date`.
- Expanded parsing for strings, byte slices, time values, and Unix timestamps.
- Updated `extension.Date()` and `extension.DateOr()` to accept date-order preferences.
- Updated related documentation, examples, and tests.

---

## [1.5.0] - 2026-09-12

### Changed

- Moved `float` from `common/extension/float` to `common/float`.
- Updated related package documentation, parsers, examples, and tests.

---

## [1.4.0] - 2026-09-04

### Changed

- Moved `object` from `common/extension/object` to `common/object`.
- Updated package documentation, examples, and tests.

---

## [1.3.0] - 2026-08-05

### Added

- Added the new `text` package for common string utilities.
- Added `text.PadLeft()`, `text.PadRight()`, and `text.PadCenter()`.
- Added `text.Truncate()` and `text.Ellipsis()`.
- Added `text.Capitalize()`, `text.Uncapitalize()`, and `text.Reverse()`.
- Added package documentation, README, examples, and unit tests.

---

## [1.2.1] - 2026-08-01

### Added

- Added `ptr.Of()` and `ptr.Deref()`.

### Changed

- Improved Go compatibility (Go 1.18+).
- Updated `ptr` documentation and examples.

### Deprecated

- Deprecated `ptr.ToPtr()` in favor of `ptr.Of()`.
- Deprecated `ptr.FromPtr()` in favor of `ptr.Deref()`.

---

## [1.2.0] - 2026-08-01

### Added

- Added `ptr.EmptyString()`.

### Changed

- Updated `ptr` documentation and examples.

---

## [1.1.0] - 2026-08-01

### Added

- Added the new `ptr` package for generic pointer utilities.
- Added comprehensive package documentation.
- Added examples for all exported functions.
- Added GitHub Actions CI workflow.
- Added Codecov integration.

### Changed

- Improved Common package documentation.
- Updated project README.
- Improved package consistency.