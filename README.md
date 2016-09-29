# CLI Optionen

## Init

Mit init werden die Daten im DataFolder aus den vorhandenen Quelldateien aus dem SourceFolder erzeugt.

`datatool -init CreateList`

Erzeugt aus den vorhandenen MOD xml Dateien json Files, welche durch den Editor weiter verarbeitet werden können.

`datatool -init ImportTiff`

Hierfür muss auf der ausführenden Maschine ImageMagick installiert sein, da sonst die Tiff Dateien nicht verarbeitet werden können. Der Befehl lädt die Tiff Scans und speichert diese als jpg in den Datafolder.
