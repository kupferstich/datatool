var postID = window.location.href.split("/").pop();

app = new Vue({
  el: '#app',
  data: {
    newTag : "",
    post: {
        PostPics: {}

    },
    status : [
      "Post in Erstellung",
      "Bilder prüfen",
      "zur Überarbeitung",
      "Metadaten prüfen",
      "fertig"
    ],
    form:[
        ["Titel", "title"],
        ["Beschreibung", "description"],
        ["Bild","image"]
    ],
    formMeta:[
        ["Post Datum", "date"],
        ["Veröffentlichung am", "publishdate"]
    ],
    
  },
  ready: function() {
    this.getData()
    
},
watch: {
       /*post: function(val,oldVal){          
          this.picUrls = [];
          if (this.post.PostPics == undefined) {
              return "";
          }
          for (p in this.post.PostPics){
              //this.person.ProfilePics[p].url = "/img/person/"+personID+"/thumb/"+p;
              //console.log(this.person.ProfilePics[p].url);
              this.picUrls.push("/img/person/"+personID+"/thumb/"+p);
          }
         //return picUrls; 
      }*/
  },
methods: {
    getData: function(cb) {
        this.$http.get("/post/"+postID).then(
        function(res){
          this.$set('post', JSON.parse(res.body));
        }
        )
    },
    setBlogMetaToNow: function(){
      now = new Date();
      this.post.date = now.toJSON();
      this.post.publishdate = now.toJSON();
    },
    addLink:function(){
        if (this.post.Links == null){
            this.post.Links = [];
        }
        var url = this.newLink.Url;
        var title = this.newLink.Title;
        this.post.Links.push({
            Url: url,
            Title: title
        })
        this.newLink.Url = "";
        this.newLink.Title = "";
    },
    removeLink:function(index){
      this.post.Links.splice(index,1)
    },
    saveData:function(redirect,cb) {
        this.unsafedChanges = false;
        this.$http.post("/post/"+this.post.id,JSON.stringify(this.post)).then(
            function(res){
                if(redirect){
                    window.location = "/list/posts";
                }
            }
        )
    },
    addPerson:function(){
      this.post.Persons.push({FullName:"",GND:""});
    },
    removePerson:function(index){
      this.post.Persons.splice(index,1)
    },
    addTag:function(){
       if (this.post.tags == null){
         this.post.tags = [];
       }
        this.post.tags.push(this.newTag);  
        this.newTag = "";    
    },
    removeTag:function(index){
      this.post.tags.splice(index,1)
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
        .accordion();
        var cleave = new Cleave('.datetime',{
          delimiters: ["-","-","T",":",":","Z"],
          blocks: [4,2,2,2,2,2,0],
          uppercase: true
        });
    })
  ;