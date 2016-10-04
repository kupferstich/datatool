
// create a wrapper around native canvas element (with id="c")
var canvas = new fabric.Canvas('picture');

app = new Vue({
  el: '#app',
  data: {
    fabricElements: [],
    pid: window.location.href.split("/").pop(),
    picSrc: "/img/"+window.location.href.split("/").pop()+"-700-700",
    pic: {
        ID:null,
        Title:null,
        Areas:[{areaID:"",rect:{fill:""},Text:"",Status:""}]
    },
    persons: null,
    newPerson: 0,
    form: [ //Label, JSON Key
        ["Titel","Title"],
        ["Thema","Topic"],
        ["Place","Place"],
        ["Jahr der Anfertigung","YearIssued"]
    ],
    status : ["neu","in Bearbeitung","fertig"],
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
methods: {
    getData: function(cb) {
        $.ajax({
            context: this,
            url: "/pic/"+this.pid,
            success: function (result) {
                this.$set("pic", JSON.parse(result));
                this.loadAreas();
                canvas.renderAll();
                cb;

            }
        })
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
            url: "/pic/"+this.pid,
            data: JSON.stringify(this.pic),
            contentType: "application/json; charset=utf-8",
            dataType: "json",
            success: function(){
              cb;
            }
        });
    },
    addPerson:function(target){
        this.pic.Persons.push(parseInt(this.newPerson));      
    },
    removePerson:function(index){
      this.pic.Persons.splice(index,1)
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
      canvas.add(this.pic.Areas[i].rect);
      this.saveData(
        window.location.href=window.location.href
      );
    },
    loadAreas:function(){
      for (i in this.pic.Areas){
        
        this.fabricElements[i] = new fabric.Rect(this.pic.Areas[i].rect);
       
        this.pic.Areas[i].rect = this.fabricElements[i];
        
        //var group = new fabric.Group([this.pic.Areas[i].rect,label]);
        canvas.add(this.pic.Areas[i].rect);
        //canvas.add(label);
        //canvas.add(group);
      }
      
    },
    addPersonToArea:function(index){
      if (this.pic.Areas[index].Persons==null){
        this.pic.Areas[index].Persons = []
      }
      this.pic.Areas[index].Persons.push(parseInt(this.newPerson)); 
    },
    removePersonFromArea:function(aindex,index){
      this.pic.Areas[aindex].Persons.splice(index,1);
    },
    removeArea:function(index){
      canvas.remove(this.pic.Areas[index].rect)
      this.pic.Areas.splice(index,1)
      this.saveData(
        window.location.href=window.location.href
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
      canvas.add(oImg);
      canvas.sendToBack(oImg);
      });

  $(document)
    .ready(function() {
        $('.ui.accordion')
        .accordion()
;
    })
  ;