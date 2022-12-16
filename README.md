# How to write to an html file in python
Use the open file function to create the HTML file.
Add input data in HTML format into the file with the help of the write function.
Finally, save and close the file.
.......................................
errorhandling in Go:
"errors" is a type to represent any issue in the execution of a function.
i.e invalid input or prob connecting to external service.....
error type is an interface type that have one function.
type error interface {
    error() string
}
The most commonly-used error implementation is the errors package’s unexported errorString type.
type errorString struct {
    s string
}
func (e *errorString) Error() string {
    return e.s
}
 New returns an error that formats as the given text.
func New(text string) error {
    return &errorString{text}
}
Handle the errors:
errors.Is is better than ==
== is relatively error-prone and can only compare the current error type but can not unwrap. Therefore, errors.Is or errors.As are better choices.
............................................
Type Assertion:
Type Assertion provide access to the exact type of a variable of an interface.
If already the data type is present in the interface, then it will retrieve the acutal datatype value held by the interface.
A type assertion takes an interface value and extracts from it a value of the specified explicit type. Basically, it is used to remove the ambiguity from the interface variables.

Syntax-> t := value.(typeName)
a type assertion can return two values: the underlying value and a boolean value.
    t, ok := i.(T)
   
    If i holds a T, then t will be the underlying value and ok will be true.

   If not, ok will be false and t will be the zero value of type T, and no panic occurs.
   .........................................................................................
   Type Switches:
   A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.
