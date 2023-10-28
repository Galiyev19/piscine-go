package main

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

type NodeI struct {
	Data int
	Next *NodeI
}

func main() {
	// fmt.Println(piscine.AtoiBase("10", "0123456789ABCDEF"))
	// HACKATHON
	// ROT14
	// result := piscine.Rot14("abcdefghijklmnopqrstuvwxyz")

	// DescendComb
	// piscine.DescendComb()

	// UNMATCH
	// a := []int{1, 2, 3, 1, 2, 3, 4, 4, 5, 5}
	// b := []int{1, 1, 1, 1}
	// unmatch := piscine.Unmatch(a)
	// unmatch1 := piscine.Unmatch(b)
	// fmt.Println(unmatch)
	// fmt.Println(unmatch1)

	// FOODDELIVERYTIME
	// fmt.Println(piscine.FoodDeliveryTime("burger"))
	// fmt.Println(piscine.FoodDeliveryTime("chips"))
	// fmt.Println(piscine.FoodDeliveryTime("nuggets"))
	// fmt.Println(piscine.FoodDeliveryTime("burger") + piscine.FoodDeliveryTime("chips") + piscine.FoodDeliveryTime("nuggets"))

	// collatzcountdown
	// steps := piscine.CollatzCountdown(12)
	// fmt.Println(steps)

	// ABORT
	// middle := piscine.Abort(2, 3, 8, 5, 7)
	// fmt.Println(middle)
	// QUEST - 4

	// shoppingsummarycounter
	// summary := " "
	// for index, element := range piscine.ShoppingSummaryCounter(summary) {
	// 	fmt.Println(index, "=>", element)
	// }

	// Compact
	// const N = 6
	// a := make([]string, N)
	// a[0] = "a"
	// a[2] = "b"
	// a[4] = "c"

	// DEALPACKOFCARDS
	// deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	// piscine.DealAPackOfCards(deck)

	// JUMPOVER
	// fmt.Print(piscine.JumpOver("1010101010"))
	// fmt.Print(piscine.JumpOver(""))
	// fmt.Print(piscine.JumpOver("t w e l v e"))
	// fmt.Print(piscine.JumpOver("12"))

	// STRING SLICE
	// fmt.Println(piscine.StringToIntSlice("A quick brown fox jumps over the lazy dog"))
	// fmt.Println(piscine.StringToIntSlice("Converted this string into an int"))
	// fmt.Println(piscine.StringToIntSlice("hello THERE"))

	// JOIN
	// toConcat := []string{"Hello!", " How", " are", " you?"}
	// fmt.Println(piscine.Join(toConcat, ":"))

	// ENIGMA
	// x := 5
	// y := &x
	// z := &y
	// a := &z

	// w := 2
	// b := &w

	// u := 7
	// e := &u
	// f := &e
	// g := &f
	// h := &g
	// i := &h
	// j := &i
	// c := &j

	// k := 6
	// l := &k
	// m := &l
	// n := &m
	// d := &n

	// fmt.Println(***a)
	// fmt.Println(*b)
	// fmt.Println(*******c)
	// fmt.Println(****d)

	// piscine.Enigma(a, b, c, d)

	// fmt.Println("After using Enigma")
	// fmt.Println(***a)
	// fmt.Println(*b)
	// fmt.Println(*******c)
	// fmt.Println(****d)

	// DESCENDAPPENDRANGE
	// fmt.Println(piscine.DescendAppendRange(10, 5))
	// fmt.Println(piscine.DescendAppendRange(5, 10))
	// fmt.Println(piscine.DescendAppendRange(0, 1))

	// SHOPPINGCARTLISTSORT
	// slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
	// fmt.Println(piscine.ShoppingListSort(slice))

	// REVERSEMENUINDEX
	// fmt.Println(piscine.ReverseMenuIndex([]string{"starters", "drinks", "mains", "desserts"}))

	// COMCHECK
	// args := os.Args[1:]
	// res := piscine.ComCheck(args)
	// fmt.Println(res)

	// PODIUMPOSITION
	// p4 := []string{"4th Place"}
	// p3 := []string{"3rd Place"}
	// p2 := []string{"2nd Place"}
	// p1 := []string{"1st Place"}

	// position := [][]string{p4, p3, p2, p1}
	// fmt.Println(piscine.PodiumPosition(position))

	// ACTIVE BITS
	// decimal := 7
	// var binary string

	// fmt.Println(max)

	// for decimal > 0 {
	// 	remainder := decimal % 2
	// 	binary = string('0'+remainder) + binary
	// 	decimal /= 2
	// }
	// count := 0
	// for _, ch := range binary {
	// 	if ch == '1' {
	// 		count++
	// 	}
	// }
	// fmt.Println(count)

	// MAX
	// a := []int{23, 123, 1, 11, 55, 93}
	// max := a[0]

	// for _, ch := range a {
	// 	if max < ch {
	// 		max = ch
	// 	}
	// }

	// ROCKANDROLL
	// fmt.Println(piscine.RockAndRoll(4))
	// fmt.Println(piscine.RockAndRoll(9))
	// fmt.Println(piscine.RockAndRoll(6))

	// LOADOFBREAD
	// fmt.Print(piscine.LoafOfBread("deliciousbread"))
	// fmt.Print(piscine.LoafOfBread("This is a loaf of bread"))
	// fmt.Print(piscine.LoafOfBread("loaf"))
	// for _, v := range a {
	// 	fmt.Println(v)
	// }

	// fmt.Println("Size after compacting:", piscine.Compact(&a))

	// for _, v := range a {
	// 	fmt.Println(v)
	// }

	// piscine.ShoppingSummaryCounter(" ")

	// 1
	// arg := 12
	// fmt.Println(piscine.IterativeFactorial(arg))

	// 2
	// arg := 36
	// fmt.Println(piscine.RecursiveFactorial(arg))

	// 3
	// fmt.Println(piscine.IterativePower(4, 3))

	// 4
	// fmt.Println(piscine.RecursivePower(4, 3))

	// 5
	// arg1 := 2
	// fmt.Println(piscine.Fibonacci(arg1))

	// 6
	// fmt.Println(piscine.Sqrt(5))

	// 7
	// fmt.Println(piscine.IsPrime(90))
	// fmt.Println(piscine.IsPrime(97))

	// 8
	// fmt.Println(piscine.FindNextPrime(2))
	// fmt.Println(piscine.FindNextPrime(90))

	// 9
	// piscine.EightQueens()

	// QUEST - 5

	// 1
	// z01.PrintRune(piscine.FirstRune("♥01"))
	// z01.PrintRune(piscine.FirstRune("Salut!"))
	// z01.PrintRune(piscine.FirstRune("Ola!"))
	// z01.PrintRune('\n')

	// 2
	// z01.PrintRune(piscine.FirstRune("♥01"))
	// z01.PrintRune(piscine.LastRune("Hello!"))
	// z01.PrintRune(piscine.LastRune("Salut!"))
	// z01.PrintRune(piscine.LastRune("Ola!"))
	// z01.PrintRune('\n')

	// 3
	// z01.PrintRune(piscine.NRune("Hello!", 3))
	// z01.PrintRune(piscine.NRune("Salut!", 2))
	// z01.PrintRune(piscine.NRune("Bye!", -1))
	// z01.PrintRune(piscine.NRune("Bye!", 5))
	// z01.PrintRune(piscine.NRune("Ola!", 4))
	// z01.PrintRune('\n')

	// 4
	// fmt.Println(piscine.Compare("Hello!", "Hello!"))
	// fmt.Println(piscine.Compare("Salut!", "lut!"))
	// fmt.Println(piscine.Compare("Ola!", "Ol"))

	// 5
	// s := "Hello 78 World!    4455 /"
	// nb := piscine.AlphaCount(s)
	// fmt.Println(nb)

	// 6
	// fmt.Println(piscine.Index("Hello!", "l"))
	// fmt.Println(piscine.Index("Salut!", "alu"))
	// fmt.Println(piscine.Index("Ola!", "hOl"))

	// 7
	// fmt.Println(piscine.Concat("Hello!", " How are you?"))

	// 8
	// fmt.Println(piscine.IsUpper("HELLO"))
	// fmt.Println(piscine.IsUpper("HELLO!"))

	// 9
	// fmt.Println(piscine.IsLower("hello"))
	// fmt.Println(piscine.IsLower("hello!"))

	// 10
	// fmt.Println(piscine.IsAlpha("Hello! How are you?"))
	// fmt.Println(piscine.IsAlpha("HelloHowareyou"))
	// fmt.Println(piscine.IsAlpha("What's this 4?"))
	// fmt.Println(piscine.IsAlpha("Whatsthis4"))

	// 11

	// fmt.Println(piscine.IsNumeric("010203"))
	// fmt.Println(piscine.IsNumeric("01,02,03"))

	// 12
	// fmt.Println(piscine.IsPrintable("Hello"))
	// fmt.Println(piscine.IsPrintable("Hello\n"))

	// 13
	// fmt.Println(piscine.ToUpper("1RQRT8-1nM}Em"))

	// 14
	// fmt.Println(piscine.ToLower("Hello! How are you?"))

	// 15

	// piscine.PrintNbrInOrder(321)
	// piscine.PrintNbrInOrder(0)
	// piscine.PrintNbrInOrder(321)
	// z01.PrintRune('\n')

	// 16
	// fmt.Println(piscine.TrimAtoi("12345"))
	// fmt.Println(piscine.TrimAtoi("str123ing45"))
	// fmt.Println(piscine.TrimAtoi("012 345"))
	// fmt.Println(piscine.TrimAtoi("Hello World!"))
	// fmt.Println(piscine.TrimAtoi("sd+x1fa2W3s4"))
	// fmt.Println(piscine.TrimAtoi("sd-x1fa2W3s4"))
	// fmt.Println(piscine.TrimAtoi("sdx1-fa2W3s4"))
	// fmt.Println(piscine.TrimAtoi("sdx1+fa2W3s4"))

	// 17
	// fmt.Println(piscine.Capitalize("7E^68_Vrm2BIT"))

	// 18

	// elems := []string{"Hello!", " How", " are", " you?"}
	// fmt.Println(piscine.BasicJoin(elems))

	// 19

	// toConcat := []string{"Hello!", " How", " are", " you?"}
	// fmt.Println(piscine.Join(toConcat, ":"))

	// 20

	// piscine.PrintNbrBase(125, "0123456789")
	// z01.PrintRune('\n')
	// piscine.PrintNbrBase(-125, "01")
	// z01.PrintRune('\n')
	// piscine.PrintNbrBase(125, "0123456789ABCDEF")
	// z01.PrintRune('\n')
	// piscine.PrintNbrBase(-125, "choumi")
	// z01.PrintRune('\n')
	// piscine.PrintNbrBase(-9223372036854775808, "0123456789")
	// z01.PrintRune('\n')

	// QUEST - 7

	// 1
	// fmt.Println(piscine.AppendRange(10, 20))
	// fmt.Println(piscine.AppendRange(10, 5))
	// fmt.Println(piscine.AppendRange(0, 1))

	// 2
	// fmt.Println(piscine.MakeRange(5, 10))
	// fmt.Println(piscine.MakeRange(10, 5))
	// fmt.Println(piscine.MakeRange(0, 1))
	// fmt.Println(piscine.MakeRange(-1646928, -1646911))

	// 3
	// test := []string{"Hello", "how", "are", "you?"}
	// fmt.Println(piscine.ConcatParams(test))

	// 4
	// fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("S>/2\\`Yx)'/Ax 93uU4@)nA4[/  ::L}J./fBO2S0 G`B405HGRRg4\\ #>3<2@D2EH3B1 d$Kz-rUL75%GJ pg%uD!OTMU{{K riuG1@9:UKO=4"))

	// 5
	// a := piscine.SplitWhiteSpaces("Hello how are you?")
	// piscine.PrintWordsTables(a)

	// 6
	// s := "HelloHAhowHAareHAyou?HA"
	// fmt.Printf("%#v\n", piscine.Split(s, "HA"))

	// 7
	// fmt.Println(piscine.AtoiBase("125", "0123456789"))
	// fmt.Println(piscine.AtoiBase("1111101", "01"))
	// fmt.Println(piscine.AtoiBase("7D", "0123456789ABCDEF"))
	// fmt.Println(piscine.AtoiBase("uoi", "choumi"))
	// fmt.Println(piscine.AtoiBase("bbbbbab", "-ab"))

	// QUEST-9

	// FOREACH
	// a := []int{1, 2, 3, 4, 5, 6}
	// piscine.ForEach(piscine.PrintNbr, a)
	// fmt.Println()

	// MAP
	// a := []int{1, 2, 3, 4, 5, 6}
	// result := piscine.Map(piscine.IsPrime, a)
	// fmt.Println(result)

	// ANY
	// a1 := []string{"Hello", "how", "are", "you"}
	// a2 := []string{"This", "is", "4", "you"}

	// result1 := piscine.Any(piscine.IsNumeric, a1)
	// result2 := piscine.Any(piscine.IsNumeric, a2)

	// fmt.Println(result1)
	// fmt.Println(result2)

	// COUNTIF
	// tab1 := []string{"Hello", "how", "are", "you"}
	// tab2 := []string{"This", "1", "is", "4", "you"}
	// tab3 := []string{"8619050580303", "8416927713951", "4256330544964", "9355794201834", "1938399807298", "8678471957948", "0391243979670", "0320555340756"}
	// answer1 := piscine.CountIf(piscine.IsNumeric, tab1)
	// answer2 := piscine.CountIf(piscine.IsNumeric, tab2)
	// answer3 := piscine.CountIf(piscine.IsNumeric, tab3)
	// fmt.Println(answer1)
	// fmt.Println(answer2)
	// fmt.Println(answer3)

	// ISSORTED
	// a1 := []int{0, 1, 2, 3, 4, 5}
	// a2 := []int{0, 2, 1, 3, 4, 5}

	// a3 := []int{298148, 12862, 272197, 857472, 850973, 530321, -922198, -241980}

	// result1 := piscine.IsSorted(MyFunc, a1)
	// result2 := piscine.IsSorted(MyFunc, a2)
	// result3 := piscine.IsSorted(MyFunc, a3)

	// fmt.Println(result1)
	// fmt.Println(result2)
	// fmt.Println(result3)

	// sortwordarr
	// result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	// piscine.SortWordArr(result)

	// fmt.Println(result)

	// advancedsortwordarr
	// result := []string{"z", "2", "1", "b", "B", "2", "c", "C", "3"}
	// piscine.AdvancedSortWordArr(result, piscine.Compare)

	// fmt.Println(result)
	// result := piscine.ConvertBase("hWI?h?oA", "WhoAmI?", "choumi")
	// fmt.Println(result)

	// piscine.EightQueens()
	// QUEST 11

	// link := &piscine.List{}

	// piscine.ListPushBack(link, "1")
	// piscine.ListPushBack(link, "2")
	// piscine.ListPushBack(link, "3")
	// piscine.ListPushBack(link, "5")

	// piscine.ListForEach(link, piscine.Add2_node)

	// it := link.Head
	// for it != nil {
	// 	fmt.Println(it.Data)
	// 	it = it.Next
	// }

	// LISTFOREACH

	// link := &piscine.List{}

	// piscine.ListPushBack(link, 1)
	// piscine.ListPushBack(link, "hello")
	// piscine.ListPushBack(link, 3)
	// piscine.ListPushBack(link, "there")
	// piscine.ListPushBack(link, 23)
	// piscine.ListPushBack(link, "!")
	// piscine.ListPushBack(link, 54)

	// PrintList(link)

	// fmt.Println("--------function applied--------")
	// piscine.ListForEachIf(link, PrintElem, piscine.IsPositiveNode)

	// piscine.ListForEachIf(link, StringToInt, piscine.IsAlNode)

	// fmt.Println("--------function applied--------")
	// PrintList(link)

	// fmt.Println()

	// LISTFIND
	// link := &piscine.List{}

	// piscine.ListPushBack(link, "hello")
	// piscine.ListPushBack(link, "hello1")
	// piscine.ListPushBack(link, "hello2")
	// piscine.ListPushBack(link, "hello3")

	// found := piscine.ListFind(link, interface{}("hello2"), piscine.CompStr)

	// fmt.Println(found)
	// fmt.Println(*found)

	// LISTREMOVEIF
	// link := &piscine.List{}
	// link2 := &piscine.List{}

	// fmt.Println("----normal state----")
	// piscine.ListPushBack(link2, 1)
	// PrintList(link2)
	// piscine.ListRemoveIf(link2, 1)
	// fmt.Println("------answer-----")
	// PrintList(link2)
	// fmt.Println()

	// fmt.Println("----normal state----")
	// piscine.ListPushBack(link, 1)
	// piscine.ListPushBack(link, "Hello")
	// piscine.ListPushBack(link, 1)
	// piscine.ListPushBack(link, "There")
	// piscine.ListPushBack(link, 1)
	// piscine.ListPushBack(link, 1)
	// piscine.ListPushBack(link, "How")
	// piscine.ListPushBack(link, 1)
	// piscine.ListPushBack(link, "are")
	// piscine.ListPushBack(link, "you")
	// piscine.ListPushBack(link, 1)
	// PrintList(link)

	// piscine.ListRemoveIf(link, 1)
	// fmt.Println("------answer-----")
	// PrintList(link)

	// LISTMERGE
	// link := &piscine.List{}
	// link2 := &piscine.List{}

	// piscine.ListPushBack(link, "a")
	// piscine.ListPushBack(link, "b")
	// piscine.ListPushBack(link, "c")
	// piscine.ListPushBack(link, "d")
	// fmt.Println("-----first List------")
	// PrintList(link)

	// piscine.ListPushBack(link2, "e")
	// piscine.ListPushBack(link2, "f")
	// piscine.ListPushBack(link2, "g")
	// piscine.ListPushBack(link2, "h")
	// fmt.Println("-----second List------")
	// PrintList(link2)

	// fmt.Println("-----Merged List-----")
	// piscine.ListMerge(link, link2)
	// PrintList(link)

	// LISTSORT
	// var link *piscine.NodeI

	// link = listPushBack(link, 5)
	// link = listPushBack(link, 4)
	// link = listPushBack(link, 3)
	// link = listPushBack(link, 2)
	// link = listPushBack(link, 1)

	// PrintList(piscine.ListSort(link))

	// LISTSORTINSERT
	// var link *piscine.NodeI

	// link = listPushBack(link, 1)
	// link = listPushBack(link, 4)
	// link = listPushBack(link, 9)

	// PrintList(link)

	// link = piscine.SortListInsert(link, -2)
	// link = piscine.SortListInsert(link, 2)
	// PrintList(link)

	// sortedlistmerge
	// var link *piscine.NodeI
	// var link2 *piscine.NodeI

	// link = listPushBack(link, 3)
	// link = listPushBack(link, 5)
	// link = listPushBack(link, 7)

	// link2 = listPushBack(link2, -2)
	// link2 = listPushBack(link2, 9)

	// PrintList(piscine.SortedListMerge(link2, link))

	// root := &piscine.TreeNode{Data: "4"}
	// piscine.BTreeInsertData(root, "1")
	// piscine.BTreeInsertData(root, "7")
	// piscine.BTreeInsertData(root, "5")
	// fmt.Println(piscine.BTreeIsBinary(root))

	// fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))

	// root := &piscine.TreeNode{Data: "4"}
	// piscine.BTreeInsertData(root, "1")
	// piscine.BTreeInsertData(root, "7")
	// piscine.BTreeInsertData(root, "5")
	// piscine.BTreeApplyByLevel(root, fmt.Println)
	// fmt.Println(decimalToHexadecimal(10))
	// root := &piscine.TreeNode{Data: "4"}
	// piscine.BTreeInsertData(root, "1")
	// piscine.BTreeInsertData(root, "7")
	// piscine.BTreeInsertData(root, "5")
	// node := piscine.BTreeSearchItem(root, "4")
	// fmt.Println("Before delete:")
	// piscine.BTreeApplyInorder(root, fmt.Println)
	// root = piscine.BTreeDeleteNode(root, node)
	// fmt.Println("After delete:")
	// piscine.BTreeApplyInorder(root, fmt.Println)
}

