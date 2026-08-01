package main

import (
	"fmt"
)

func main() {
	// Sample texts and patterns
	testCases := []struct {
		text    string
		pattern string
	}{
		{"ABABDABACDABABCABAB", "ABABCABAB"},
		{"AABAACAADAABAABA", "AABA"},
		{"GEEKSFORGEEKS", "GEEK"},
		{"AAAAAA", "AA"},
	}

	// Test each pattern matching algorithm
	for i, tc := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Text: %s\n", tc.text)
		fmt.Printf("Pattern: %s\n", tc.pattern)

		// Test naive pattern matching
		naiveResults := NaivePatternMatch(tc.text, tc.pattern)
		fmt.Printf("Naive Pattern Match: %v\n", naiveResults)

		// Test KMP algorithm
		kmpResults := KMPSearch(tc.text, tc.pattern)
		fmt.Printf("KMP Search: %v\n", kmpResults)

		// Test Rabin-Karp algorithm
		rkResults := RabinKarpSearch(tc.text, tc.pattern)
		fmt.Printf("Rabin-Karp Search: %v\n", rkResults)

		fmt.Println("------------------------------")
	}
}

// NaivePatternMatch performs a brute force search for pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func NaivePatternMatch(text, pattern string) []int {
	if len(pattern) == 0 || len(pattern) > len(text) {
    return []int{}
}
	// TODO: Implement this function
	n:=len(text)
	m:=len(pattern)
	pat := []int{}
	for i:=0;i<=n-m;i++{
	    j:=0
	    for j<m{
	        if text[i+j] != pattern[j]{
	            break
	        }
	        j++
	    }
	    if j==m{
	        pat = append(pat,i)
	    }
	}
	return pat
}

// KMPSearch implements the Knuth-Morris-Pratt algorithm to find pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func KMPSearch(text, pattern string) []int {
	if len(pattern) == 0 || len(pattern) > len(text) {
    return []int{}
}
	// TODO: Implement this function
	
	res := []int{}
	n := len(text)
	lenp := 0
	m := len(pattern)
	lps := make([]int, m)
	
	lps[0]=0
	
	i:=1
	for i<m{
	    if pattern[i] == pattern[lenp]{
	        lenp += 1
	        lps[i] = lenp
	        i++
	    }else{
	        if lenp != 0{
	            lenp = lps[lenp-1]
	        }else{
	            lps[i] = 0
	            i++
	        }
	    }
	}
	
	j := 0
	k := 0
	
	for j<n{
	    if text[j] == pattern[k]{
	        j++
	        k++
	        if k == m {
	            res = append(res,j-k)
	            k = lps[k-1]
	        }
	    }else{
	        if k != 0{
	            k = lps[k-1]
	        }else{
	            j++
	        }
	    }
	} 
	
	return res
}

// RabinKarpSearch implements the Rabin-Karp algorithm to find pattern in text.
// Returns a slice of all starting indices where the pattern is found.
func RabinKarpSearch(text, pattern string) []int {
	if len(pattern) == 0 || len(pattern) > len(text) {
    return []int{}
}
	// TODO: Implement this function
	d := 256
	q := 101
	m := len(pattern)
	n := len(text)
	p := 0
	t := 0
	h := 1
	ans := []int{}
	

	for i:=0;i<m-1;i++{
		h = (h*d)%q
	}

	for i:=0;i<m;i++{
		p = (d * p + int(pattern[i])) % q
		t = (d * t + int(text[i])) % q
	}

	for i:=0;i<n-m+1;i++{
		if p==t {
			match := true
			for j:=0;j<m;j++{
				if text[i+j] != pattern[j]{
					match = false
					break
				}
			}
			if match {
				ans = append(ans, i)
			}
		}
		if i< n-m{
			t = (d* (t - int(text[i])*h) + int(text[i+m])) % q
			if t<0{
				t += q
			}
		}
	}

	return ans
}
