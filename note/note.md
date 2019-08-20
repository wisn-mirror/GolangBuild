**内建类型**

bool string
int int8 int16 int32 int64 uintptr
byte rune//32位的没有char 
float32 float64 complex64 complex128 

iota自增


Fields Split Join
Contains Index
ToLower ToUpper
Trim TrimRight TrimLeft


任何类型在未初始化时都对应一个零值：布尔类型是 false，整型是 0，字符串是 ""，而指针，函数，interface，slice，channel 和 map 的零值都是 nil
