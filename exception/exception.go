package exception

import "fmt"
//移动越过高原边界
var RIPPositionError = fmt.Errorf("RIP Exception: Move out of boundary!")
//处于失足风险
var SkipCmdByRipInfoError = fmt.Errorf("SKip move command, for Rip position.")