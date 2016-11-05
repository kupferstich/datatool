var personID = window.location.href.split("/").pop();

app = new Vue({
    el: '#app',
    data: {
        person: {
            personID: null,
            masterID: 0,
            GND: '',
            Title: null,
            Text: "",
            FullName: null,
            FamilyName: '',
            GivenName: '',
            ProfilePics: {},
            YearBirth: 0,
            YearDeath: 0,
            CityBirth: '',
            CityDeath: '',
            Links: []
        },
        status: [
            "In Erstellung",
            "zur Überarbeitung",
            "Metadaten prüfen",
            "fertig"
        ],
        blogForm: [
            ["Blog Datum", "BlogDate"],
            ["Veröffentlichung am", "PublishDate"]
        ],
        persons: {},
        ppUrls: [],
        newLink: {
            Url: "",
            Title: ""
        }
    },
    watch: {
        person: function (val, oldVal) {

            this.ppUrls = [];
            if (this.person.ProfilePics == undefined) {
                return "";
            }
            for (p in this.person.ProfilePics) {
                //this.person.ProfilePics[p].url = "/img/person/"+personID+"/thumb/"+p;
                //console.log(this.person.ProfilePics[p].url);
                this.ppUrls.push("/img/person/" + personID + "/thumb/" + p);
            }
            //return picUrls; 
        }
    },
    ready: function () {
        this.getData();
        var cleave1 = new Cleave('#BlogDate', {
            delimiters: ["-", "-", "T", ":", ":", "Z"],
            blocks: [4, 2, 2, 2, 2, 2, 0],
            uppercase: true
        });
        var cleave2 = new Cleave('#PublishDate', {
            delimiters: ["-", "-", "T", ":", ":", "Z"],
            blocks: [4, 2, 2, 2, 2, 2, 0],
            uppercase: true
        });

    },
    methods: {
        getData: function (cb) {
            $.ajax({
                context: this,
                url: "/person/" + personID,
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
        addLink: function () {
            if (this.person.Links == null) {
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
        removeLink: function (index) {
            this.person.Links.splice(index, 1)
        },
        saveData: function (redirect, cb) {
            this.unsafedChanges = false;
            this.$http.post("/person/" + personID, JSON.stringify(this.setIntegers())).then(
                function (res) {
                    if (redirect) {
                        window.location = "/list/persons";
                    }
                    // getData() have to be called, that the areas inside the canvas
                    // and the areas overview has been updated inside the view.
                    //this.getData();
                    //cb();
                }
            );
            /* $.ajax({
                 type: "POST",
                 url: "/person/"+personID,
                 data: JSON.stringify(this.setIntegers()),
                 contentType: "application/json; charset=utf-8",
                 dataType: "json",
                 success: function(){
                    if (redirect){
                         window.location = "/list/persons";
                    }
                   //cb;
                 }
             });*/
        },
        setIntegers: function () {
            this.person.masterID = parseInt(this.person.masterID);
            //this.person.GND = parseInt(this.person.GND);
            this.person.YearBirth = parseInt(this.person.YearBirth);
            this.person.YearDeath = parseInt(this.person.YearDeath);
            return this.person;
        },
        setFullName: function () {
            this.person.FullName = this.person.FamilyName + ", " + this.person.GivenName;
        },
        addPerson: function () {
            this.pic.Persons.push({ FullName: "", GND: "" });
        },
        removePerson: function (index) {
            this.pic.Persons.splice(index, 1)
        },
        setBlogMetaToNow: function () {
            now = new Date();
            this.person.BlogDate = now.toJSON();
            this.person.PublishDate = now.toJSON();
        },


    },
    filters: {
        marked: function (value) {
            if (value === undefined) {
                value = "";
            }
            return marked(value);
        }
    }
})


$(document)
    .ready(function () {
        $('.ui.accordion')
            .accordion();

    })
    ;