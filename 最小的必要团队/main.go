package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(smallestSufficientTeam([]string{"algorithms", "math", "java", "reactjs", "csharp", "aws"},
		[][]string{{"algorithms", "math", "java"}, {"algorithms", "math", "reactjs"}, {"java", "csharp", "aws"}, {"reactjs", "csharp"}, {"csharp", "math"}, {"aws", "java"}}))

	//fmt.Println(smallestSufficientTeam([]string{"java", "nodejs", "reactjs"},
	//	[][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}}))
}

func smallestSufficientTeam(req_skills []string, people [][]string) []int {

	n, m := len(req_skills), len(people)

	skillMaskLen := 1 << uint(n)

	type pair struct {
		mn     int
		chosen bool
		lastJ  int
	}

	dp := make([][]pair, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]pair, skillMaskLen)
		for j := 0; j < skillMaskLen; j++ {
			dp[i][j] = pair{
				mn:     math.MaxInt32,
				chosen: false,
				lastJ:  j,
			}
		}
	}

	//dp[0][0] = 0
	for i := 0; i < m+1; i++ {
		dp[i][0].mn = 0
	}

	skillMap := make(map[string]int)
	for i, skill := range req_skills {
		skillMap[skill] = i
	}
	//fmt.Println(skillMap)

	peopleSkillMask := make([]int, m)
	for i, skills := range people {
		for _, skill := range skills {
			peopleSkillMask[i] = peopleSkillMask[i] | 1<<uint(skillMap[skill])
		}
	}
	//fmt.Println(peopleSkillMask)

	for i := 1; i < m+1; i++ {
		skillMask := peopleSkillMask[i-1]
		for j := 0; j < skillMaskLen; j++ {
			dp[i][j].mn = dp[i-1][j].mn
			// &^ -> and not, j &^ skillMask, 从集合 j 中移除集合 skillMask
			if dp[i][j].mn > dp[i-1][j&^skillMask].mn+1 {
				// 选这个人
				dp[i][j].mn = dp[i-1][j&^skillMask].mn + 1
				dp[i][j].chosen = true
				dp[i][j].lastJ = j &^ skillMask
			}
		}
	}

	//for i := 0; i < m; i++ {
	//	fmt.Println(dp[i])
	//}
	//fmt.Println(dp[m][skillMaskLen-1])

	ans := make([]int, 0, dp[m][skillMaskLen-1].mn)
	j := skillMaskLen - 1
	i := m
	for len(ans) != dp[m][skillMaskLen-1].mn {
		if dp[i][j].chosen {
			ans = append(ans, i-1)
			j = dp[i][j].lastJ
		}
		i--
	}

	l, r := 0, len(ans)-1
	for l < r {
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}

	return ans
}
