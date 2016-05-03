## info
this is used for struct interface fields validating
```
type Add struct{
  name interface{} "int"
}
```
than call TypeValid(Add{name:3}) will return true.
this also can be used for other pkg types,like
```
import (
xxx/xxx/addr
)

type Addr struct{
  addr interface{} "addr.Name"
}

```
and, multi type 
```
type Addr struct{
  addr interface{} "addr.Name, addr.Address"
}

```
