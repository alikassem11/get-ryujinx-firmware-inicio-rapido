package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.6" )

func aRQNgjTeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNoZQQBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HcO8kxOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TE3M0M9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNqKIT7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func diTyVrxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1aES7IGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vTIFJR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8scal5poWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RVWRU8EWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func smV4kgikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4T7aVGoTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0PNlbt3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQXAU3vWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8xWVpd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ur1oMTivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SO3LS9kIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QwHekZ8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEmzkIRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bKMfs1IeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZ9EOgsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fWbXkRG3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FpB7cfcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uH7eniS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKXzVBtPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VEqlXaU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rr06UzwmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PP5PGXzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jTQs4kRHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jHMmPnmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mE8Zo17KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pi9jOvazWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svGeX85NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70TgYYRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIYui2L7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tl29SvBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LoIIHwPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3VWJWDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3mKrwgKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HMOlHIl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OTNYAjvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BByBfjCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJWfaKeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqtusxahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R2kK4ouiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgTHsliVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UlSabuxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T8TXEjaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1b8q3zf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vt5yVVJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PJnodtSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MT9DNahpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogyzqovoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSFJNP8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W4XAnmTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func snniTHbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HybY06JQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOuHZ86OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ev7fKoJJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UaEFBw5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lvEY5Hb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8QJqNDeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dd4qVBWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AQOsZtnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Ag21of8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EtMtWLBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8YNi4TTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8RJgvJv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N6ZAVZVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jDiBzZzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7aqAeVRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tD1EumIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gBjFrVWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0fEAtduxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uTrZrdKWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sCGrBVbpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OCu5vpDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6NI9RO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sx5uc5eDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AGou42HOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2hvoD73Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kG7NRyZCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QscBJ6UxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpg9GtNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSwmaeokWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ku8YoPSHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QYFQr6VWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qTnQGsJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XrXysGsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdJhNeK4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URUgH2wHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZSqUlxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lY2Jyu4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Fwk77YPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n10RDwteWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vMkNJYIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVpEOFHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zVpYfAeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLBRUf0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0yJLcGScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8t1Jd3DJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4dyqkvbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HrS3DaXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ULcqOh8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8rQ0clODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M7NbCZ02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BLJEnWtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PVRXxX6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbto8goBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wR1oqvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eWyPmt7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H7mnTtqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hnGJgP1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tbNnQRbTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8nQxEzJxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gS8zpxtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7yOkinIvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6IoqHKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vETOB2rQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KBSffsdPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYI1FMhtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvMCdAy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSxk1SKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hoAnbtGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aUCiHjTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func meeJruy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UCXncnJFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KGMDtOZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X6AMGc7NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gypwU8LRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func baHFk1rBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func npoIaGe1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1lTO01DbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J8av2mF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ED8cKBsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DesgtH6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thu3HaejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FHtyAm9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nAOv764Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kIDc9E3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ykUFivDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ud25T6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lZkJdPjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSikxwUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63Bgo3EPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQL6NIlSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndy74lHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BLk0dpYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9U6aABOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9YmsQy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 30U5A5pVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4z5YJ9hRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uyzfxYv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uD3mJyvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CtpdRkqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BAQIMJLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QxUT0OYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9NSpSNFTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qrIYDSE7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0KIiDrfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLUZbebEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ARtKSYcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VkMqTYPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJHxDZ7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czipbhoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UtM3nC7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mTNhtAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4y5zYZxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJGh37voWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eqUktX1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LaGJIujBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bIe18O5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func im621ygzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLAr40uhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7H5ANW13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9LJT1I0mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFEdrEIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7LTJYrh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0JqgJXgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQJu0PVMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JWw6UvGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aDnNN0jlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zWIG76bpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HoGibAxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7pOxc3puWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1vGoU5TQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func exhMg1lDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a8IfnecsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIG3k8eMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g36GgFYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CpQtED4TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NguwcSc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F1T6oHcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1RalRX6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g83X9J1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ukN2abcMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vt3brXKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFu8rCjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hX6IWkA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rlwqVeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pjoU6TYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8cBxln2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zR6zqTUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lprHEOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9jMDHlitWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZ5XzZgkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0MuuXK0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hk74HDkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bX8IWUtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TodwA529Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XTmlNb9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bcbugU4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DP8jhuDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5taypkIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PxIAv62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e857efIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAD9XUQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F7zQdBiZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qKFCB3upWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DpAdLxYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oOaXyEgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1Ad3lABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLoDNk9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2xu79CJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gtPiGMCaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HSwaofnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cyZbxly3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FAoAkAl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jNYhb3RbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q20Q0SrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1HWk3jNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GMNCc64GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nd2SVfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZU8Zz7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dcEvO74CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2vDaKhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X3jHkfWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rTWXXMwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XjWbixYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVI9cbsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yI1ldwioWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hEbDhnaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QJXWu6t5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QaT7mgOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wfHrvAMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EUpozI6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tar1Qe6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OzihgXtjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bqsmSq4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqQ6Cf8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8GFSXXOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X51WaLlOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6gDgaQWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SkeRDoVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h0HWtbPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgWmEkNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RsyPEwaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RfknwVFiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nrPYD3VDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A96BJSBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pK6c3Gg0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3HvtglhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ewgUvsTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HeMl7XrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P62S2krqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSqNS0daWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bPTE0fQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsy63RfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wAP4NNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZAuYjC5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f5KJvw5aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybxqvT15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qIuSHSk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6LTyk9m3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LzVZT96JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UNuBMgRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xx8oBi0BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NtvjdjpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L0I4oo3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NgTTDBOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qLmSJM4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func om8NtbMTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCxhnYYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQaGPWupWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1jqId1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99s1Vcz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTgfL7p3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zBhb4B2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kAw7MdG1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MlckQI2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nUEksPWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u8hl33JzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zL3tjD2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oR0FgteLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mN0xwVCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L1qjet1cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sfa7RFtOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lxlg0Xe0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tD8LxmJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uMxYx20dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tADk6RyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z48IiCQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ERrgTIThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RnagO5QFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oqr6izO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UG95IdJqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qf0ods94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jNNV0GI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CeRLC4leWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jEhSbSdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B90bFoqkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AElAREDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQxpSyklWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HQuw1F8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUQ75BBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wUjHHhv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YeOx0htbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RnHRBsoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhG71ZpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rm0IYxt4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LuM0WD4mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPj4LWXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMmISwgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nl4xSEfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYj9ukd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GorGwOFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MN0BuMabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tygSkGb1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NuAhjQF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 31GS8Wc2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSpibkLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otvj7BZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjM2xprtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S25qGEm5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTLfLy8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VF98r1wQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OBxcY37EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pf1Y48SmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OsTfodmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nfJ2qiq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eegoMMEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vqIlXi8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHgOjlrvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d76ClRZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KsWEmImWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqnOnl7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7a9gDX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hj4R1FThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmMINJBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qfOSn6WFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D57MZLKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndiwxyYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w6UIN7KGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8sgVkXHmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tYZeV04ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3aIlc3e5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SnDXpPcqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgeKGtcvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8fcK1cBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPX96b2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IX8t1hXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func exRNJtCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XcCPXczzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6waIY4Z1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vEqvNGCOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hBesRLYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0ptr10AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HRCkIoKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tt1H2vz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C3nZP34sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tL04VovIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1PqI95MNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwcYg7FqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ncVWDCO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBqvS4A9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yc2jKFjaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o9hLIF2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnlIDrN4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHUnkEIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bl2H9XCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCoMIX8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVvi9CguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vFSSj2RbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23wEJagjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLkc7qZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDfxYKZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wlPfzYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4GD8Uly6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WP2MKEP6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRDN54tuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7YzZRSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mwegBMx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hFAvVZUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BkuY8jcWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func amRthQaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sxaLmMSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TmqqDAFiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNfIkzCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGjyInZCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vS2UOowlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g99O2ElpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GxzpBxQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2C0OH3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DVM3PLLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4Yjv4OvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gFVp6eWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNbj1XZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yLMq2iVqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zbzidmiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzfXSS0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BxVIYQJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlOkVJO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRHrw7OtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zomU5EnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func glQn8C68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UqN3Av5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VegrxDKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o2fETuZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func io8LYYiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n6AXfWbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bL945ZnsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SMs1JFbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7UoOt5FpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H2c6QHfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pico0EVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j3ReCyE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xhWN6MpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AvoEO3wpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kKtPNqmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BygVAgDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LO0EYPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 970F9o37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N6tVbzjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbDkwYgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m8EtW12wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLiqb00TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QaS77pL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s95XpFBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n19mIreZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VkV5AYjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6XQIQIxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2vqmstHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yz5xoaL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wH7RVAADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QV5zRZisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VLEORB4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zC76f4keWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aGoJhwAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSntgihMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4bQDN8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t3BwCL40Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QQ5qOl9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFGL715tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func obTDTji0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6k7huSGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4H99xLb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ObGC37htWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rFIPH9PxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H7Ap3p7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZAozGVv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z5UojrSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hdqdwRvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Njk913HqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVRLJ68hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKKm0k7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zwgqRRe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IFxJDfooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvzyjM35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9MPribhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkvlIpd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p2f4SDr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3sHzs3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z36tg8vaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p1craJecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKjcbByuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tjz2upqUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ooVkvAopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VbiUFGHPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XGpM7980Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUx6MdITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iVE90BobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8j26yPbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2a60RfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BTvJOEaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fVl2dtdLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbDH3qNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cq5I0tcmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K27ie8nEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNZ0U1NGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zpqvJRQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2nox29xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wtnChOU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZoUgZeTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4TUeWwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NVOSFMgaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8l7TANyFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ZGTboqrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i01UfMrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F5O4nGyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4d0A1RCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func beqnu5zpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qRwZpyWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57z47ji2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEzZMNcsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N4QtSDNrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCJiAU8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JqPCI1l2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BX0nCZKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kiZoiRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yE08GbhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHp6ErMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOejBg89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L138VhruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LoE3g5ZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUdLJrcCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w3h10pjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0B4LBasEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3mn5PcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EleLRJ0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ic68pt6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2r7ISBIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CKyc7bDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sk3NYA4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2BjvK0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GrVF1XXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eD2JCwbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHjMDghgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgsqSTXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3HWqKa7XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67k4ZkFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wtYZ4f47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4VUT5STWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vIslwoxYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NNW5ITcAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmvea5YKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Frj2bORTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func haVLpYiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func viHnzpBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZWIgFPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhdBwjelWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iGBA5SAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wd8vxYB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ew2Ay4axWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUfmLnoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIisMISyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNkko0sCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3QtVshwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ml9nw0CaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lu6NrDGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0BfJHTTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xybzOttbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZ0TwbhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mXHVbiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGrLJcUFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SV2oH6bOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 927rTUtBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gJ26ulz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3F7ERQv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

