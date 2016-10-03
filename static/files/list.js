var pic = {
        ID:null,
        Title:null,
        Areas:[{areaID:"",rect:{fill:""},Text:""}]
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
      sortOrders: sortOrders
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
app = new Vue({
  el: '#app',
  data: {
    searchQuery: '',
    gridColumns: ['ID','Title','Topic','YearIssued','Status'],
    gridData: [pic],
    pics: [pic],
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
            url: "/pic/all",
            success: function (result) {
                this.$set("gridData", JSON.parse(result));
                $.ajax({
                context: this,
                url: "/person/all",
                success: function (result) {
                    this.$set("persons", JSON.parse(result));
                    //this.setPersonsNames();
            }
        })
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