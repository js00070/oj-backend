package routine

import (
	"fmt"
	"oj/cache"
	"oj/model"
	"strconv"
	"strings"
	"time"
)

// Routine aaa
func Routine() {
	for {
		if cache.RedisClient.LLen("res_commit").Val() != 0 && cache.RedisClient.LLen("res_output").Val() == cache.RedisClient.LLen("res_commit").Val() {
			commit, err := cache.RedisClient.RPop("res_commit").Result()
			if err != nil {
				panic(err)
			}
			commitID, err := strconv.Atoi(commit)
			if err != nil {
				panic(err)
			}
			output, err := cache.RedisClient.RPop("res_output").Result()
			if err != nil {
				panic(err)
			}
			status := 1
			if strings.Contains(output, "Failures") {
				status = -1
			}
			fmt.Printf("set commit: %v %v\n %v\n", commitID, status, output)
			if err := model.SetCommit(commitID, output, status); err != nil {
				panic(err)
			}
		} else {
			time.Sleep(time.Second)
			continue
		}
	}
}
