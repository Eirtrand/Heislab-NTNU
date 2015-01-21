// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld_c.c -lpthread
#include <pthread.h>
#include <stdio.h>
// Note the return type: void*
int i= 0;
pthread_mutex_t lock;


void* someThreadFunction1(){
for(int k = 0;k<1000000;k++){
	pthread_mutex_lock(&lock);
	i++;
	pthread_mutex_unlock(&lock);
}
return NULL;
}

void* someThreadFunction2(){
for(int k = 0;k<1000000;k++){
	pthread_mutex_lock(&lock);
	i--;
	pthread_mutex_unlock(&lock);
}
return NULL;
}


int main(){
	if (pthread_mutex_init(&lock, NULL) != 0)
    {
        printf("\n mutex init failed\n");
        return 1;
    }

	pthread_t Thread1;
	pthread_t Thread2;

	pthread_create(&Thread1, NULL, someThreadFunction1, &lock);
	pthread_create(&Thread2, NULL, someThreadFunction2, &lock);
	// Arguments to a thread would be passed here ---------^
	pthread_join(Thread1, NULL);
	pthread_join(Thread2, NULL);

	pthread_mutex_destroy(&lock);

	printf("%d\n",i);
return 0;
}




