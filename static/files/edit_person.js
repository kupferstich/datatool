var personID = window.location.href.split("/").pop();

app = new Vue({
  el: '#app',
  data: {
    person: {
        personID:null,
        masterID:0,
        GND:0,
        Title:null,
        Text:"",
        FullName:null,
        FamilyName:'',
        GivenName:'',
        ProfilePics:{},
        YearBirth: 0,
        YearDeath: 0,
        CityBirth: '',
        CityDeath: '',
        Links:[]
    },
    persons: {},
    ppUrls:[],
    newLink: {
        Url: "",
        Title: ""
    }
  },
  watch: {
       person: function(val,oldVal){
                    
          this.ppUrls = [];
          if (this.person.ProfilePics == undefined) {
              return "";
          }
          for (p in this.person.ProfilePics){
              //this.person.ProfilePics[p].url = "/img/person/"+personID+"/thumb/"+p;
              //console.log(this.person.ProfilePics[p].url);
              this.ppUrls.push("/img/person/"+personID+"/thumb/"+p);
          }
         //return picUrls; 
      }
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
                this.unsafedChanges = false;
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
    addLink:function(){
        if (this.person.Links == null){
            this.person.Links = [];
        }
        var url = this.newLink.Url;
        var title = this.newLink.Title;
        this.person.Links.push({
            Url: url,
            Title: title
        })
        this.newLink.Url = "";
        this.newLink.Title = "";
    },
    removeLink:function(index){
      this.person.Links.splice(index,1)
    },
    saveData:function(cb) {
         this.unsafedChanges = false;
        $.ajax({
            type: "POST",
            url: "/person/"+personID,
            data: JSON.stringify(this.setIntegers()),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function(){
               
              cb;
            }
        });
    },
    setIntegers:function(){
        this.person.masterID = parseInt(this.person.masterID);
        this.person.GND = parseInt(this.person.GND);
        this.person.YearBirth = parseInt(this.person.YearBirth);
        this.person.YearDeath = parseInt(this.person.YearDeath);
        return this.person;
    },
    setFullName:function(){
        this.person.FullName = this.person.FamilyName + ", " + this.person.GivenName;
    },
    addPerson:function(){
      this.pic.Persons.push({FullName:"",GND:""});
    },
    removePerson:function(index){
      this.pic.Persons.splice(index,1)
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