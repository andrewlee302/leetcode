func canConstruct(ransomNote string, magazine string) bool {
    ransomMap := make(map[byte]int)
    magazineMap := make(map[byte]int)
    for i := 0; i < len(ransomNote); i++ { ransomMap[ransomNote[i]]++ }
    for i := 0; i < len(magazine); i++ { magazineMap[magazine[i]]++ }
    for k, v := range ransomMap {
        if magazineMap[k] < v { return false }
    }
    return true
}
