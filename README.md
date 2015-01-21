# Bark wraps the log

Bark is a convenience wrapper around http://golang.org/pkg/log/, converting a
two- or three-liner error checking code into one-liners.

Some examples.

Instead of

```
if err != nil {
    log.Println(err.Error())
    return err
}
```

you can write

```
if err != nil {
    return barkLog.LogIf(err)
}
```

If you want to amend an error with some additional message, like this:

```
if err != nil {
    log.Printf("Some error has occurred while doing %q", action)
    log.Println(err.Error())
    return err
}
```

you can abbreviate:

```
if err != nil {
    return barkLog.LogIff(err, "Some error has occurred while doing %q", action)
}
```

If you only want to record a non-critical error and continue:

```
err := MaybeFailingAPI()
if err != nil {
    log.Println(err.Error)
}
Continue()
```

you can save some typing:

```
err := MaybeFailingAPI()
barkLog.LogIf(err)
Continue()
```
