
// create a wrapper around native canvas element (with id="c")
//var canvas = new fabric.Canvas('picture');

// Constant values of the canvas size
cSize = {
      width: 700,
      height: 700
    };


app = new Vue({
  el: '#app',
  data: {
    canvas: new fabric.Canvas('picture'),
    fabricElements: [],
    pid: window.location.href.split("/").pop(),
    picSrc: "/img/"+window.location.href.split("/").pop()+"-"+cSize.width+"-"+cSize.height,
    pic: {
        ID:null,
        Title:null,
        Areas:[{areaID:"",rect:{fill:""},Text:"",Status:""}],
        canvasWidth: cSize.width,
        canvasHeight: cSize.height
    },
    persons: null,
    newPerson: "",
    tags: null,
    newTag: "",
    form: [ //Label, JSON Key
        ["Titel","Title"],
        ["Thema","Topic"],
        ["Jahr der Anfertigung","YearIssued"]
    ],
    status : [
      "Bild zugeschnitten",
      "Texte gepflegt",
      "zur Überarbeitung",
      "fertig"
    ],
    colors: [
      "black",
      "blue",
      "brown",
      "darkcyan",
      "gold",
      "green",
      "olive",
      "white"
    ]
  },
  computed: {
   
  },
  ready: function() {
    this.getData()
    
},
watch:{
  canvas: function(newVal,oldVal){
    //this.canvas.renderAll();
    //this.pic = newVal;
    //this.loadAreas();
  }
},
methods: {
    getData: function(cb) {
      this.$http.get("/pic/"+this.pid).then(
        function(res){
                var self = this;
          this.$set('pic', JSON.parse(res.body));
          this.canvas.clear();
          //self.canvas = new fabric.Canvas('picture');
          
          this.loadAreas();
          fabric.Image.fromURL(self.picSrc, function(oImg) {
                  oImg.set('selectable', false);
                  self.canvas.add(oImg);
                  self.canvas.sendToBack(oImg);
                  });
          this.canvas.renderAll();
        }
      );
      this.$http.get("/person/all").then(
        function(res){
          this.$set('persons', JSON.parse(res.body));
        }
      );
      
    },
    saveData:function(cb) {
      this.$http.post("/pic/"+this.pid,JSON.stringify(this.pic)).then(
        function(res){
          
          // getData() have to be called, that the areas inside the canvas
          // and the areas overview has been updated inside the view.
          //this.getData();
          if ($.isFunction(cb)){
            cb();
          }
         
        }
      );
    },
    saveDataRedirect:function(){
      this.saveData(function(){
        window.location = "/list/pictures";
      });
    },
    addPerson:function(target){
        this.pic.Persons.push(this.newPerson);     
        this.newPerson = ""; 
    },
    removePerson:function(index){
      this.pic.Persons.splice(index,1)
    },
    addTag:function(target){
       if (this.pic.Tags == null){
         this.pic.Tags = [];
       }
        this.pic.Tags.push(this.newTag);  
        this.newTag = "";    
    },
    removeTag:function(index){
      this.pic.Tags.splice(index,1)
    },
    addArea:function(){
     
      var area = {
      "areaID": "Neuer Bereich",
      "rect": {
        "type": "rect",
        "left": 100,
        "top": 100,
        "width": 100,
        "height": 50,
        "fill": "orange",
        "opacity": 0.5,
        "scaleX": 1,
        "scaleY": 1,
        "hasRotatingPoint": false
      },
      "Shape": "rect",
      "Coords": "",
      "Persons": null,
      "Text": "",
      "Links": null
    }
      var i = this.fabricElements.length
      if (this.pic.Areas == null) {
        this.pic.Areas = [];
      }
      this.pic.Areas[i] = area
      this.fabricElements[i] = new fabric.Rect(this.pic.Areas[i].rect);
      this.pic.Areas[i].rect = this.fabricElements[i];
      this.canvas.add(this.pic.Areas[i].rect);
      this.canvas.renderAll();
      //this.$set('pic',this.pic)
      this.saveData(function(){
        //window.location.href=window.location.href
        //console.log(this);
        app.getData();
      }
      );
    },
    loadAreas:function(){
      for (i in this.pic.Areas){
        this.fabricElements[i] = new fabric.Rect(this.pic.Areas[i].rect);
        this.pic.Areas[i].rect = this.fabricElements[i];
        
        //var group = new fabric.Group([this.pic.Areas[i].rect,label]);
        this.canvas.add(this.pic.Areas[i].rect);
        //canvas.add(label);
        //canvas.add(group);
        // Set canvas information for the picture
        this.pic.canvasWidth = cSize.width;
        this.pic.canvasHeight = cSize.height;
      }
      
    },
    addPersonToArea:function(index){
      if (this.pic.Areas[index].Persons==null){
        this.pic.Areas[index].Persons = []
      }
      this.pic.Areas[index].Persons.push(this.newPerson); 
      this.newPerson = "";
    },
    removePersonFromArea:function(aindex,index){
      this.pic.Areas[aindex].Persons.splice(index,1);
    },
    removeArea:function(index){
      this.canvas.remove(this.pic.Areas[index].rect)
      this.pic.Areas.splice(index,1)
      this.saveData(function(){
            window.location.href=window.location.href
      }
      );
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

fabric.Image.fromURL(app.picSrc, function(oImg) {
                  oImg.set('selectable', false);
                  app.canvas.add(oImg);
                  app.canvas.sendToBack(oImg);
                  });


  $(document)
    .ready(function() {
        $('.ui.accordion')
        .accordion()
;
    })
  ;