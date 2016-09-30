app = new Vue({
  el: '#app',
  data: {
    pid: window.location.href.split("/").pop(),
    picSrc: "/img/"+window.location.href.split("/").pop()+"-800-800",
    pic: {
        ID:null,
        Title:null
    },
    form: [ //Label, JSON Key
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
    },
    addPerson:function(){
      this.pic.Persons.push({FullName:"",GND:""});
    },
    removePerson:function(index){
      this.pic.Persons.splice(index,1)
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
// create a wrapper around native canvas element (with id="c")
var canvas = new fabric.Canvas('picture');
// create a rectangle object
var rect = new fabric.Rect({
  left: 100,
  top: 100,
  fill: 'red',
  width: 200,
  height: 200,
  opacity: 0.55
});
fabric.Image.fromURL(app.picSrc, function(oImg) {
  oImg.set('selectable', false);
  canvas.add(oImg,rect);
});



// "add" rectangle onto canvas
canvas.add(rect);


  $(document)
    .ready(function() {
        $('.ui.accordion')
        .accordion()
;
    })
  ;