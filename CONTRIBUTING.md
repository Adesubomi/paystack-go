# Contributing to Paystack-Go

Thanks for your interest in contributing! 🎉  
This project aims to be a clean, well-tested, idiomatic Go SDK for interacting with the Paystack API.  
Contributions are welcome from everyone.

---

## 🧱 How to Contribute

### 1. Fork the Repository
Click **Fork** on GitHub and clone your fork:

```bash
git clone https://github.com/Adesubomi/paystack-go.git
cd paystack-go
```

### 2. Create a Feature Branch
```bash
git checkout -b feature/my-awesome-feature
```
Use descriptive branch names:
- fix/response-parsing
- feature/transaction-initialize
- chore/docs-update

### 3. Testing
All code must include tests.
- Run tests with:
    ```bash
    go test ./...
    ```
- Wherever application, please write both unit tests and integration tests (using httptest).
- Ensure new features maintain or increase coverage.
- Do not hit the real Paystack API during tests — use mocks or test servers.

### 4. Coding Guidelines
- Follow idiomatic Go conventions (gofmt, go vet).
- Do not expose unnecessary types or functions.
- All exported symbols must have clear, concise comments using the GoDoc format.
- API calls must accept context.Context.
- Return (value, error) consistently.
- Avoid adding heavy dependencies.

### 5. Commit Messages
- Use clear, descriptive commit messages
- For improved clarity, the use multi-line commits are favored and highly recommended. Here is an example
  ```text
  Payment Initialization
  
  This feature handles the processes required to initialize a payment request.
    
  Sub-features include:
  - Validating the request payload before sending it to the API.
  - Creating the HTTP request and attaching authentication headers.
  - Sending the request to the Paystack initialization endpoint.
  ```

### 6. Pull Requests
Before opening a PR:
1. Ensure your branch is rebased on main.
2. Ensure tests pass.
3. Document new functionality.
4. Provide a clear PR description explaining:
   - What changed
   - Why it changed
   - How it was tested

PRs that break tests or reduce quality will not be merged.

### 🙏 Thank You
Your contributions help make this SDK stable, reliable, and trusted by the Go community.
Thank you for taking the time to improve it!
