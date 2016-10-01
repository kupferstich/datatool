
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
        Areas:[{areaID:"",rect:{fill:""},Text:""}]
    },
    form: [ //Label, JSON Key
        ["Title","Title"],
        ["Topic","Topic"],
        ["Captured","Captured"],
        ["Place","Place"],
        ["YearIssued","YearIssued"]
    ]
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
            url: "/pic/"+this.pid,
            success: function (result) {
                this.$set("pic", JSON.parse(result));
                this.loadAreas();

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
    },
    loadAreas:function(){
      for (i in this.pic.Areas){
        
        this.fabricElements[i] = new fabric.Rect(this.pic.Areas[i].rect);
        var label = new fabric.Text(this.pic.Areas[i].areaID, {
          left: this.pic.Areas[i].rect.left,
          top: this.pic.Areas[i].rect.top,
          fontSize: 20
        });
        
        this.pic.Areas[i].rect = this.fabricElements[i];
        
        //var group = new fabric.Group([this.pic.Areas[i].rect,label]);
        canvas.add(this.pic.Areas[i].rect);
        canvas.add(label);
        //canvas.add(group);
      }
      
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