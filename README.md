> [!WARNING]
> This project is a **WIP** and until there is an initial release, using it is not advised.
> There is no LICENSE yet, so all rights are reserved.

# ThunderGo
> You can also call it **TSGO** for short ;)

Go client for the [Thunderstore](https://thunderstore.io) API.

# Installation
Enter the following line into your project's terminal.

```console
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

	cl, _ := pkg.Latest.GetChangelog()
	fmt.Println(cl)
}
```

Visit the [wiki](https://github.com/The-Egg-Corp/ThunderGo/wiki) for the full documentation.

# Contact
Feel free to join my [discord](https://discord.gg/BwfzZpytjf) for support or suggestions.
