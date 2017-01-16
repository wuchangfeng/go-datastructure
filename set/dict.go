package set

import "sync"

var pool = sync.Pool{}

// Set is an implementation of ISet using the builtin map type. Set is threadsafe.
type Set struct {
	// interface{} 为 key,,struct{}属于 value
	// items 属于 map 类型的
	items     map[interface{}]struct{}
	// 保证线程安全
	lock      sync.RWMutex
	// flattened 扁平
	flattened []interface{}
}

// 可以向 set 添加任意类型的，任意多个元素
func (set *Set) Add(items ...interface{}) {
	set.lock.Lock()
	defer set.lock.Unlock()

	set.flattened = nil
	// 迭代，将各个元素插入进去
	for _, item := range items {
		// item 用 key 存储元素，value 赋值为 null
		// 因为 map 中 key 不允许重复，刚好符合 set 的要求
		set.items[item] = struct{}{}
	}
}

// Remove will remove the given items from the set.
func (set *Set) Remove(items ...interface{}) {
	set.lock.Lock()
	defer set.lock.Unlock()

	set.flattened = nil
	for _, item := range items {
		delete(set.items, item)
	}
}

// Exists returns a bool indicating if the given item exists in the set.
func (set *Set) Exists(item interface{}) bool {
	set.lock.RLock()

	_, ok := set.items[item]

	set.lock.RUnlock()

	return ok
}

// 将 set 中的元素以 list 形式返回
func (set *Set) Flatten() []interface{} {
	set.lock.Lock()
	defer set.lock.Unlock()

	if set.flattened != nil {
		return set.flattened
	}

	// []X 表示 X 类型的切片，而 []interface{} 表示任意类型的切片
	set.flattened = make([]interface{}, 0, len(set.items))
	for item := range set.items {
		set.flattened = append(set.flattened, item)
	}
	return set.flattened
}

// 返回 set 中的长度
func (set *Set) Len() int64 {
	set.lock.RLock()

	size := int64(len(set.items))

	set.lock.RUnlock()

	return size
}

// 清除所有 items
func (set *Set) Clear() {
	set.lock.Lock()

	set.items = map[interface{}]struct{}{}

	set.lock.Unlock()
}

// All returns a bool indicating if all of the supplied items exist in the set.
func (set *Set) All(items ...interface{}) bool {
	set.lock.RLock()
	defer set.lock.RUnlock()

	for _, item := range items {
		if _, ok := set.items[item]; !ok {
			return false
		}
	}

	return true
}

// 彻底消除 set 对象，并且释放对象
func (set *Set) Dispose() {
	set.lock.Lock()
	defer set.lock.Unlock()

	for k := range set.items {
		delete(set.items, k)
	}

	// 避免引用对象占用内存
	for i := 0; i < len(set.flattened); i++ {
		set.flattened[i] = nil
	}

	set.flattened = set.flattened[:0]
	pool.Put(set)
}


// set 的构造函数，从 pool 池中获取 set 对象
func New(items ...interface{}) *Set {
	set := pool.Get().(*Set)
	for _, item := range items {
		set.items[item] = struct{}{}
	}

	return set
}

// pool 对象池，返回 set 对象
func init() {
	pool.New = func() interface{} {
		return &Set{
			items: make(map[interface{}]struct{}, 10),
		}
	}
}