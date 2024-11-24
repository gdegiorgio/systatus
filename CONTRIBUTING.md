# Contributing to Systatus

🎉 Welcome, and thank you for considering contributing to **Systatus**! Your contributions make it possible to build and improve this project, and we’re excited to work with you to make Systatus even better. 🎉

This guide provides an overview of our contribution process, from reporting issues to submitting code changes. Please read through this document before starting, and feel free to ask any questions if something isn’t clear.

---

## Table of Contents

1. [Code of Conduct](#code-of-conduct)
2. [How Can I Contribute?](#how-can-i-contribute)
   - [Reporting Issues](#reporting-issues)
   - [Suggesting Features](#suggesting-features)
   - [Improving Documentation](#improving-documentation)
   - [Contributing Code](#contributing-code)
3. [Development Workflow](#development-workflow)
   - [Getting Started](#getting-started)
   - [Making Changes](#making-changes)
   - [Pull Request Guidelines](#pull-request-guidelines)
4. [Contact](#contact)

---

## Code of Conduct

Please review our [Code of Conduct](CODE_OF_CONDUCT.md) to understand the expectations for behavior within our community. We are committed to fostering an inclusive and respectful environment for everyone.

---

## How Can I Contribute?

### Reporting Issues

If you find any bugs, glitches, or unexpected behavior, please let us know by [opening an issue](https://github.com/gdegiorigo/systatus/issues). When reporting an issue, be as specific as possible:

- Describe the issue clearly, including steps to reproduce it.
- Include details about your environment (Go version, operating system, etc.).
- Attach screenshots or logs if applicable.

### Suggesting Features

We’re always open to new ideas! If you have a feature suggestion, please check the existing [feature requests](https://github.com/gdegiorgio/systatus/issues) to see if it’s already been suggested. If not, [open a new issue](https://github.com/gdegiorgio/systatus/issues/new) and include:

- A clear description of the feature and its purpose.
- Any use cases where the feature could be helpful.
- Examples or mockups if relevant.

### Improving Documentation

We value high-quality documentation. If you see areas for improvement or clarification, feel free to submit updates! Documentation contributions are a great way to get started with open source. 

### Contributing Code

If you're ready to contribute code, start by checking the current [issues](https://github.com/gdegiorgio/systatus/issues) and see if any align with your skills and interests. Issues marked with `good first issue` are a great place to start.

---

## Development Workflow

### Getting Started

1. **Fork the Repository**: Fork the Systatus repository to your GitHub account.

2. **Clone Your Fork**:
   ```bash
   git clone https://github.com/gdegiorgio/systatus.git
   cd Systatus
   ```
3. **Set Up the Environment**: Ensure you have Go installed, and configure your local environment if necessary.

4. **Install Dependencies**:

```bash
go mod tidy
```

5. **Making Changes**

Create a Branch: Use a descriptive name for your branch. For example:

```bash
git checkout -b feature/improve-health-check
```

6. **Write Code and Tests**:

Follow the project's coding style and conventions.
Write unit tests for any new code or changes in functionality.

7. Run Tests: Ensure all tests pass before submitting:

```bash
go test ./...
```

Commit Your Changes: Use clear and concise commit messages following [Coventional Commit](https://www.conventionalcommits.org/en/v1.0.0/)

```bash
git commit -m "feat: implement /health"
```

8. **Push to Your Fork**:

```bash
git push origin feature/improve-health-check
```

9. **Open a Pull Request**


### Contacts
For questions or additional support, you can contact me using info provided in my [Github Profile Page](https://github.com/gdegiorgio)

Thank you for contributing to Systatus! Let’s work together to create something awesome. 🚀