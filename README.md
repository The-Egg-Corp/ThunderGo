> [!WARNING]
> This project is a **WIP** and until there is a stable release, using it is not advised.

# ThunderGo
Go client for the [Thunderstore](https://thunderstore.io) API.

> You can call it **TSGO** for short ;)
> 
# Installation
Enter the following line into your project's terminal.

```console
go get github.com/The-Egg-Corp/ThunderGo
```

# Usage
This simple example illustrates how to interact with **ThunderGo**.
```go
import (
    TSGO "github.com/The-Egg-Corp/ThunderGo/experimental"
)

func main() {
    mod, err := TSGO.GetPackage("Owen3H", "CSync")

    if err != nil {
        fmt.Println(err)
        return
    }

    changelog, _ := mod.Latest.GetChangelog()
    fmt.Println(changelog)
}
```

Visit the [wiki](https://github.com/The-Egg-Corp/ThunderGo/wiki) for the full documentation.

# Contact
Feel free to join my [discord](https://discord.gg/BwfzZpytjf) for support or suggestions.
