# Security Policy

## Supported Versions

Entiqon Common provides security updates for the most recent releases.

| Version | Supported |
|---------|:---------:|
| Latest release | ✅ |
| Previous minor release | ⚠️ Best effort |
| Older versions | ❌ |

Users are strongly encouraged to upgrade to the latest stable release to receive security and reliability updates.

---

# Reporting a Vulnerability

If you discover a security vulnerability in Entiqon Common, please report it responsibly.

**Do not** disclose security vulnerabilities through public GitHub issues or discussions.

Instead, report the issue privately to the project maintainer.

## Contact

- **Maintainer:** Isidro A. López G.
- **Project:** https://github.com/entiqon/common

Please include:

- A description of the vulnerability.
- Steps to reproduce the issue.
- The potential impact.
- A suggested mitigation, if known.

We will acknowledge receipt of your report as soon as reasonably possible and investigate the issue promptly.

---

# Security Considerations

Entiqon Common is a collection of reusable Go libraries intended to be embedded into applications.

The project follows these security principles:

- Keep dependencies to a minimum.
- Avoid introducing unnecessary attack surface.
- Follow secure coding practices.
- Address reported vulnerabilities promptly.
- Preserve backward compatibility whenever practical while delivering security fixes.

Applications using Entiqon Common remain responsible for their own security, configuration, deployment, authentication, authorization, and operational controls.

---

# Security Updates

Security fixes are released following Semantic Versioning whenever possible.

Critical vulnerabilities may result in an out-of-band patch release to ensure users can upgrade quickly.