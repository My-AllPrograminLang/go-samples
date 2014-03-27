void voidfunc();
int intfunc(int);

// Built with:
//
//
// CFLAGS = -Wall -std=c99 -g -O0
// mylib: mylib.c mylib.h
//   gcc $(CFLAGS) -fpic -c mylib.c
//   gcc -shared -o libmylib.so mylib.o
//
// The implementation is:
//
//
//#include "mylib.h"
//#include <stdio.h>


//void voidfunc() 
//{
    //printf("mylib's voidfunc called\n");
//}


//int intfunc(int arg)
//{
    //int n = 1;
    //for (int i = 1; i <= arg; ++i) {
        //n *= i;
    //}
    //return n;
//}


