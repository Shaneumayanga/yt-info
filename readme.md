# Usage
```
go get github.com/Shaneumayanga/yt-info

```

```go
package main

import (
    "fmt"

    ytdl "github.com/Shaneumayanga/yt-info"
)

func main() {
    info := ytdl.Info("https://www.youtube.com/watch?v=n10Ntw06xBA")
    fmt.Printf("video id %s \n", info.Id)
    fmt.Printf("title %s \n", info.Title)
    for _, format := range info.Formats {
        fmt.Printf("Format %s \n", format.Format_note)
        fmt.Printf("Download URL %s \n", format.Url)
    }
}
```


