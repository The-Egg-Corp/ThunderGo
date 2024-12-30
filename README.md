# ThunderGo
Go client for the [Thunderstore](https://thunderstore.io) API.\
You can call it **TSGO** for short ;)

Currently, ThunderGo does not do any caching. Please implement this on your own end if required.

> [!WARNING]
> This project is a **WIP** and may not work correctly and/or have missing features.\
> It is not advised to use this in production until there is a stable release.

## Installation
Enter the following line into your project's terminal.

```console
go get github.com/the-egg-corp/thundergo/
```

## Import
To correctly import TSGO, you must choose which version of Thunderstore to interact with.\
Since you will likely be using V1 most of the time, you can just do the following:
```go
import TSGO "github.com/the-egg-corp/thundergo/v1"
```

However, you may need to rely on both versions, in which case I recommend doing something like so:
```go
import (
    TSGOV1 "github.com/the-egg-corp/thundergo/v1"
    TSGOExp "github.com/the-egg-corp/thundergo/experimental"
)
```

## Basic Examples
Visit the [wiki](https://github.com/the-egg-corp/thundergo/wiki) for the full documentation, or take a look at the [tests](./tests/).

### V1 
```go
var comm = TSGOV1.Community{
    Identifier: "lethal-company",
}

func main() {
    mod := comm.GetPackage("Owen3H", "CSync")
    if err != nil {
        fmt.Println(err)
        return
    }

    latest := mod.LatestVersion()
    fmt.Println(latest.VersionNumber)
}
```

### Experimental
```go
func main() {
    mod, err := TSGO.GetPackage("Owen3H", "CSync")
    if err != nil {
        fmt.Println(err)
        return
    }

    changelog, _ := mod.Latest.Changelog()
    fmt.Println(changelog)
}
```

# Contact
Feel free to join my [discord](https://discord.gg/BwfzZpytjf) for support or suggestions.
