import random
import time

def manual_uuid():
    return "%08x-%04x-%04x-%04x-%012x" % (
        random.getrandbits(32),
        random.getrandbits(16),
        (random.getrandbits(16) & 0x0fff) | 0x4000,
        (random.getrandbits(16) & 0x3fff) | 0x8000,
        random.getrandbits(48),
    )

start = time.time()
for _ in range(1_000_000):
    manual_uuid()
end = time.time()
print(f"Python: {end - start:.5f} seconds")

