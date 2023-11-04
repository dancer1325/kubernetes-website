import time

def process():
    n = 0
    while True:
        print(f'Iteration {n}')
        n += 1
        time.sleep(5)

if __name__ == "__main__":
    process()