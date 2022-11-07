# package circular dependencies

This is an example of an "import cycle" (circular dependency).

To see the "import cycle" error remove the comments in the file packagec/packagec.go and then run the command

The ciruclar dependency is created because:

- packagea imports packageb
- packageb imports packagec
- packagec imports packagea
