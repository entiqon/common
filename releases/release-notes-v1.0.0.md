# Release Notes — v1.0.0 (2025-09-19)

This is the **first stable release** of `entiqon/common`.  
It consolidates foundational utilities and extensions into a single, reusable package.

---

## ✨ Added
- **Extension Packages**
  - `boolean`: parsing helpers with extended tokens (`on/off`, `y/n`, `t/f`) and `BoolToString`.
  - `date`: deterministic parsing and cleaning (`CleanAndParse`, `CleanAndParseAsString`, strict `YYYYMMDD`).
  - `decimal`, `float`, `number`: safe parsers with fallback (`Or` variants) and rounding support.
  - `integer`: `ParseFrom(any)` and `IntegerOr`, consistent with other numeric parsers.
  - `collection`: helpers for slices and maps.
  - `object`: generic object utilities (`Exists`, `GetValue`, `SetValue`).
- **Parser Shortcuts**
  - `Or` variants (`BooleanOr`, `NumberOr`, `FloatOr`, `DecimalOr`, `DateOr`) for explicit fallback values.
- **Error Types**
  - Clearer error reporting with invalid-type messages.

---

## 🔄 Changed
- Object helpers moved from `common/object` → `common/extension/object`.
- `BoolToStr` renamed → `BoolToString` (old name kept for compatibility).
- Documentation improved:
  - Root-level `doc.go` and `README.md`.
  - Package-level docs and runnable examples.
  - Dockerfile updated to include extension docs.

---

## 🧪 Tests & Coverage
- 100% unit test coverage for all helpers and utilities.
- Extensive runnable examples in `example_test.go`.

---

## 📦 Availability
```bash
go get github.com/entiqon/common@v1.0.0
```
