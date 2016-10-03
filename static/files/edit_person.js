var personID = window.location.href.split("/").pop();

app = new Vue({
  el: '#app',
  data: {
    person: {
        personID:null,
        Title:null,
        FullName:null
    },
    persons: {},
    form: [ //Label, JSON Key
        ["Master","masterID"],
        ["Family Name","FamilyName"],
        ["Given Name","GivenName"],
        ["GND","GND"]
    ]
   
  },
  computed: {
    
  },
  ready: function() {
    this.getData()
    
},
methods: {
    getData: function(cb) {
        $.ajax({
            context: this,
            url: "/person/"+personID,
            success: function (result) {
                this.$set("person", JSON.parse(result));
                cb;

            }
        });
        $.ajax({
            context: this,
            url: "/person/all",
            success: function (result) {
                this.$set("persons", JSON.parse(result));
                cb;

            }
        })
    },
    saveData:function(cb) {
        $.ajax({
            type: "POST",
            url: "/person/"+personID,
            data: JSON.stringify(this.person),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function(){
              cb;
            }
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



  $(document)
    .ready(function() {
        $('.ui.accordion')
        .accordion()
;
    })
  ;