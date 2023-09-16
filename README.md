# Sweather
Sweather is a simple javascript-less\* weather PWA written in HTML5, CSS3, and Golang. The name Sweather is a portmanteau between the words "simple" and "weather". Sweather uses the [National Weather Service's API](https://www.weather.gov/documentation/services-web-api) to get the weather, and the [US Census Bureau Geocoding Service](https://geocoding.geo.census.gov/geocoder/Geocoding_Services_API.html) to turn addresses into latitude and longitude coordinates for use with the other API. Other than those two APIs, Sweather has no dependencies.

There used to be a Heroku instance of Sweather, however, it has been disabled due to Heroku ending their free plan. The code should still work for your own instances.

## Why it was made
Weather apps are too complicated for something as simple as the weather. Sweather aims to be simpler. Sweather doesn't use crazy UI frameworks like React.js, or even JS at all\*. Sweather uses 15 CSS declarations for the entirety of its styling (might be outdated or subject to change, but you get the point). With Sweather, there aren't any downloads, accounts, ads, or non-weather information. Sweather wasn't built to be the best looking, the fastest (I'm working on it though), or even the most accurate or comprehensive. Sweather is made to be simple and usable. Sweather will give you the basic weather you'll need, and if you need more you can just go to another site. 

## Running Sweather
If you want to, you can run a Sweather instance. All you'll need to do is run `go build`, set `$PORT` to your desired port to use (although some services such as Heroku does this for you), and run the compiled binary.

***\* Some JS is actually used but only for the sole purpose of being able to install
and function as a PWA.***
