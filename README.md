# CLI Optionen

## Init

Mit init werden die Daten im DataFolder aus den vorhandenen Quelldateien aus dem SourceFolder erzeugt.

`datatool -init CreateList`

Erzeugt aus den vorhandenen MOD xml Dateien json Files, welche durch den Editor weiter verarbeitet werden können.

`datatool -init ImportTiff`

Hierfür muss auf der ausführenden Maschine ImageMagick installiert sein, da sonst die Tiff Dateien nicht verarbeitet werden können. Der Befehl lädt die Tiff Scans und speichert diese als jpg in den Datafolder.

# Installation

## Externe Abhängigkeiten

### JavaScript

Innerhalb des Ordners `static/files/library` müssen folgende JavaScript Quellen hinzugefügt werden.

* Jquery: jquery.min.js
* Markdown: marked.min.js  
* Vue: vue.js
* https://github.com/fahrenheit-marketing/jquery-canvas-area-draw

### Semantic UI

Für das Layout wird Semantic UI benötigt. Folgende Dateien müssen im Ornder `static/files/semantic` vorhanden sein:

* semantic.min.js
* semantic.css
