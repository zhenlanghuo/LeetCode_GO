package main

import "fmt"

func main() {
	//fmt.Println(similarPairs([]string{"aba", "aabb", "abcd", "bac", "aabc"}))
	fmt.Println(similarPairs([]string{"ytnuzbuvbeznygmljiqdswsanxpncswgcdtoujixzridvjxdza","kpqxkpgbqhgkuetkwcnwamwgllgfwjyrcdhvehppabubfpukfu","tubebzuitpibmplzecrtbyjbfixltwltnyzoshbhlgqxnpqmvq","oaciiaurgpvzfrtruhiemmybgdkvfgldqlchyrsavlddfezfaw","tdwrddmzuqoxjctiezbxhxhhweirpgvpqfhrrsszigllllyszn","tbemuhmhvdbaggmbhqrcmaqyrbdtqhksrdagzpzyoqjeffecbj","bqnzwuqcvougpsdettehvpawougrnouvxvcvdlqqxvzkpabema","hbjpxhjzxfwvlovlkheapsvwputqkohyctmhuzxwxsvzycviqf","bvnffuobwirswtzeobudtosjclcgayeyzmhoeongnsxndjbiuo","jukhmgltrnialopwpwkwlornsilziukxzyoreftjyliuzycpsi","fdchbkogbqcamcgfjsygvaidivqymyjqacpdlbjyhsnmzjkdht","pnrfhpfzpklkuorpdkqikbwvqrlbjkevrenkgdrtevjlemzten","lyrpmnjbwvrqnufitvofejjlhdorgckxfvzruusdzzynnygsep","iiugqjvmuxoqofvdfxgzbypcmcuotrjajtmckhifiucqryxdjz","xvwfnzmvbmtaeglfpclxbwmwdzplpiktzvvtwrlvhotmlahtwn","iplodkacptgvaxsvmbpfiwgufdrjaxaarsdnneixmvubrpmwgx","eykynmjtpthxzwytmxkswqbcxqcfgcxbpmebvukenxgnwdxkvh","fljalvwokyurpefprjfdgtkddfnjsvhfrgqmmqzrzelccwafyf","cqsquwmvlvzfvztuahpmrezlfdktacqdlbgqvkswxvinegowpu"}))
}

func similarPairs(words []string) int {

	//wordsToInt := make([]int, len(words))
	wordsIntMap := make(map[int]int)
	for _, word := range words {
		temp := 0
		for _, b := range word {
			temp = temp | 1<<uint(b-'a')
		}
		wordsIntMap[temp]++
	}

	fmt.Println(wordsIntMap)

	result := 0
	for _, v := range wordsIntMap {
		if v <= 1 {
			continue
		}
		result += v * (v - 1) / 2
	}

	return result
}
