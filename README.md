# ThunderGo
> You can also call this **TSGO** ;)

Go client for the [Thunderstore](https://thunderstore.io) API.\

# Installation
Enter the following line into your project's terminal.

```bash
go get github.com/The-Egg-Corp/ThunderGo
```

# Usage
This simple example illustrates how to interact with **ThunderGo**.
```go
import (
    TSGO "github.com/The-Egg-Corp/ThunderGo"
)

func main() {
    mod, err := TSGO.GetPackage("Owen3H", "CSync")

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(mod.Latest.GetChangelog())
}
```

Visit the [wiki](https://github.com/The-Egg-Corp/ThunderGo/wiki) for the full documentation.

# Contact
Feel free to join my [discord]() for support or suggestions.