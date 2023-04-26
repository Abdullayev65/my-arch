package stream

import "golang.org/x/exp/constraints"

func Mapper[In any, Out any](ins []In, mapper func(In) Out) []Out {
	outs := make([]Out, len(ins))
	for i, in := range ins {
		outs[i] = mapper(in)
	}
	return outs
}

func SliceToMap[Obj any, Id constraints.Ordered](objs []Obj, getId func(Obj) Id) map[Id][]Obj {
	m := make(map[Id][]Obj, len(objs)>>3)
	var id Id
	for _, obj := range objs {
		id = getId(obj)
		if slc, ok := m[id]; ok {
			m[id] = append(slc, obj)
		} else {
			slc = make([]Obj, 1, 6)
			slc[0] = obj
			m[id] = slc
		}
	}
	return m
}

func ForEach[T any](objs []T, fun func(T)) {
	for _, obj := range objs {
		fun(obj)
	}
}

func AnyMatch[T any](objs []T, matcher func(T) bool) bool {
	for _, obj := range objs {
		if matcher(obj) {
			return true
		}
	}
	return false
}
