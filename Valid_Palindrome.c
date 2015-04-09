// 125 Valid Palindrome Easy
#include <string.h>
#include <stdlib.h>
#define bool int
#define true 1
#define false 0

bool isAlphanumeric(char c) {
    if(c>='A'&&c<='Z' || c>='a'&&c<='z' || c>='0'&&c<='9')
        return true;
    else
        return false;
}

bool isEqualIgnoreCase(char a, char b) {
    if(a>='0'&&a<'9'){
        return a==b?true:false;
    }else {
        if(abs(a-b)==0 || abs(a-b)==abs('Z'-'z'))
            return true;
        else
            return false;
    }
}

bool isPalindrome(char *s) {
    int len = strlen(s);
    if(len == 0)
        return true;
    char *start = s;
    char *end = s+len-1;
    while(!isAlphanumeric(*start) && start < end)
        start++;
    while(!isAlphanumeric(*end) && start < end)
        end--;
    while(start < end){
        if(!isEqualIgnoreCase(*start, *end)){
            return false;
        }
        start ++;
        end --;
        while(!isAlphanumeric(*start) && start < end)
            start++;
        while(!isAlphanumeric(*end) && start < end)
            end--;
    }
    return true;
}
