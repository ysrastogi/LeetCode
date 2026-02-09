func searchInsert(arr []int, target int) int {
    n := len(arr)
    l := 0
    r := n-1
    mid := l+ (r-l)/2
    for l<=r{
        mid = l+ (r-l)/2
        if arr[mid]==target{
            return mid
        }else if arr[mid]>target{
            r = mid-1
        }else{
            l = mid+1
        }

    }
    if arr[mid]>target{
        return mid
    }
    return mid+1
    
}

// Approach:
// -> search for space