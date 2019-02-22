package model

/**
* Created by wanjx in 2019/2/21 22:32
**/

type Article struct {
	Uid string
	Title string
	Date string
	Content string
	Author User
}
