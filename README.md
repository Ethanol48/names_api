Looking for a new name but you don't think of any? look no more my friend

I develop this API powered by go just for you, just fork this repo run the main.go and call to the port 4000 of your machine with the following format:

```json
"http://localhost:4000/name/{country}/{gender}"
```
the country being the country name (in english) with a Capital letter, per example: Germany, France

and the gender being ``` "Male" ``` or ``` "Female" ```


An example of the response from:

```json
"http://localhost:4000/name/France/Male"
```

```json
{
  "Name": "Alexandre",
  "Surname": "Rousseau"
}
```

but if you feel extra crazy the API has an option of totally random names taken from different languages,

```json
"http://localhost:4000/random"
```
An example of the output:

```json
{
  "Name": "عباس",
  "Surname": "Palmieri"
}
```
you can get pretty crazy with it.

I created the python file to reformat the json, I found that working with the format that it got was a little difficult, so I made a simple script so if you wanted to contribute for the names make the pull request to the "names.json" and the script will do the rest

*This repo is a fork of thm/uninames*
