import json

# load json file
f = json.load(open("names.json", 'r'))

dictio = {}
for d in f:
    dictio[d.get("region")] = {
        "male": d.get("male"),
        "female": d.get("female"),
        "surnames": d.get("surnames")
        # "field": value
    }

# create a new json file with desired format
f = open("names_new.json", 'w')
json.dump(dictio,f)
