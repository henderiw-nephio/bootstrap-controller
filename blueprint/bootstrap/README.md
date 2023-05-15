# bootstrap

## Description
bootstrap controller

## Usage

### Fetch the package
`kpt pkg get REPO_URI[.git]/PKG_PATH[@VERSION] bootstrap`
Details: https://kpt.dev/reference/cli/pkg/get/

### View package content
`kpt pkg tree bootstrap`
Details: https://kpt.dev/reference/cli/pkg/tree/

### Apply the package
```
kpt live init bootstrap
kpt live apply bootstrap --reconcile-timeout=2m --output=table
```
Details: https://kpt.dev/reference/cli/live/
