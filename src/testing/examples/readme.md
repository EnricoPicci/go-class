# Examples

Examples are functions whose name starts with the prefix Example and that are contained in files with \_test postfix. Examples are functions defined in a package that follows the \_text postfix convention.

After the Example prefix the name of an example has either to be the name of an exported function or the name of an exported type followed by \_ and the name of an exported method.

Examples can be executable. In order to be executed they have to have an // Output: directive followed by the expected output. Anything that is printed on the stdout is compared with the value defined by the directive. If there is no match the example fails and is signalled as any failed test.

Examples that have the // Output: directive are executed by the go test command.
