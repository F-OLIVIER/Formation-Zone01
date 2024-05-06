const express = require('express')
const { execSync } = require('child_process')

const app = express()

app.use(express.json())
app.use(express.static('.'))

console.log("Average : http://localhost:3000/?guesser=average");
console.log("Big range : http://localhost:3000/?guesser=big-range");
console.log("Correlation coef : http://localhost:3000/?guesser=correlation-coef");
console.log("Huge range : http://localhost:3000/?guesser=huge-range");
console.log("Linear regr : http://localhost:3000/?guesser=linear-regr");
console.log("Median : http://localhost:3000/?guesser=median");
console.log("Mse : http://localhost:3000/?guesser=mse");
console.log("Nic : http://localhost:3000/?guesser=nic");

//Post requests from front-end, requesting the results of the programs
app.post('/', async (req, res) => {
  let { guesser, data_set } = req.body

  //Figuring it if it is the student or the ai program to test
  let test_file = !guesser.startsWith('./') ? './ai/' + guesser : guesser

  let values_to_test = await data_set.reduce((acc, add) => acc + add + '\n', '')

  //executes the program with the numbers as stdin entries and saves its ouput in `output`
  let output
  try {
    output = execSync(test_file, {
      input: values_to_test,
    })
  } catch (err) {
    res.send({ status: 500, msg: 'No file found: ' + guesser })
    throw Error('No file found\n' + err)
  }

  let value = output.toString().split('\n')
  value.pop()

  let result = 0
  let returning = []
  let correct = 0

  //This loop will iterate through each number and append to the `returning` array an instance of:
  //- the result at the moment
  //- the range given by the program
  //- the value given by the `data_set`
  //- the color depending if the program got the range right or not

  for (let i = 0; i < data_set.length; i++) {
    if (value[i - 1]) {
      let [bet_low, bet_high] = value[i - 1]
        .split(' ')
        .map((nbr) => parseInt(nbr))
      let gotRight = data_set[i] >= bet_low && data_set[i] <= bet_high

      result += gotRight
        ? Math.round(
            10000000 / (1 + bet_high - bet_low) / (data_set.length - 1)
          )
        : 0

      correct += gotRight ? 1 : 0

      returning.push({
        result: result,
        range: [bet_low, bet_high],
        value: data_set[i],
        color: gotRight ? 'green' : 'red',
        correct: correct,
      })
    } else {
      returning.push({
        value: data_set[i],
        color: test_file.includes('ai') ? 'dodgerblue' : 'darkorchid',
      })
    }
  }

  res.send(JSON.stringify({ ranges: returning }))
})

app.listen(process.env.PORT || 3000)

console.log('Listening in port 3000')
