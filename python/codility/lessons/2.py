def solution(A, K):
    a = K % len(A)
    return A[-a:] + A[:-a]