func decimalToHexadecimal(decimal int) string {
	hexDigits := "0123456789abcdef"
	var result string

	for decimal > 0 {
		remainder := decimal % 16
		result = string(hexDigits[remainder]) + result
		decimal /= 16
	}

	return result
}

func Capitalize(s string) string {
	result := make([]rune, len(s))
	found := true
	for i, ch := range s {
		if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' {
			if found {
				result[i] = ch - 'a' - 'A'
				found = false
			} else {
				result[i] = ch + 'a' - 'A'
			}
		} else {
			result[i] = ch
			found = true
		}
	}

	return string(result)
}

// func PrintList(l *piscine.NodeI) {
// 	it := l
// 	for it != nil {
// 		fmt.Print(it.Data, " -> ")
// 		it = it.Next
// 	}
// 	fmt.Print(nil, "\n")
// }

// func listPushBack(l *piscine.NodeI, data int) *piscine.NodeI {
// 	n := &piscine.NodeI{Data: data}

// 	if l == nil {
// 		return n
// 	}
// 	iterator := l
// 	for iterator.Next != nil {
// 		iterator = iterator.Next
// 	}
// 	iterator.Next = n
// 	return l
// }

// func PrintElem(node *piscine.NodeL) {
// 	fmt.Println(node.Data)
// }

// func StringToInt(node *piscine.NodeL) {
// 	node.Data = 2
// }

// func MyFunc(a, b int) int {
// 	if a > b {
// 		return 1
// 	} else if a == b {
// 		return 0
// 	} else {
// 		return -1
// 	}
// }
