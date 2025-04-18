import time

def loop_counter():
    start_time = time.time()
    count = 0

    while time.time() - start_time < 10:  # Loop for 60 seconds
        count += 1
        # print(f"Python: Looped {count}")

    print(f"Looped {count} times in 1 minute.")

loop_counter()