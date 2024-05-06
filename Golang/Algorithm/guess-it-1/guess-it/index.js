const genData = async (data_set) => {
  let random = Math.floor(Math.random() * 5) + 1
  return fetch(`./data_sets/${data_set}/${random}.txt`).then((file) =>
    file.text()
  )
}

// Get the guesser AI chosen in the URL
let { guesser = '' } = Object.fromEntries(new URLSearchParams(location.search))
document.querySelector('.guesser-name.blue-text').innerHTML = guesser

//draws graph with the points
const drawGraph = (data_set) => {
  let canvas = document.getElementById('graph')
  new Chart(canvas, {
    type: 'scatter',
    data: {
      labels: Array.from(data_set.keys()),
      datasets: [
        {
          label: 'Numbers Generated',
          data: data_set,
          fill: false,
          borderColor: 'rgb(75,192,192)',
          tension: 0.1,
        },
      ],
    },
  })
  document
    .querySelector('.options-wrapper')
    .appendChild(document.createElement('div'))
}

// Create buttons with numbers corresponding to which data set will be created
//(see gen-data.js to check out which data_set will be generated)
// The onclick method for each button, starts the displaying of the results of
//each program with the random data set according to the number
for (let i = 1; i <= 5; i++) {
  let btn = document.createElement('button')
  btn.class = `button-${i}`
  btn.innerHTML = 'Test Data ' + i
  btn.onclick = () => {
    document.querySelectorAll('.final-result').forEach((div) => {
      div.innerHTML = ''
    })
    if (guesser) {
      genData(i).then((data) => {
        data = data.split('\n')
        data.pop()

        fetchAndDisplay(data, guesser).then((_) => drawGraph(data))
      })
    } else {
      throw Error('The program needs one guesser to go against')
    }
  }
  document.querySelector('.buttons').appendChild(btn)
}

const numberWithSpaces = (x) =>
  x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ' ')

// Returns a Promise that contains a `p` element for every value and range
const showEveryEntry = ({ ranges }) =>
  new Promise((res, _) => {
    let entries = []
    for (const [i, values] of ranges.entries()) {
      if (i != 0) {
        let entry = document.createElement('p')
        entry.style = `color: ${values.color};`
        entry.innerText = `\n${values.range[0]} <> ${
          values.range[1]
        }\nResult= ${numberWithSpaces(values.result)}\nValue #${i + 1}: ${
          values.value
        } \n`
        entries.push(entry)
      } else {
        let entry = document.createElement('p')
        entry.style = `color: ${values.color};`
        entry.innerText = `Value #${i + 1}: ${values.value} \n`
        entries.push(entry)
      }
    }
    return res(entries)
  })

const loadResult = (result, corrPerc, guesserNbr) => {
  document.getElementById(
    'final-result-' + guesserNbr
  ).innerText = `Final Result: ${numberWithSpaces(
    result
  )}\n Correct Guesses: ${numberWithSpaces(corrPerc)}%`
}

// Makes a post request to the backend in order to obtain the values needed for display
const fetching = (guesserName, data) =>
  fetch('/', {
    headers: {
      Accept: 'application/json',
      'Content-Type': 'application/json',
    },
    method: 'POST',
    body: JSON.stringify({
      guesser: guesserName,
      data_set: data,
    }),
  })

const workerHandler = (ranges, guesserNbr) => {
  // Connection to the worker in order to display the p elements one after the other
  const guesser = new Worker('./lib/test.js')

  guesser.postMessage(ranges)
  guesser.onmessage = ({ data }) => {
    let entry = document.createElement('p')
    if (typeof data[0] == 'string') {
      // Means is the last one and updates the final result
      document.getElementById(
        'final-result-' + guesserNbr
      ).innerText = `Result:${data.split(':')[1]}`
    } else if (data[2]) {
      entry.style = `color:${data[0].color}`
      entry.innerHTML = `Value #${data[1] + 1}: ${data[0].value}\n`
      document.getElementById('guesser-' + guesserNbr).appendChild(entry)
    } else {
      entry.style = `color:${data[0].color}`
      entry.innerHTML = `\n${data[0].range[0]} <> ${
        data[0].range[1]
      }\nResult= ${numberWithSpaces(data[0].result)}\nValue #${data[1] + 1}: ${
        data[0].value
      }\n`
      document.getElementById('guesser-' + guesserNbr).appendChild(entry)
    }
  }

  return guesser
}

// Scrolls to bottom automatically
const scrollBottom = (guesserScreen) => {
  const isScrolledToBottom =
    guesserScreen.scrollHeight - guesserScreen.clientHeight <=
    guesserScreen.scrollHeight + 1

  if (isScrolledToBottom) {
    guesserScreen.scrollTop =
      guesserScreen.scrollHeight - guesserScreen.clientHeight
  }
}

const fetchAndDisplay = async (data, aiProg) => {
  let studProg = './student/script.sh'

  let guesserStud = document.getElementById('guesser-1'),
    guesserAI = document.getElementById('guesser-2')

  // Fetching of the values of the student and AI programs
  let resStud = await (await fetching(studProg, data)).json()
  if (resStud.status) throw Error(resStud.msg)
  let resAI = await (await fetching(aiProg, data)).json()
  if (resAI.status) throw Error(resAI.msg)

  guesserStud.innerHTML = ''
  guesserAI.innerHTML = ''

  let workerStud = workerHandler(resStud.ranges, 1),
    workerAI = workerHandler(resAI.ranges, 2)

  // When the `Quick` button is clicked, both workers stop, and every entry is loaded at once to the corresponding displays
  document.getElementById('quick-btn').addEventListener('click', async () => {
    workerStud.terminate()
    workerAI.terminate()

    guesserStud.innerHTML = 'loading...'
    guesserAI.innerHTML = 'loading...'

    let all = await Promise.all([
      showEveryEntry(resStud, guesserStud),
      showEveryEntry(resAI, guesserAI),
    ])

    setTimeout(() => {
      guesserStud.innerHTML = ''
      guesserAI.innerHTML = ''

      new Promise((res, _) => guesserStud.append(...all[0]) && res())
      new Promise((res, _) => guesserAI.append(...all[1]) && res())

      guesserAI.scrollTop = guesserAI.scrollHeight
      guesserStud.scrollTop = guesserStud.scrollHeight

      let lastEntryAI = resAI.ranges[resAI.ranges.length - 1]
      let lastEntryStud = resStud.ranges[resStud.ranges.length - 1]

      let corrPercAI = (lastEntryAI.correct * 100) / resAI.ranges.length
      let corrPercStud = (lastEntryStud.correct * 100) / resStud.ranges.length

      loadResult(lastEntryStud.result, Math.round(corrPercStud * 100) / 100, 1)
      loadResult(lastEntryAI.result, Math.round(corrPercAI * 100) / 100, 2)
    }, 0)
  })

  scrollBottom(guesserStud)
  scrollBottom(guesserAI)
}
