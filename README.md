# Paystack-Go SDK

[![golangci-lint](https://github.com/Adesubomi/paystack-go/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/Adesubomi/paystack-go/actions/workflows/golangci-lint.yml)

A Go SDK for interacting with the Paystack API. This project is not stable and is currently in active development.

---

## Table of Contents
- Installation
- Usage
- Features
- Code Quality
- Contributing
- License

---

## Code Quality

We use **golangci-lint** under the hood to maintain code quality, enforce consistent formatting, and catch potential issues early.

For development, use the Makefile commands to ensure the code passes all quality checks:

```bash 
make lint
make fmt
make test
```

These commands run linters, format the code consistently, and execute tests, keeping the codebase reliable and maintainable.

All pull requests and pushes to `main` are automatically checked with **golangci-lint** in our CI pipeline, ensuring that contributions adhere to the same quality standards.

---

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on contributing to this project.

---

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

---

## Code of Conduct

This project is governed by a [Code of Conduct](CODE_OF_CONDUCT.md), which is adapted 
from the [**Contributor Covenant**](https://www.contributor-covenant.org), version 2.1.
All participants, including maintainers, contributors, and users, are expected to
follow these guidelines to foster a welcoming and inclusive community.
