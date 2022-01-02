import math

def run():
  number = int(input("Enter a starting number: "))

  if number < 1:
    return print("Number must be greater than zero.")

  i = 1

  while number != 1:
    if (number % 2) == 0:
      number /= 2
    else:
      number = (number * 3) + 1

    i += 1
    
    print(f"Step [{i}] {math.trunc(number)}")

run()