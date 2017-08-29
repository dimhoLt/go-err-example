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

### Scope
Errors should always return errors that represent their package and their functionality, and may convert errors to
errors of its own type as seen fit.

An example is a package that talks to a database, but doesn't want the user of the package to care about the underlying
database implementation. Instead of then returning an `mgo.NotFound` error, which would break functionality if the
package would start returning `pgsql.NotFound`, it may convert those into its own errors with the same name.


## Examples
To show how the code is expected to work, run `$ go run main.go`. It will showcase both an implementation of an error
handling in code, and simple printing of errors - what most people currently do.

For an example on how this can be used in real life, run `$ go run server.go`, which will send the proper HTTP status
codes depending on the errors it receives.