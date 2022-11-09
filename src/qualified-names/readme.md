# qualified names

This is an example of using qualified names when importing a package.

`packagea` imports `packageb` and `packagecba`.

If you open the file [packageAcode.go](./packagea/packageAcode.go) you can see how the package names are used in the qualified names.

It is also worth noting that the folder containing `packagecba` is named `packagec`, in other words package name and folder name are not the same. The consequence is that the import of package `packagecba` must be the following

`packagecba "github.com/EnricoPicci/go-class/src/qualified-names/packagec"`

On the contrary, `packageb` is contained in the folder `packageb`, i.e. the name of the package and the name of the folder are the same, therefore the import statement is simpler

`"github.com/EnricoPicci/go-class/src/qualified-names/packageb"`
