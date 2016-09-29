app = new Vue({
  el: '#app',
  data: {
    test: 'Hello Vue.js!',
    pid: window.location.href.split("/").pop(),
    pic: {
        ID:null,
        Title:null
    },
    form: [ //Label, JSON Key
        ["Picture ID","ID"],
        ["Title","Title"],
        ["Topic","Topic"],
        ["Captured","Captured"],
        ["Place","Place"],
        ["YearIssued","YearIssued"]
    ]
  },
  ready: function() {
    this.getData()
    
},
methods: {
    getData: function() {
        $.ajax({
            context: this,
            url: "/pic/"+this.pid,
            success: function (result) {
                this.$set("pic", JSON.parse(result)) 
            }
        })
    },
    saveData:function() {
        $.ajax({
            type: "POST",
            url: "/pic/"+this.pid,
            data: JSON.stringify(this.pic),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
        });
    }
},
filters: {
  marked: function(value){
    if (value === undefined){
      value = "";
    }
    return marked(value);
  }
}
})
/*
var apiURL = 'https://api.github.com/repos/vuejs/vue/commits?per_page=3&sha='

var demo = new Vue({

  el: '#demo',

  data: {
    branches: ['master', 'dev'],
    currentBranch: 'master',
    commits: null
  },

  created: function () {
    this.fetchData()
  },

  watch: {
    currentBranch: 'fetchData'
  },

  filters: {
    truncate: function (v) {
      var newline = v.indexOf('\n')
      return newline > 0 ? v.slice(0, newline) : v
    },
    formatDate: function (v) {
      return v.replace(/T|Z/g, ' ')
    }
  },

  methods: {
    fetchData: function () {
      var xhr = new XMLHttpRequest()
      var self = this
      xhr.open('GET', apiURL + self.currentBranch)
      xhr.onload = function () {
        self.commits = JSON.parse(xhr.responseText)
      }
      xhr.send()
    }
  }
})

*/