onmessage = (data) => {
  let i = 0
  let ranges = data.data
  // The worker keeps responding in an interval of 16 ms the ranges to the frontend display
  let int = setInterval(() => {
    i == 0
      ? postMessage([ranges[i], i, 'first_value'])
      : postMessage([ranges[i], i])

    i++

    if (i == ranges.length) {
      clearInterval(int)
      postMessage(
        `\nFinal result: ${numberWithSpaces(
          ranges[ranges.length - 1].result
        )}\n`
      )
    }
  }, 16)
}

const numberWithSpaces = (x) => {
  return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ' ')
}
