# ThunderGo
Go client for the [Thunderstore](https://thunderstore.io) API.\
You can call it **TSGO** for short ;)

> [!WARNING]
> This project is a **WIP** and may not work correctly and/or have missing features.\
> It is not advised to use this in production until there is a stable release.

# Installation
Enter the following line into your project's terminal.

```console
go get github.com/the-egg-corp/thundergo/
```

# Usage
This simple example illustrates how to interact with **ThunderGo**.
```go
import (
    TSGO "github.com/the-egg-corp/thundergo/experimental"
)

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

Visit the [wiki](https://github.com/the-egg-corp/thundergo/wiki) for the full documentation.

# Contact
Feel free to join my [discord](https://discord.gg/BwfzZpytjf) for support or suggestions.
