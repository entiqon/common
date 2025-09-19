# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/)
and this project adheres to [Semantic Versioning](https://semver.org/).

## [v1.0.0] - 2025-09-19

### Added
- **Extension Packages**:
    - `boolean`: parsing helpers with extended tokens (`on/off`, `y/n`, `t/f`) and `BoolToString` (renamed from `BoolToStr`).
    - `date`: deterministic parsing and cleaning (`CleanAndParse`, `CleanAndParseAsString`, strict `YYYYMMDD`).
    - `decimal`, `float`, `number`: safe parsers with fallback (`Or` variants) and rounding support.
    - `integer`: `ParseFrom(any)` and `IntegerOr`, consistent with other numeric parsers.
    - `collection`: helpers for slices and maps.
    - `object`: generic object utilities (`Exists`, `GetValue`, `SetValue`) with flexible type support.
- **Parser Shortcuts**:
    - Added `Or` variants (`BooleanOr`, `NumberOr`, `FloatOr`, `DecimalOr`, `DateOr`) to provide explicit fallback values.
- **Error Types**:
    - Enhanced error reporting with clear invalid-type messages.

### Changed
- **Object Helpers**: moved from `common/object` → `common/extension/object` and normalized APIs.
- **Boolean Helpers**: `BoolToStr` renamed → `BoolToString`; old name retained for backward compatibility.
- **Documentation**:
    - Root-level `doc.go` and `README.md` describing purpose and subpackages.
    - Package-level `README.md`, `doc.go`, and `example_test.go` for each extension.
    - Dockerfile updated to copy extension docs into site builds.
    - Navigation updated with alphabetized extension packages.

### Tests & Coverage
- Extensive unit tests for all parsing helpers, `Or` fallbacks, and object utilities.
- 100% coverage across extension packages, aligned with Entiqon quality standards.
- Runnable examples provided in `example_test.go`.

---
