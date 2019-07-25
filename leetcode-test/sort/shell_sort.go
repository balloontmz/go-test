//Package sort 希尔排序
/*
希尔排序
时间复杂度跟增量的选取有关O(n^（1.3—2）)
希尔排序是插入排序的改进版本，其核心思想是将原数据集合分隔成若干个子序列，然后再对子序列分别进行直接插入排序，使子序列基本有序，
再对全体记录进行一次直接插入排序。
最关键的是跳跃和分隔策略，间隔多大的问题。通常将相聚某个增量的记录组成一个子序列，这样直接插入排序得到的结果是基本有序的。
本例通过 i = int(i/3)+1 来确定增量的值--此处替换为地板除
*/
package sort

//ShellSort 希尔排序
func ShellSort (li []int) {
	l := len(li)
	for increment := len(li); increment > 1;{
		// 此处只能前置处理，暂时没找到后置处理的办法。因为最后一次遍历必然 increment 要为 1
		increment = increment / 3 + 1
		// i 从第二个数字开始，如果第二个小于第一个，就互换
    	// 边界条件是一个需要关注的大点
		for i := increment; i < l; i++ {
			if li[i] < li[i-increment] {
				temp := li[i]
				j := i - increment 
				for j >=0 {
					if temp >= li[j] {
						break
					}
					li[j+increment] = li[j]
					j = j - increment
				}
				li[j+increment] = temp
			}
		}
	}
}