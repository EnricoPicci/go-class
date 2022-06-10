# Applicaton "order-manager-tax-calculator-integrated"

## package dependency structure

The <b>orderManager</b> package imports the <b>order</b> package and the <b>taxCalculator</b> package.

The <b>taxCalculator</b> package imports the <b>order</b> package.

The entire application is integrated and there is no decoupling.

                   orderManager
                    / \
                   /   \
        taxCalculator   \
                     \   \
                      \   |
                      order

### build

From the GO-CLASS project folder run the command
`go build -o ./bin/order-manager-tax-calculator-integrated ./src/decoupling-through-functions/order-manager-tax-calculator-integrated/main`

### run

From the GO-CLASS project folder run the command
`./bin/order-manager-tax-calculator-integrated`
