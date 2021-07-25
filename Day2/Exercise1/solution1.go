package main

import "fmt"

func frequency(str string,result map[string]int) int{

	for i:=0;i<len(str);i++ {
		char:= str[i:i+1]
		result[char]+=1
	}
	return 1
}
func worker(jobs <-chan string, done chan<- int,result map[string]int){
	for str:= range jobs {
		done<- frequency(str,result)
	}
}
func main(){

	set:= [] string{"quick","brown","fox","lazy","dog"}
	//fmt.Println(frequency(set[1]))
	jobs:=make(chan string,len(set))
	done:=make(chan int,len(set))
	result:=make(map[string]int)
	go worker(jobs,done,result)
    for _,str:= range set{
    	 jobs<-str
	 }
	 close(jobs)

    for i:=0;i<len(set);i++{
    	<-done
	}
   fmt.Println(result)
}
