# Applicaton "order-manager-tax-calculator-decoupled"

## Problem

There are different ways to calculate taxes and the right implementation depends on the specific configuration
used to launch the app (e.g. on the country which is set at startup).

## package dependency structure

Each tax calculation logic is implemented in a separate package. All tax implementation packages use the <b>order</b> type and therefore
import the <b>order</b> package.

The <b>orderManager</b> package imports the <b>order</b> package but none of the tax implementation packages available. It exports though a function that allows to set a variable of the type of the function used to calculate the tax.

The <b>main</b> package imports the <b>orderManager</b> package and all the packages that implement tax calculation logic and sets, at startup, which
is the logic (i.e. the specific function belonging to a specific package) to use.

The entire application is integrated and there is no decoupling.

                       main
                ______ / | \______
                |        |        |
                |   orderManager  |
                |        |        |
                |        |        |
                |        |        |
    taxCalculatorSimple  | taxCalculatorComplex
                     \   |  /
                      \  | /
                      order

### build

From the GO-CLASS project folder run the command
`go build -o ./bin/order-manager-tax-calculator-decoupled ./src/decoupling-through-functions/order-manager-tax-calculator-decoupled/main`

### run

From the GO-CLASS project folder run the command
`./bin/order-manager-tax-calculator-decoupled`
