package queue

type quen []interface{}
//
func ( q *quen) Push(value interface{})  {
	*q=append(*q,value)
}
func ( q *quen) Pop()interface{}  {
	value:=(*q)[0]
	*q=(*q)[1:]
	return value
}
func ( q *quen) isEmpty()bool  {
	return len(*q)==0
}
