function setTransiName() {
  setTimeout(() => {
    console.log('coucou')
    let transitext = document.querySelector('.transitext')
    var pathArray = window.location.pathname.split('/');
    console.log(pathArray, transitext)
    if (pathArray[1] === "") {
      pathArray[1] = "Home"
    }
    if (pathArray[1] === "user") {
      pathArray[1] = "Profile"
    }
    transitext.innerText = pathArray[1]
  }, 10);
  }

  export default setTransiName