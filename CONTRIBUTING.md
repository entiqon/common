# Contributing to Entiqon Common

Thank you for your interest in contributing to **Entiqon Common**.

Entiqon Common is part of the **Entiqon open-source ecosystem** and provides foundational packages shared across Entiqon projects.

- **Original Author:** Isidro A. López G.
- **Organization:** Entiqon Project
- **Official Repository:** https://github.com/entiqon/common

---

# Attribution & Intellectual Ownership

Entiqon Common was originally designed and implemented by **Isidro A. López G.** under the **Entiqon Project**.

All contributions must respect the following principles:

- The original copyright notice **must not be removed**.
- The original author attribution **must remain intact**.
- Derivative works should reference the **official repository**.

The MIT License permits reuse, modification, and redistribution under its terms.

If you reuse or fork this project, please preserve the project attribution:

```text
Original project: Entiqon Common
Author: Isidro A. López G.
Repository: https://github.com/entiqon/common
```

---

# Development Setup

Clone the repository:

```bash
git clone https://github.com/entiqon/common.git
cd common
```

Download dependencies:

```bash
go mod download
```

Run the recommended test suite using **GoTestX**, the official testing tool of the Entiqon ecosystem:

```bash
gotestx
```

Alternatively, use the standard Go toolchain:

```bash
go test ./...
```

---

# Coding Guidelines

Please follow standard Go conventions.

## Formatting

```bash
gofmt -w .
```

## Static Analysis

```bash
go vet ./...
```

## Tests

Every new feature or bug fix should include unit tests whenever practical.

New packages should also include:

- `doc.go`
- `README.md`
- `example_test.go`

Examples should be executable and verified through the Go testing framework.

---

# Package Design Principles

When adding new packages or APIs, follow these guidelines:

- Keep packages small and focused.
- Prefer idiomatic Go over unnecessary abstractions.
- Minimize external dependencies.
- Preserve backward compatibility within the current major version.
- Document every exported type, function, constant, and variable.

---

# Commit Guidelines

Follow the Conventional Commits specification.

Examples:

```text
feat(ptr): add generic pointer helpers
fix(errors): preserve wrapped causes
docs(common): improve package documentation
test(decimal): add parser edge cases
ci(common): add Codecov integration
```

---

# Pull Request Process

1. Fork the repository.
2. Create a feature branch.
3. Implement your changes.
4. Add or update tests.
5. Update documentation when appropriate.
6. Submit a Pull Request.

Pull requests should include:

- A clear description of the change.
- The motivation behind the change.
- Example usage when introducing new APIs.

---

# What Not to Contribute

Please avoid submitting pull requests that:

- Introduce unnecessary dependencies.
- Add unrelated functionality.
- Break existing public APIs without discussion.
- Duplicate existing functionality.
- Reduce documentation or test coverage.

---

# License

By contributing to this project, you agree that your contributions will be licensed under the **MIT License**.

See the `LICENSE` file for details.

---

# Maintainer

**Original Author:** Isidro A. López G.

**Organization:** Entiqon Project