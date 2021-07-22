package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
func PerformanceRating(wait *sync.WaitGroup,rating *int){
	defer wait.Done()
	*rating+= rand.Intn(100)
}
func GetAverage(rating int,num int){
	fmt.Println("Average rating of the class is ",rating/num)
}
func main(){
	var wait sync.WaitGroup
	var rating int =0
	var NoOfStudents int =200
	for student:=1;student<=NoOfStudents;student++{
		wait.Add(1)
		go PerformanceRating(&wait,&rating)
	}
	time.Sleep(time.Second)
	wait.Wait()
	GetAverage(rating,NoOfStudents)
}