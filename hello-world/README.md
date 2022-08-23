When you write a program in Go you will have a main package defined with a main func inside it. Packages are ways of grouping up related Go code together.

The func keyword is how you define a function with a name and a body.

With import "fmt" we are importing a package which contains the Println function that we use to print.

It is good to separate your "domain" code from the outside world (side-effects). The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.

## Writing tests

- Writing a test is just like writing a function, with a few rules
- It needs to be in a file with a name like `xxx_test.go`
- The test function must start with the word Test
- The test function takes one argument only `t *testing.T`
- In order to use the `*testing.T` type, you need to import "testing", like we did with fmt in the other file

For now, it's enough to know that your t of type *testing.T is your "hook" into the testing framework so you can do things like t.Fail() when you want to fail.

You can read more about the placeholder strings in the fmt go doc. For tests %q is very useful as it wraps your values in double quotes.

## Go doc

Another quality of life feature of Go is the documentation. You can launch the docs locally by running godoc -http :8000 If you go to localhost:8000/pkg you will see all the packages installed on your system. The vast majority of the standard library has excellent documentation with examples. Navigating to http://localhost:8000/pkg/testing/ would be worthwhile to see what's available to you.

If you don't have godoc command, then maybe you are using the newer version of Go (1.14 or later) which is no longer including godoc. You can manually install it with go install golang.org/x/tools/cmd/godoc@latest.

## subtests

Here we are introducing another tool in our testing arsenal, subtests. Sometimes it is useful to group tests around a "thing" and then have subtests describing different scenarios.

## `t.Helper()`

t.Helper() is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper. This will help other developers track down problems easier. If you still don't understand, comment it out, make a test fail and observe the test output. Comments in Go are a great way to add additional information to your code, or in this case, a quick way to tell the compiler to ignore a line. You can comment out the t.Helper() code by adding two forward slashes // at the beginning of the line. You should see that line turn grey or change to another color than the rest of your code to indicate it's now commented out.
## Discipline

1. Let's go over the cycle again
2. Write a test
3. Make the compiler pass
4. Run the test, see that it fails and check the error message is meaningful
5. Write enough code to make the test pass
6. Refactor