go cron
-------

golang [cron](https://en.wikipedia.org/wiki/Cron) like service

## Supported system

    Linux
    Windows 8
    Windows 7
    Windows XP
    OS X
    
## Features
    
    ┌───────────── sec (0 - 59)
    │ ┌───────────── min (0 - 59)
    │ │ ┌────────────── hour (0 - 23)
    │ │ │ ┌─────────────── day of month (1 - 31)
    │ │ │ │ ┌──────────────── month (1 - 12)
    │ │ │ │ │ ┌───────────────── day of week (0 - 6)
    │ │ │ │ │ │                                         
    │ │ │ │ │ │
    │ │ │ │ │ │
    * * * * * *  

## Clone

```bash
    git clone https://github.com/e154/cron.git /path/to/clone   
```

## init

```bash
    go get github.com/e154/cron   
```

```go

    import (
        "time"
        "strings"
        "strconv"
        "github.com/e154/cron"
    )
    
    cron := NewCron()
    
    cron.NewTask("* * * * * *", func(){log.Println("task1")})
    cron.NewTask("10,44 * * * * *", func(){log.Println("task2")})
    cron.NewTask("5 * * * * *", func(){log.Println("task3")})
    //cron.NewTask("*/5 * * * * *", func(){log.Println("task4")})	// not supported
    cron.NewTask("1-45,47,50 * * * * *", func(){log.Println("task5")})
    cron.NewTask("45-59,20 * * * * *", func(){log.Println("task6")})
    cron.NewTask("0 23 * * * *", func(){log.Println("task7")})
    cron.NewTask("* 24 16 * * *", func(){log.Println("task8")})
    cron.NewTask("0 24 16 * * *", func(){log.Println("task9")})

    cron.Run()

    for {
        time.Sleep(time.Second)
    }
```

#### LICENSE

go cron is licensed under the [MIT License (MIT)](https://github.com/e154/cron/blob/dev/LICENSE)

