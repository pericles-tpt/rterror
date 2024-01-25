# rterror
This library prepends additional runtime information to error messages, in order to make it easier to identify where the error originates at a glance.

## Example
An error that occurred in `package strconv`, in the function `ParseInt`, with a maximum identifier length for package and function names of 4 would output:
```
[strc_pi] parsing "a": invalid syntax 
```
