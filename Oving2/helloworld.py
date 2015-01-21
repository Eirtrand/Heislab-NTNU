# Python 3.3.3 and 2.7.6
# python helloworld_python.py

from threading import Thread, Lock

i = 0
mutex = Lock()

def someThreadFunction1():
    global i
    for x in range(0,1000000):
        mutex.acquire()
        i += 1
        mutex.release()
        

def someThreadFunction2():
    global i
    for y in range(0,1000001):
        mutex.acquire()
        i -= 1
        mutex.release()


def main():
    someThread1 = Thread(target = someThreadFunction1, args = (),)
    someThread2 = Thread(target = someThreadFunction2, args = (),)
    someThread1.start()
    someThread2.start()
    
    someThread1.join()
    someThread2.join()
    print("Hello from main!")
    print(i)


main()
