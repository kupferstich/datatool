function getCookie(name) {
  var value = "; " + document.cookie;
  var parts = value.split("; " + name + "=");
  if (parts.length == 2) return parts.pop().split(";").shift();
}


// register the grid component
Vue.component('grid', {
  template: '#grid-template',
  replace: true,
  props: {
    data: Array,
    columns: Array,
    filterKey: String
  },
  data: function () {
    var sortOrders = {}
    this.columns.forEach(function (key) {
      sortOrders[key] = 1
    })
    return {
      sortKey: '',
      sortOrders: sortOrders,
      linkBase: this.$parent.gridLinkBase
    }
  },
  computed: {
    filteredData: function () {
      var sortKey = this.sortKey
      var filterKey = this.filterKey && this.filterKey.toLowerCase()
      var order = this.sortOrders[sortKey] || 1
      var data = this.data
      if (filterKey) {
        data = data.filter(function (row) {
          return Object.keys(row).some(function (key) {
            return String(row[key]).toLowerCase().indexOf(filterKey) > -1
          })
        })
      }
      if (sortKey) {
        data = data.slice().sort(function (a, b) {
          a = a[sortKey]
          b = b[sortKey]
          return (a === b ? 0 : a > b ? 1 : -1) * order
        })
      }
      return data
    }
  },
  filters: {
    capitalize: function (str) {
      return str.charAt(0).toUpperCase() + str.slice(1)
    }
  },
  methods: {
    sortBy: function (key) {
      this.sortKey = key
      this.sortOrders[key] = this.sortOrders[key] * -1
    }
  }
})


//////////////////
var types = {
    "pictures":{
        "columns":['ID','Title','Topic','YearIssued','BlogDate','Status'],
        "defaultData": {
            ID:null,
            Title:null,
            Areas:[{areaID:"",rect:{fill:""},Text:""}]
        },
        "url": "/pic/all",
        "linkBase": "/form/",
        "dataTrans": function(inData){
            return inData;
        }
    },
    "persons":{
        "columns":['personID','FullName','GND','BlogDate','Status'],
        "defaultData": {
            ID:null,
            Title:null,
            Areas:[{areaID:"",rect:{fill:""},Text:""}]
        },
        "url": "/person/all",
        "linkBase": "/edit/person/",
        "dataTrans": function(inData){
            return  Object.keys(inData).map(function (key) { 
                inData[key].ID = key;
                return inData[key]; });
        }
    },
    "posts":{
        "columns":['id','title','date','status'],
        "defaultData": {
            ID:null,
            Title:null,
            Areas:[{areaID:"",rect:{fill:""},Text:""}]
        },
        "url": "/post/all",
        "linkBase": "/edit/post/",
        "dataTrans": function(inData){
            return  Object.keys(inData).map(function (key) { 
                inData[key].ID = key;
                return inData[key]; });
        }
    }
}
var type = window.location.href.split("/").pop()

var searchQueryCookie = getCookie("searchQuery");
app = new Vue({
  el: '#app',
  data: {
    searchQuery: '',
    gridColumns: types[type]['columns'],
    gridData: types[type]['defaultData'],
    gridLinkBase: types[type]['linkBase'],
    persons: {}
  },
  computed: {
    
  },
  ready: function() {
    this.getData()
    
},
methods: {
    getData: function() {
        $.ajax({
            context: this,
            url: types[type]['url'],
            success: function (result) {
                var trans = types[type]['dataTrans'] 
                this.$set("gridData", trans(JSON.parse(result)));
            }
        });
        
    },
    setPersonsNames: function(){
        self = this;
        this.gridData.forEach(
            function(pic,index){
                pic.Persons.forEach(
                    function(person,pindex){
                        self.gridData[index].PersonsName[pindex] = self.persons[person]
                    }
                )
            }
        )
    },
    saveData:function(cb) {
        $.ajax({
            type: "POST",
            url: "/pic/"+this.pid,
            data: JSON.stringify(this.pic),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function(){
              cb;
            }
        });
    },
   
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



  $(document)
    .ready(function() {
        $('.ui.accordion')
        .accordion()
;
    })
  ;