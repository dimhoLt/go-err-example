# Golang Error Handling

This project is intended to be a demo of error handling using Golang.

## Guidelines
When handling errors, the following guidelines should be adhered to:

### Type
Each package should have an `Error`-type. It must satisfy the `error`-interface, which is done simply by declaring a
function named `Error` that takes no parameters and returns a string.

### Structure
Each package's own `Error`-type should be implemented in an `errors.go`-file within the package. The trailing _s_ is
there since we will also initialize a few variables of potential errors that can be returned.

### Variables
Every error that the user may want to handle specifically should be initialized in the errors file, as specified above.

Every error variable, as just mentioned, should be named `Err*`, where the asterisk is replaced with a more detailed
description of the error's purpose.

### Formatting
The package should add it's own name, lower-case, to any errors it returns that is not its own - but only in exported
methods! This makes sure the package name is only added once.

### Returns
Regardless of the error type returned, the return value declared in the function should **always** be `error`, and never
the custom error declared in the package.
