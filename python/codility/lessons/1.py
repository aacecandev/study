import re

def solution(N):
    return len(max(re.findall("0+1", format(N, "b")), default=[''], key=len)) - 1

def binary_array_search(A, X):
    left = 0
    right = len(A) - 1
    while left <= right:
        mid = (left + right) // 2
        if A[mid] == X:
            return mid
        elif A[mid] < X:
            left = mid + 1
        else:
            right = mid - 1
    return -1

def gener
