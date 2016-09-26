new Vue({
  el: '#app',
  data: {
    test: 'Hello Vue.js!',
    pid: window.location.href.split("/").pop()
  }
})

