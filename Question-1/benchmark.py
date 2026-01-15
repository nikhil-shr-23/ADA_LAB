import time
import random
import copy
from sorting_algorithms import (
    insertion_sort,
    bubble_sort,
    merge_sort,
    quick_sort,
    counting_sort,
    radix_sort,
    bucket_sort
)

def generate_data(size):
    return [random.randint(0, 10000) for _ in range(size)]

def benchmark(algorithms, sizes):
    print(f"{'Algorithm':<20} | {'Size':<10} | {'Time (s)':<15} | {'Status':<10}")
    print("-" * 65)
    
    for size in sizes:
        data = generate_data(size)
        for name, func in algorithms.items():
            arr_copy = copy.deepcopy(data)
            start_time = time.time()
            if name == "Quick Sort":
                sorted_arr = func(arr_copy) 
            else:
                func(arr_copy) 
                sorted_arr = arr_copy
            
            end_time = time.time()
            elapsed_time = end_time - start_time
            
            is_sorted = all(sorted_arr[i] <= sorted_arr[i+1] for i in range(len(sorted_arr)-1))
            status = "Pass" if is_sorted else "Fail"
            
            print(f"{name:<20} | {size:<10} | {elapsed_time:<15.6f} | {status:<10}")
        print("-" * 65)

if __name__ == "__main__":
    algorithms = {
        "Insertion Sort": insertion_sort,
        "Bubble Sort": bubble_sort,
        "Merge Sort": merge_sort,
        "Quick Sort": quick_sort,
        "Counting Sort": counting_sort,
        "Radix Sort": radix_sort,
        "Bucket Sort": bucket_sort
    }
    
    sizes = [100, 1000, 5000]
    benchmark(algorithms, sizes)
