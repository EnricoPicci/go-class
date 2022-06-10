# Decoupling using functions

Decoupling can be obtained using functions and not necessarly always using interfaces.

In these examples we can see we can decouple the implementation of a certain piece of logic, the "taxation logic in this case",
using a design based on functions.

The main logic (`CalculatePrice(o order.Order)` in this case) calls the function stored in the `taxCalculateLogic` variable.

The actual function used is decided at startup by the `main` function which sets the value of `taxCalculateLogic`.

As a result, the actual logic applied to calculate the taxation can be developed decoupled from the core logic implementation with minimal dependencies from the core implementation (the only dependency is on the `order` package that provides the type of the input variable).

Therefore different teams can work on different packages autonomously with the minimal common dependency on the `order` package.
