## Create module

Here we create our own module (either lib/application)

```
go mod init <name> // name should meaning full if you want to publish in future
go mod init example.com/practice // onlu work local
```

It will create go.mod file which will write dependence just like node Package.json

To Import you local module we need to tell which module you mean

```
go mod edit -replace <import module name>=<path>
go mod edit -replace example.com/practice=../practice


after that you need to add version which is requires
# this will assign current verions
go mod tidy
```

To run File just type

```
. mean run file with main function and package
go run .
```

## Receiver

it will stucture which you can define as receving values of function

`func (d structureName) functionName`
