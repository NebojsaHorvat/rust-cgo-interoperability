// C Wrapper: wrapper.c

#include <stdio.h>
// #include "_cgo_export.h"


extern int getState(int);

int getStateWrapper(int index) {
    printf("Function getState from C called!\n");
    return getState(index);
}