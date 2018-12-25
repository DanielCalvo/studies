import folium
import pandas
import math

def color_producer(elevation):
    if elevation < 1000:
        return "green"
    elif 1000 <= elevation < 3000:
        return "orange"
    elif elevation > 3000:
        return "red"
    else:
        return "gray"


map = folium.Map(location=[38.58, -99.09], zoom_start=7, tiles="Mapbox Bright")

fgv = folium.FeatureGroup(name="Volcanoes")

data = pandas.read_csv("volcanoes.csv")
#print(data.columns)
#print(type(data))

lat = list(data["LAT"])
lon = list(data["LON"])
elev = list(data["ELEV"])


for lt, ln, el in zip(lat, lon, elev):
    if not math.isnan(lt) and not math.isnan(lt) and not math.isnan(el):
        #fg.add_child(folium.Marker(location=[lt, ln], popup=str(el)+"m", icon=folium.Icon(icon="arrow-dropdown", color=color_producer(el))))
        fgv.add_child(folium.CircleMarker(location=[lt, ln], radius=6, popup=str(el), fill_color=color_producer(el), color="grey", fill_opacity=0.7))


fgp = folium.FeatureGroup(name="Population")

fgp.add_child(folium.GeoJson(data=open("world.json", "r", encoding="utf-8-sig").read(),
                            style_function=lambda x: {"fillColor":"green" if x["properties"]["POP2005"] < 10000000
                            else "orange" if 10000000 <= x["properties"]["POP2005"] < 20000000 else "red"}))

map.add_child(fgv)
map.add_child(fgp)

map.add_child(folium.LayerControl())

map.save("Map1.html")
print("complete!")