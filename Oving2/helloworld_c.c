
// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld_c.c -lpthread
#include <pthread.h>
#include <stdio.h>
#include <stdint.h>

	pthread_mutex_t lock;
	pthread_mutex_t lock2;

void adder();
void adder2();


int i = 0;




int main() {


	pthread_mutex_init(&lock, NULL);
	pthread_mutex_init(&lock2, NULL);

	pthread_t adder_thr;
	pthread_t subst_thr;

	pthread_create(&adder_thr, NULL, adder2, NULL );
	pthread_create(&subst_thr, NULL, adder, NULL );

	pthread_join(adder_thr, NULL);
	pthread_join(subst_thr, NULL);
	pthread_mutex_destroy(&lock2);
	pthread_mutex_destroy(&lock);
	printf("%i\n", i);

	return 0;
}


void adder() {
	for(int x = 0; x < 1000000; x++){

		pthread_mutex_lock(&lock);
		i--;
		pthread_mutex_unlock(&lock);

	}
	return NULL;
}

void adder2() {
	for(int x = 0; x < 1000000; x++){

		pthread_mutex_lock(&lock2);
		i++;
		pthread_mutex_unlock(&lock2);
		
	}
	return NULL;
}
