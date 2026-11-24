# Search in Rotated Sorted Array

> Striver SDE Sheet 064 · Day 11 Binary Search — at any midpoint one half is still sorted; range-check the target against that half to pick a side. Go, stdlib only, tested.

## Problem

A sorted array of distinct values was rotated at an unknown pivot (`[4,5,6,7,0,1,2]`). Find a target's index in O(log n) without first locating the pivot.

## Approach

Rotation breaks the array into two sorted runs, and any midpoint lands in one of them — meaning at least one side of mid is fully sorted. That's decidable locally: `nums[lo] <= nums[mid]` says the left half is sorted. A sorted half supports an exact range test (`lo-value ≤ target < mid-value`), so either the target provably lives there (search it) or it provably doesn't (search the other half). One comparison to classify, one to route: still a clean halving each step, no pivot hunt needed.

**Complexity:** O(log n) time, O(1) space. (Duplicates would degrade the classification to O(n) worst case — this variant assumes distinct values.)

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
