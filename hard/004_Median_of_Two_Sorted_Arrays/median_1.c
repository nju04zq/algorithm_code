// find kth (0 <= k < n)
// if a[m/2] < b[n/2]
// #1 then at least (m+n)/2 elements larger than a[m/2],
//    that is, at most (m+n)/2 elements smaller than a[m/2]
// #2 then at least (m+n)/2 elements smaller than b[n/2]
// #1 if k >= (m+n)/2, then we can drop a[0] to a[m/2] due to above #1
// #2 if k <= (m+n)/2, then we can drp b[n/2] to b[n] due to above #2
