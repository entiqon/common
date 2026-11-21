# Contributing to Entiqon Common

Thank you for your interest in contributing to **Entiqon Common**.

Entiqon Common is part of the **Entiqon open-source ecosystem** and provides foundational packages shared across Entiqon projects.

- **Original Author and Copyright Holder:** Isidro A. López G.
- **Organization:** [Entiqon](https://github.com/entiqon)
- **Official Repository:** https://github.com/entiqon/common

---

# Attribution & Intellectual Ownership

Entiqon Common was originally designed and implemented by **Isidro A. López G.**, who holds its copyright. It is maintained under the [Entiqon](https://github.com/entiqon) GitHub organization.

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

## Go Version

All packages support **Go 1.18 or later**. CI tests Go 1.18, `oldstable`, and `stable`,
so avoid language features and standard library APIs newer than Go 1.18.

## Formatting

```bash
gofmt -w .
```

## Static Analysis

```bash
go vet ./...
```

Code should produce no warnings from `gofmt`, `go vet`, or gopls.

## Tests

Every new feature or bug fix should include unit tests whenever practical.

Total test coverage must stay at or above **80%**; CI fails below that.

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

# Project Roles

| Role            | Who                          | Responsibilities                                                      |
|-----------------|------------------------------|-----------------------------------------------------------------------|
| **Owner**       | Isidro A. López G.           | Sets direction, reviews and approves Pull Requests, may write code.   |
| **Maintainer**  | Developers with write access | Write code, review the owner's Pull Requests, merge, and release.     |
| **Contributor** | Anyone else                  | Submit Pull Requests from a fork.                                     |

---

# Pull Request Process

1. Fork the repository (maintainers can branch in the main repository instead).
2. Create a branch from `main`:
    - `features/<short-name>` for new functionality.
    - `fixes/<short-name>` for bug fixes.
    - `chores/<short-name>` for CI, documentation, and tooling.
3. Implement your changes.
4. Add or update tests.
5. Update documentation when appropriate.
6. Add an entry under `## [Unreleased]` in `CHANGELOG.md` for any change users would notice.
7. Submit a Pull Request against `main`.

Pull requests should include:

- A clear description of the change.
- The motivation behind the change.
- Example usage when introducing new APIs.

---

# Reviews

- Every Pull Request needs **one approving review** and **passing CI**.
- The owner is requested as a reviewer automatically through `.github/CODEOWNERS`.
- When the **owner** writes the code, a **maintainer** reviews it.
- Reviewers approve, request changes, or comment. They do not push to the author's
  branch without asking.

---

# Merging

Pull Requests are merged with **Squash and merge** only, so each one becomes a single
commit on `main`. The squash commit message must follow Conventional Commits and
describe the whole change, not list work-in-progress commits.

| Pull Request written by | Reviewed by  | Merged by                       |
|-------------------------|--------------|---------------------------------|
| A maintainer            | The owner    | The author, after approval      |
| The owner               | A maintainer | The reviewing maintainer        |
| A contributor (fork)    | The owner    | Any maintainer, after approval  |

Delete the branch after merging.

---

# Releases

Maintainers make releases from a `main` commit whose CI has passed.

1. In `CHANGELOG.md`, rename `## [Unreleased]` to `## [X.Y.Z] - YYYY-MM-DD` and merge
   that change.
2. Tag the merged commit and push the tag:

   ```bash
   git tag -a vX.Y.Z -m "vX.Y.Z"
   git push origin vX.Y.Z
   ```

3. The **Release** workflow validates the tagged code and creates a **draft
   pre-release** with notes from the matching `CHANGELOG.md` section.
4. Review the draft, optionally replace the notes with the detailed release notes,
   then publish it:

   ```bash
   gh release edit vX.Y.Z --notes-file releases/release-notes-vX.Y.Z.md
   gh release edit vX.Y.Z --draft=false
   ```

Never delete or move a published tag. Go's module proxy caches versions permanently,
so a mistake is fixed by releasing the next patch version instead.

Breaking changes wait for the next major version. Until then, mark the old API with a
`// Deprecated:` notice that names its replacement.

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

**Original Author and Owner:** Isidro A. López G.

**Organization:** [Entiqon](https://github.com/entiqon) on GitHub