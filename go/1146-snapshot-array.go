type SnapshotArray struct {
    snaps [][][2]int // <id, value>
    snapId int
}

func Constructor(length int) SnapshotArray {
    snaps := make([][][2]int, length)
    for i := 0; i < length; i++ {
        snaps[i] = [][2]int{{0,0}}
    }
    return SnapshotArray{snaps, 0}
}

func (this *SnapshotArray) Set(index int, val int)  {
    snap := this.snaps[index]
    lastItem := snap[len(snap)-1]
    if lastItem[0] != this.snapId && lastItem[1] != val {
        this.snaps[index] = append(this.snaps[index], [2]int{this.snapId, val})
    } else if lastItem[0] == this.snapId {
        snap[len(snap)-1][1] = val
    }
}

func (this *SnapshotArray) Snap() int {
    this.snapId++
    return this.snapId - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
    snap := this.snaps[index]
    // find the smallest pos whose id is > snap_id.
    pos := sort.Search(len(snap), func(i int) bool { return snap[i][0] > snap_id })
    return snap[pos-1][1]
}
