# kk
try and make kubectl more betterer

## Setup

1. git clone https://github.com/matoszz/kk
2. cd kk
3. go build
4. ./kk (command)

kk
1. service / svc
    1. prints service list for current namespace
2. service -A
    1. prints service list for all namespaces

you can specify a "grep" like command to filter by service name

1. kk svc argo -A
    1. this would print a list of all service names containing "argo" across all namespaces

hitting "enter" on the service will then output the selection with "-o yaml" option


Inspiration / credit:
- ckube
- kubectl-grep
