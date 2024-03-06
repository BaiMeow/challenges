#include<stdio.h>

int main(){
    FILE *file;
    file = fopen("/flag","r");
    char flag[64];
    fscanf(file,"%s",flag);
    printf("%s",flag);
}