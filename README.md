# AliCloudDisk OF Golang
AliCloudDisk is a Go language written Ali cloud disk interface library,Due to my technical problems, I will keep learning and updating it


## Basic function
> Obtain user login information
>
> Gets a list of user files
## Unfinished features
> Dynamic acquisition Authorization
>
> Download file
>
> Upload file
>
> Get share file list and share link
>
> ............

# Getting Start
## Getting AliCloudDisk
With Go module support, simply add the following import

```shell  
  import "github.com/XiaoTuanZiJiang/AliCloudDisk"
 ```
to your code, and then will automatically fetch the necessary dependencies.```go [build|run|test]```

Otherwise, run the following Go command to install the package:```AliCloudDisk```

```shell  
  $ go get -u github.com/XiaoTuanZiJiang/AliCloudDisk 
 ```

## Running AliCloudDisk
First you need to import Gin package for using AliCloudDisk, one simplest example likes the follow :example.go
```Golang
package main

import (
	"fmt"
	AliCloudDisk "github.com/XiaoTuanZiJiang/AliCloudDisk/apis"
)

func main() {
	c := AliCloudDisk.NewCloudDiskConnection("Fill in your Authorization",&AliCloudDisk.Config{})
	fmt.Println(c.UserInfo)
}
```
And use the Go command to run the demo:
```shell  
  $ go run example.go
 ```
### Authorization acquisition mode
Open Ali cloud disk and log in, open the developer mode F12, click the network and refresh, and open an item at any time to find it
![无标题](https://user-images.githubusercontent.com/85481419/236196610-12403a4f-9715-423e-a5e4-708773c4d1ff.png)

