# SUMMARY

### Variable declaration
  - `var foo int`
  - `var foo int = 42`
  - `foo := 42`
### Can't redeclare variables, but can shadow them.
### All variable must be used.
### Visibility
  - lower case frist letter for package scope
  - upper case frist letter to export
  - no private scope

### Naming conventions
  - Pascal or camelCase
    - Capitaliaze acronyms (HTTP, URL)
  - As short as readsonable
    - longer names for longer lives

### Type conversions
  - destination Type(variable)
  - use `strconv` package for strings