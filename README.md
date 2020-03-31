# kk
try and make kubectl more betterer

## Setup

git clone https://github.com/mateo1647/kk
cd kk
go build
./kk (command)

kk
    service / svc
        prints service list for current namespace
    service -A
        prints service list for all namespaces

you can specify a "grep" like command to filter by service name

ex: kk svc argo -A
    this would print a list of all service names containing "argo" across all namespaces

hitting "enter" on the service will then output the selection with "-o yaml" option
