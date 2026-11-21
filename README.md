<h1>
    <img src="https://github.com/entiqon/entiqon/blob/main/assets/entiqon_sharicon.png?raw=true" align="center" height="64" width="64" alt="logo"> Common Module
</h1>

[![Go Reference](https://pkg.go.dev/badge/github.com/entiqon/common.svg)](https://pkg.go.dev/github.com/entiqon/common)
[![Go Version](https://img.shields.io/badge/Go-1.18%2B-blue)](https://go.dev)
[![Release](https://img.shields.io/github/v/release/entiqon/common)](https://github.com/entiqon/common/releases)
[![Official Repository](https://img.shields.io/badge/common-blue?logo=github)](https://github.com/entiqon/common)
[![Build Status](https://github.com/entiqon/common/actions/workflows/ci.yml/badge.svg)](https://github.com/entiqon/common/actions)
[![codecov](https://codecov.io/gh/entiqon/common/graph/badge.svg?token=3bK4fVc4xS)](https://codecov.io/gh/entiqon/common)
[![License](https://img.shields.io/github/license/entiqon/common)](LICENSE)

## 🌱 Overview

The **common** module provides foundational shared utilities and reusable
components used across the Entiqon ecosystem.

It promotes consistency, code reuse, and well-tested implementations for common
development tasks.

---

## 🚀 Quick Start

```bash
go get github.com/entiqon/common
```

---

## 📦 Subpackages

| Package                                     | Description                                                        |
|---------------------------------------------|--------------------------------------------------------------------|
| [🧩 common/boolean](./boolean/README.md)     | Boolean parsing from common representations and string formatting. |
| [🧩 common/date](./date/README.md)           | Date and timestamp parsing and formatting utilities.               |
| [🧩 common/errors](./errors/README.md)       | Structured error types with cause, reason, and process stage.      |
| [🧩 common/extension](./extension/README.md) | Convenience conversions with zero-value or custom fallbacks.       |
| [🧩 common/float](./float/README.md)         | Floating-point parsing and conversion utilities.                   |
| [🧩 common/object](./object/README.md)       | Dynamic access and manipulation for maps and structs.              |
| [🧩 common/ptr](./ptr/README.md)             | Generic pointer utilities for working with optional values.        |
| [🧩 common/testkit](./testkit/README.md)     | Dependency-free helpers for writing and testing Go tests.          |
| [🧩 common/text](./text/README.md)           | String padding, truncation, and presentation utilities.            |

---

## 📄 License

[MIT](LICENSE) — © Entiqon Labs
