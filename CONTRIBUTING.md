# Contributing to Mongorm

Welcome, contributor! We're excited to have you help make Mongorm even better. Here’s a guide to get you started with contributing.

## How to Contribute

### Reporting Issues

If you encounter a bug or have a feature request, please open an issue on our [GitHub Issues](https://github.com/mmycin/mongorm/issues) page. When reporting issues, try to provide:

- A clear and descriptive title.
- Steps to reproduce the issue.
- The expected and actual results.
- Any relevant logs or error messages.

### Submitting Pull Requests

We appreciate pull requests! Here’s how you can contribute code:

1. **Fork the Repository**: Click the "Fork" button at the top-right of this repository to create your copy.

2. **Clone Your Fork**: Clone your fork locally with:
    ```bash
    git clone https://github.com/your-username/mongorm.git
    ```

3. **Create a New Branch**: Create a new branch for your changes:
    ```bash
    git checkout -b your-branch-name
    ```

4. **Make Your Changes**: Make your changes to the code. Ensure you add tests for your changes if applicable.

5. **Commit Your Changes**: Commit your changes with a clear and descriptive message:
    ```bash
    git add .
    git commit -m "Your detailed description of changes"
    ```

6. **Push Your Changes**: Push your changes to your fork:
    ```bash
    git push origin your-branch-name
    ```

7. **Open a Pull Request**: Go to the [Pull Requests](https://github.com/mmycin/mongorm/pulls) page of the original repository and open a pull request. Provide a detailed description of your changes and any additional context.

## Code Style

- Follow Go's official [code style](https://golang.org/doc/effective_go).
- Ensure your code is formatted using `gofmt`.
- Use clear and descriptive variable and function names.

## Testing

Before submitting your pull request, make sure your changes pass all existing and new tests. Run the tests locally with:
```bash
go test ./...
```

## Documentation

If you make changes to the functionality of the code, please update the documentation accordingly. Ensure all changes are reflected in the `README.md` and other relevant documentation files.

## Style Guide

We follow Go’s standard style guide. If you are adding new features or changing existing ones, make sure they align with the project’s style.

## Thank You!

Thank you for contributing to Mongorm! Your efforts help improve the project and are greatly appreciated. If you have any questions, feel free to reach out via [GitHub Issues](https://github.com/mmycin/mongorm/issues) or our [discussion page](https://github.com/mmycin/mongorm/discussions).

