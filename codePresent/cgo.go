// +build OMIT
package main // OMIT START

/*
#include <stdio.h>

void pyramid(int p) {
    for(int i = 0 ; i <= p ; i++ ) {
        for(int j = 1 ; j <= i ; j++ ) {
        printf("*");
        }
    printf("\n");
    }
}
*/
import "C" // 注意：此語句上方不能有其他空白列

func main() {
	for i := 1; i <= 3; i++ {
		C.pyramid((C.int)(i)) // OMIT END
	}
}
