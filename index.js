const readline = require("readline").createInterface({
  input: process.stdin,
  output: process.stdout,
});

readline.question(
  "Enter a starting number: ",
  (number) => {
    if (number < 1) {
      return console.log("Number must be greater than zero.")
    }

    let i = 1;

    while (number !== 1) {
      if ((number % 2) == 0) {
        number /= 2
      } else {
        number = (number * 3) + 1
      }

      console.log(`[Step ${i++}] ${number}`);
    }

    readline.close();
  }
);