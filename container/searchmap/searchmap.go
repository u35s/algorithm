package searchmap

type SearchMap struct {
	dataMap map[string]interface{}
	idxMap  map[string][]interface{}
}

func (this *SearchMap) Init() {
	this.dataMap = make(map[string]interface{})
	this.idxMap = make(map[string][]interface{})
}

func (this *SearchMap) Size() int    { return len(this.dataMap) }
func (this *SearchMap) IdxSize() int { return len(this.idxMap) }

func (this *SearchMap) Add(k string, v interface{}) {
	if _, ok := this.dataMap[k]; !ok {
		this.dataMap[k] = v
		this.execKeyIdx(k, func(idx string) {
			this.idxMap[idx] = append(this.idxMap[idx], v)
		})
	}
}

func (this *SearchMap) Remove(k string, v interface{}) {
	if _, ok := this.dataMap[k]; ok {
		delete(this.dataMap, k)
		this.execKeyIdx(k, func(idx string) {
			if slc, ok := this.idxMap[idx]; ok {
				for i := range slc {
					if v == slc[i] {
						this.idxMap[idx] = append(slc[:i], slc[i+1:]...)
						break
					}
				}
			}
		})
	}
}

func (this *SearchMap) Search(k string) ([]interface{}, bool) {
	//TODO %,_,[],[^]通配符查询,B TREE
	v, ok := this.idxMap[k]
	return v, ok
}

func (this *SearchMap) execKeyIdx(k string, exec func(idx string)) {
	runes := []rune(k)
	added := make(map[string]bool)
	for i := 1; i <= len(runes); i++ {
		for j := i - 1; j >= 0; j-- {
			sub := string(runes[j:i])
			if _, ok := added[sub]; !ok {
				added[sub] = true
				exec(sub)
			}

		}
	}
}
