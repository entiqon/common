# Release Process

This document describes the release process for **Entiqon Common**.

Entiqon Common follows a lightweight release workflow designed for Go libraries.

---

# Versioning

Entiqon Common follows **Semantic Versioning (SemVer)**.

Format:

```text
MAJOR.MINOR.PATCH
```

Example:

```text
1.1.0
```

Meaning:

| Type | When to use |
|------|-------------|
| MAJOR | Breaking API changes |
| MINOR | New packages or backward-compatible features |
| PATCH | Bug fixes, documentation updates, and internal improvements |

---

# Preparing a Release

Before creating a new release, ensure the following:

1. The working tree is clean.

```bash
git status
```

2. All tests pass.

Entiqon Common is continuously validated using **GoTestX**, the official testing tool of the Entiqon ecosystem. We recommend using GoTestX during development for a richer testing experience, enhanced output, and integrated coverage reporting.

```bash
gotestx
```

Alternatively, the standard Go toolchain can be used:

```bash
go test ./...
```

3. Race detection passes.

```bash
go test -race ./...
```

4. Static analysis passes.

```bash
go vet ./...
```

5. Source code is properly formatted.

```bash
gofmt -w .
git diff --exit-code
```

6. Dependencies are up to date.

```bash
go mod tidy
```

7. Documentation is updated if necessary.

- `README.md`
- `CHANGELOG.md`
- Package `README.md`
- Examples (`example_test.go`)
- GoDoc (`doc.go`)

8. CI passes successfully.

---

# Updating the Changelog

Update `CHANGELOG.md` with the new version.

Example:

```markdown
## [1.1.0] - 2026-08-01

### Added

- Added the `ptr` package.
- Added GitHub Actions CI.
- Added Codecov integration.

### Changed

- Improved package documentation.
- Added runnable examples.
```

---

# Creating a Release

Create a signed Git tag.

```bash
git tag -s v1.1.0 -m "Release v1.1.0"
git push origin main
git push origin v1.1.0
```

GitHub Actions will automatically execute the CI workflow.

---

# Publishing a GitHub Release

Create the GitHub release using the GitHub CLI.

Generate release notes automatically:

```bash
gh release create v1.1.0 \
    --title "v1.1.0" \
    --generate-notes
```

Or provide custom release notes:

```bash
gh release create v1.1.0 \
    --title "v1.1.0" \
    --notes-file release-notes.md
```

Release notes should summarize:

- New packages
- New features
- Improvements
- Bug fixes
- Breaking changes (if any)

---

# Installation

Install the latest version:

```bash
go get github.com/entiqon/common@latest
```

Install a specific version:

```bash
go get github.com/entiqon/common@v1.1.0
```

---

# Verifying the Release

Verify that the tagged version can be downloaded:

```bash
go list -m github.com/entiqon/common@v1.1.0
```

Optionally create a temporary project and verify that imports resolve correctly.

---

# Release Guidelines

Maintain the following principles:

- Releases should be small and incremental.
- Follow Semantic Versioning.
- Avoid breaking public APIs within the same major version.
- Document all user-visible changes.
- Ensure documentation, examples, and tests remain synchronized.
- Never create a release if CI is failing.

---

# Maintainer

- **Original Author:** Isidro A. López G.
- **Organization:** Entiqon Project
- **Official Repository:** https://github.com/entiqon/common