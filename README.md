# eazyE - A repetitive pattern in error handling
==========

All eazyE is is the following snippet of code
```
func New(s string, err error) (fErr error) {
  if err != nil {
    fmt.Errorf(s, err)
    return
  }
  return
}
```

JK, there's one more:

```
func Newl(s string, err error) {
  if err !=  nil {
  log.Println(fmt.Errorf(s, err))
  return
  }
}
```
