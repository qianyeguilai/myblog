package runner

var (
    RET_OK                      = 0
    RET_ERROR_MISS_PARAMETER    = 1


    RetDescription = map[int]string{
        0: "OK",
        1: "MISS PARAMETER",
    }
)
