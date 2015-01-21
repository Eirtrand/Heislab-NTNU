// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld_c.c -lpthread
#include <pthread.h>
#include <stdio.h>
// Note the return type: void*
int i= 0;

void* someThreadFunction1(){
int j;
for(j = 0;j<1000000;j++)
	i++;
return NULL;
}

void* someThreadFunction2(){
int k;
for(k = 0;k<1000000;k++)
	i--;
return NULL;
}


int main(){

pthread_t Thread1;
pthread_t Thread2;

pthread_create(&Thread1, NULL, someThreadFunction1, NULL);
pthread_create(&Thread2, NULL, someThreadFunction2, NULL);
// Arguments to a thread would be passed here ---------^
pthread_join(Thread1, NULL);
pthread_join(Thread2, NULL);
printf("%d",i);
return 0;
}




