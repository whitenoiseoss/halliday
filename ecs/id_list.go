package ecs

type IDList []uint64

func (list IDList) Len() int           { return len(list) }
func (list IDList) Swap(i, j int)      { list[i], list[j] = list[j], list[i] }
func (list IDList) Less(i, j int) bool { return list[i] < list[j] }
