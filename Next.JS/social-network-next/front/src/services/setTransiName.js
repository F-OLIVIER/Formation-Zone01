function setTransiName() {
  setTimeout(() => {
    let transitext = document.querySelector('.transitext')
    var pathArray = window.location.pathname.split('/');
    if (pathArray[1] === "") {
      pathArray[1] = "Home"
    }
    if (pathArray[1] === "user") {
      pathArray[1] = "Profile"
    }
    transitext.innerText = pathArray[1]
  }, 50);
  }

  export default setTransiName